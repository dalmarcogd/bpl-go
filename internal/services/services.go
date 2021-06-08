package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dalmarcogd/bpl-go/internal/models"
)

type (
	Generic interface {
		Sis() Sis
		Init(ctx context.Context) error
		Health(ctx context.Context) error
		Close() error
	}
	Database interface {
		Generic
		WithSis(c Sis) Database
		WithCoreDatabase() Database
		WithMasterClient(*sql.DB) Database
		WithReplicaClient(*sql.DB) Database
		OpenTransactionMaster(ctx context.Context) (context.Context, error)
		TransactionMaster(ctx context.Context, f func(tx DatabaseTransaction) error) error
		OpenTransactionReplica(ctx context.Context) (context.Context, error)
		TransactionReplica(ctx context.Context, f func(tx DatabaseTransaction) error) error
		CloseTransaction(ctx context.Context, err error) error
	}
	DatabaseTransaction interface {
		Insert(query string, args ...interface{}) (sql.Result, error)
		Update(query string, args ...interface{}) (sql.Result, error)
		Get(query string, args ...interface{}) (*sql.Rows, error)
	}
	Validator interface {
		Generic
		WithSis(c Sis) Validator
		Validate(ctx context.Context, obj interface{}) error
		ValidateSlice(ctx context.Context, objs interface{}) error
	}
	Cache interface {
		Generic
		WithSis(c Sis) Cache
	}
	Logger interface {
		Generic
		WithSis(c Sis) Logger
		Info(ctx context.Context, message string, fields ...map[string]interface{})
		Warn(ctx context.Context, message string, fields ...map[string]interface{})
		Error(ctx context.Context, message string, fields ...map[string]interface{})
		Fatal(ctx context.Context, message string, fields ...map[string]interface{})
	}
	HttpServer interface {
		Generic
		WithSis(c Sis) HttpServer
		Run() error
	}
	Environment interface {
		Generic
		WithSis(c Sis) Environment
		Environment() string
		Service() string
		Version() string
		DebugPprof() bool
		DatabaseDsn() string
		CacheAddress() string
	}
	Handlers interface {
		Generic
		WithSis(c Sis) Handlers
		CreateUser(ctx context.Context, u *models.User) error
		UpdateUser(ctx context.Context, u *models.User) error
		GetUser(ctx context.Context, u *models.User) error
		GetUsers(ctx context.Context, u *[]models.User) error
		DeleteUser(ctx context.Context, u *models.User) error
	}

	Sis interface {
		WithDatabase(d Database) Sis
		Database() Database
		WithCache(d Cache) Sis
		Cache() Cache
		WithLogger(d Logger) Sis
		Logger() Logger
		WithHttpServer(d HttpServer) Sis
		HttpServer() HttpServer
		WithHandlers(d Handlers) Sis
		Handlers() Handlers
		WithEnvironment(d Environment) Sis
		Environment() Environment

		Context() context.Context
		Init() error
		Close() error
		Health(ctx context.Context) error
	}

	sisImpl struct {
		ctx         context.Context
		cancel      context.CancelFunc
		database    Database
		validator   Validator
		cache       Cache
		log         Logger
		httpServer  HttpServer
		handlers    Handlers
		environment Environment
	}
)

func New() *sisImpl {
	return &sisImpl{
		database:    NewNoopDatabase(),
		cache:       NewNoopCache(),
		log:         NewNoopLogger(),
		httpServer:  NewNoopHttpServer(),
		handlers:    NewNoopHandlers(),
		environment: NewNoopEnvironment(),
		validator:   NewNoopValidator(),
	}
}

func (s *sisImpl) Init() error {
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

func (s *sisImpl) Health(ctx context.Context) error {
	if err := s.Logger().Health(ctx); err != nil {
		return err
	}
	if err := s.Environment().Health(s.ctx); err != nil {
		return err
	}
	if err := s.HttpServer().Health(s.ctx); err != nil {
		return err
	}
	if err := s.Cache().Health(s.ctx); err != nil {
		return err
	}
	if err := s.Database().Health(s.ctx); err != nil {
		return err
	}
	s.Logger().Info(s.ctx, "All services initialized")
	return nil
}

func (s *sisImpl) Close() error {
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

func (s *sisImpl) Context() context.Context {
	return s.ctx
}

func (s *sisImpl) WithDatabase(d Database) Sis {
	s.database = d.WithSis(s)
	return s
}

func (s *sisImpl) Database() Database {
	return s.database
}

func (s *sisImpl) WithCache(d Cache) Sis {
	s.cache = d.WithSis(s)
	return s
}

func (s *sisImpl) Cache() Cache {
	return s.cache
}

func (s *sisImpl) WithLogger(d Logger) Sis {
	s.log = d.WithSis(s)
	return s
}

func (s *sisImpl) Logger() Logger {
	return s.log
}

func (s *sisImpl) WithHttpServer(d HttpServer) Sis {
	s.httpServer = d.WithSis(s)
	return s
}

func (s *sisImpl) HttpServer() HttpServer {
	return s.httpServer
}

func (s *sisImpl) WithHandlers(d Handlers) Sis {
	s.handlers = d.WithSis(s)
	return s
}

func (s *sisImpl) Handlers() Handlers {
	return s.handlers
}
func (s *sisImpl) WithEnvironment(d Environment) Sis {
	s.environment = d.WithSis(s)
	return s
}

func (s *sisImpl) Environment() Environment {
	return s.environment
}
