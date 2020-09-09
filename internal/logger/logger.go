package logger

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/bpl-go/internal/services"
	"github.com/sirupsen/logrus"
	"runtime"
)

type (
	DefaultFields func(ctx context.Context, fields *map[string]interface{})
	ServiceImpl   struct {
		serviceManager services.ServiceManager
		ctx            context.Context
		logger         *logrus.Logger
		defaultFields  DefaultFields
	}
)

func New() *ServiceImpl {
	return &ServiceImpl{}
}

func (s *ServiceImpl) WithDefaultFields(f DefaultFields) *ServiceImpl {
	s.defaultFields = f
	return s
}

func (s *ServiceImpl) Init(ctx context.Context) error {
	s.ctx = ctx
	s.logger = logrus.New()
	return nil
}

func (s *ServiceImpl) Close() error {
	return nil
}

func (s *ServiceImpl) WithServiceManager(c services.ServiceManager) services.Logger {
	s.serviceManager = c
	return s
}

func (s *ServiceImpl) ServiceManager() services.ServiceManager {
	return s.serviceManager
}

func (s *ServiceImpl) output(ctx context.Context, fields ...map[string]interface{}) *logrus.Entry {
	f := make(map[string]interface{}, 0)

	if s.defaultFields != nil {
		s.defaultFields(ctx, &f)
	}

	for _, ff := range fields {
		for k, v := range ff {
			f[k] = v
		}
	}

	stack := ""
	pc, _, line, ok := runtime.Caller(2)
	if ok {
		stack = fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), line)
	}
	f["caller"] = stack

	return s.logger.WithFields(f)
}

func (s *ServiceImpl) Info(ctx context.Context, message string, fields ...map[string]interface{}) {
	s.output(ctx, fields...).Info(message)
}

func (s *ServiceImpl) Warn(ctx context.Context, message string, fields ...map[string]interface{}) {
	s.output(ctx, fields...).Warn(message)
}

func (s *ServiceImpl) Error(ctx context.Context, message string, fields ...map[string]interface{}) {
	s.output(ctx, fields...).Error(message)
}

func (s *ServiceImpl) Fatal(ctx context.Context, message string, fields ...map[string]interface{}) {
	s.output(ctx, fields...).Fatal(message)
}
