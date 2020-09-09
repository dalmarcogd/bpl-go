package services

import (
	"context"
	"fmt"
	"github.com/dalmarcogd/bpl-go/internal/models"
	"gorm.io/gorm"
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
		DB(ctx context.Context) *gorm.DB
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
	Environment interface {
		Generic
		WithServiceManager(c ServiceManager) Environment
		Environment() string
		Service() string
		Version() string
		DebugPprof() bool
		DatabaseDsn() string
		CacheAddress() string
	}
	Handlers interface {
		Generic
		WithServiceManager(c ServiceManager) Handlers
		CreateUser(ctx context.Context, u *models.User) error
		UpdateUser(ctx context.Context, u *models.User) error
		GetUser(ctx context.Context, u *models.User) error
		GetUsers(ctx context.Context, u *[]models.User) error
		DeleteUser(ctx context.Context, u *models.User) error
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
		WithHandlers(d Handlers) ServiceManager
		Handlers() Handlers
		WithEnvironment(d Environment) ServiceManager
		Environment() Environment

		Context() context.Context
		Init() error
		Close() error
	}

	ServiceManagerImpl struct {
		ctx         context.Context
		cancel      context.CancelFunc
		database    Database
		cache       Cache
		log         Logger
		httpServer  HttpServer
		handlers    Handlers
		environment Environment
	}
)

func New() *ServiceManagerImpl {
	return &ServiceManagerImpl{}
}

func (s *ServiceManagerImpl) Init() error {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	if err := s.Logger().Init(s.ctx); err != nil {
		return err
	}
	if err := s.Environment().Init(s.ctx); err != nil {
		return err
	}
	if err := s.HttpServer().Init(s.ctx); err != nil {
		return err
	}
	if err := s.Cache().Init(s.ctx); err != nil {
		return err
	}
	if err := s.Database().Init(s.ctx); err != nil {
		return err
	}
	s.Logger().Info(s.ctx, "All services initialized")
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
	if errC := s.httpServer.Close(); errC != nil {
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

func (s *ServiceManagerImpl) WithHandlers(d Handlers) ServiceManager {
	s.handlers = d.WithServiceManager(s)
	return s
}

func (s *ServiceManagerImpl) Handlers() Handlers {
	return s.handlers
}
func (s *ServiceManagerImpl) WithEnvironment(d Environment) ServiceManager {
	s.environment = d.WithServiceManager(s)
	return s
}

func (s *ServiceManagerImpl) Environment() Environment {
	return s.environment
}
