package services

import (
	"context"
	"fmt"
)

type (
	Generic interface {
		ServiceManager() ServiceManager
		Init(ctx context.Context) error
		Close() error
	}

	Database interface {
		Generic
		WithServiceManager(c ServiceManager) Database
	}
	Cache interface {
		Generic
		WithServiceManager(c ServiceManager) Cache
	}
	Logger interface {
		Generic
		WithServiceManager(c ServiceManager) Logger
		Info(ctx context.Context, message string, fields ...map[string]interface{})
		Warn(ctx context.Context, message string, fields ...map[string]interface{})
		Error(ctx context.Context, message string, fields ...map[string]interface{})
		Fatal(ctx context.Context, message string, fields ...map[string]interface{})
	}
	HttpServer interface {
		Generic
		WithServiceManager(c ServiceManager) HttpServer
		Run() error
	}

	ServiceManager interface {
		WithDatabase(d Database) ServiceManager
		Database() Database
		WithCache(d Cache) ServiceManager
		Cache() Cache
		WithLogger(d Logger) ServiceManager
		Logger() Logger
		WithHttpServer(d HttpServer) ServiceManager
		HttpServer() HttpServer

		Context() context.Context
		Init() error
		Close() error
	}

	ServiceManagerImpl struct {
		ctx        context.Context
		cancel     context.CancelFunc
		database   Database
		cache      Cache
		log        Logger
		httpServer HttpServer
	}
)

func New() *ServiceManagerImpl {
	return &ServiceManagerImpl{}
}

func (s *ServiceManagerImpl) Init() error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	if err := s.log.Init(s.ctx); err != nil {
		return err
	}
	if err := s.httpServer.Init(s.ctx); err != nil {
		return err
	}
	if err := s.cache.Init(s.ctx); err != nil {
		return err
	}
	if err := s.database.Init(s.ctx); err != nil {
		return err
	}
	return nil
}

func (s *ServiceManagerImpl) Close() error {
	var err error
	if errC := s.cache.Close(); errC != nil {
		err = errC
	}
	if errC := s.database.Close(); errC != nil {
		err = fmt.Errorf("%v - %v", err, errC)
	}
	if errC := s.log.Close(); errC != nil {
		err = fmt.Errorf("%v - %v", err, errC)
	}
	s.cancel()
	return err
}

func (s *ServiceManagerImpl) Context() context.Context {
	return s.ctx
}

func (s *ServiceManagerImpl) WithDatabase(d Database) ServiceManager {
	s.database = d.WithServiceManager(s)
	return s
}

func (s *ServiceManagerImpl) Database() Database {
	return s.database
}

func (s *ServiceManagerImpl) WithCache(d Cache) ServiceManager {
	s.cache = d.WithServiceManager(s)
	return s
}

func (s *ServiceManagerImpl) Cache() Cache {
	return s.cache
}

func (s *ServiceManagerImpl) WithLogger(d Logger) ServiceManager {
	s.log = d.WithServiceManager(s)
	return s
}

func (s *ServiceManagerImpl) Logger() Logger {
	return s.log
}

func (s *ServiceManagerImpl) WithHttpServer(d HttpServer) ServiceManager {
	s.httpServer = d.WithServiceManager(s)
	return s
}

func (s *ServiceManagerImpl) HttpServer() HttpServer {
	return s.httpServer
}
