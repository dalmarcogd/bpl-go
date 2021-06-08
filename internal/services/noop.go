package services

import (
	"context"
	"database/sql"
	"github.com/dalmarcogd/bpl-go/internal/models"
)

type (
	NoopHealth   struct{}
	NoopSis      struct{}
	NoopDatabase struct {
		NoopHealth
	}
	NoopValidator struct {
		NoopHealth
	}
	NoopHttpServer struct {
		NoopHealth
	}
	NoopCache struct {
		NoopHealth
	}
	NoopLogger struct {
		NoopHealth
	}
	NoopHandlers struct {
		NoopHealth
	}
	NoopEnvironment struct {
		NoopHealth
	}
)

func (n *NoopHealth) Health(_ context.Context) error {
	return nil
}

func NewNoopValidator() *NoopValidator {
	return &NoopValidator{}
}

func (n *NoopValidator) Sis() Sis {
	return nil
}

func (n *NoopValidator) Init(_ context.Context) error {
	return nil
}

func (n *NoopValidator) Close() error {
	return nil
}

func (n *NoopValidator) WithSis(_ Sis) Validator {
	return n
}

func (n *NoopValidator) Validate(_ context.Context, _ interface{}) error {
	return nil
}

func (n *NoopValidator) ValidateSlice(_ context.Context, _ interface{}) error {
	return nil
}

func NewNoopDatabase() *NoopDatabase {
	return &NoopDatabase{}
}

func (n *NoopDatabase) Sis() Sis {
	return nil
}

func (n *NoopDatabase) Init(_ context.Context) error {
	return nil
}

func (n *NoopDatabase) Close() error {
	return nil
}

func (n *NoopDatabase) WithSis(_ Sis) Database {
	return n
}

func (n *NoopDatabase) WithCardsAutomaticUpdaterDatabase() Database {
	return n
}

func (n *NoopDatabase) WithCardsEmbossingsDatabase() Database {
	return n
}

func (n *NoopDatabase) WithCoreDatabase() Database {
	return n
}

func (n *NoopDatabase) WithMasterClient(_ *sql.DB) Database {
	return nil
}

func (n *NoopDatabase) WithReplicaClient(_ *sql.DB) Database {
	return nil
}

func (n *NoopDatabase) Ping(_ context.Context) error {
	return nil
}

func (n *NoopDatabase) OpenTransactionMaster(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (n *NoopDatabase) TransactionMaster(context.Context, func(tx DatabaseTransaction) error) error {
	return nil
}

func (n *NoopDatabase) OpenTransactionReplica(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (n *NoopDatabase) TransactionReplica(context.Context, func(tx DatabaseTransaction) error) error {
	return nil
}

func (n *NoopDatabase) CloseTransaction(_ context.Context, err error) error {
	return err
}

func (n *NoopDatabase) Insert(_ context.Context, _ string, _ ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (n *NoopDatabase) Update(_ context.Context, _ string, _ ...interface{}) (sql.Result, error) {
	return nil, nil
}

func (n *NoopDatabase) Get(_ context.Context, _ string, _ ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

func NewNoopHttpServer() *NoopHttpServer {
	return &NoopHttpServer{}
}

func (n *NoopHttpServer) Sis() Sis {
	return nil
}

func (n *NoopHttpServer) Init(_ context.Context) error {
	return nil
}

func (n *NoopHttpServer) Close() error {
	return nil
}

func (n *NoopHttpServer) WithSis(_ Sis) HttpServer {
	return n
}

func (n *NoopHttpServer) Run() error {
	return nil
}

func NewNoopCache() *NoopCache {
	return &NoopCache{}
}

func (n *NoopCache) Sis() Sis {
	return nil
}

func (n *NoopCache) Init(_ context.Context) error {
	return nil
}

func (n *NoopCache) Close() error {
	return nil
}

func (n *NoopCache) WithSis(_ Sis) Cache {
	return n
}

func NewNoopLogger() *NoopLogger {
	return &NoopLogger{}
}

func (n *NoopLogger) Sis() Sis {
	return nil
}

func (n *NoopLogger) Init(_ context.Context) error {
	return nil
}

func (n *NoopLogger) Close() error {
	return nil
}

func (n *NoopLogger) WithSis(_ Sis) Logger {
	return n
}

func (n *NoopLogger) Info(_ context.Context, _ string, _ ...map[string]interface{}) {}

func (n *NoopLogger) Warn(_ context.Context, _ string, _ ...map[string]interface{}) {}

func (n *NoopLogger) Error(_ context.Context, _ string, _ ...map[string]interface{}) {}

func (n *NoopLogger) Fatal(_ context.Context, _ string, _ ...map[string]interface{}) {}

func NewNoopHandlers() *NoopHandlers {
	return &NoopHandlers{}
}

func (n *NoopHandlers) Sis() Sis {
	return nil
}

func (n *NoopHandlers) Init(_ context.Context) error {
	return nil
}

func (n *NoopHandlers) Close() error {
	return nil
}

func (n *NoopHandlers) WithSis(_ Sis) Handlers {
	return n
}

func (n *NoopHandlers) CreateUser(_ context.Context, _ *models.User) error {
	return nil
}

func (n *NoopHandlers) UpdateUser(_ context.Context, _ *models.User) error {
	return nil
}

func (n *NoopHandlers) GetUser(_ context.Context, _ *models.User) error {
	return nil
}

func (n *NoopHandlers) GetUsers(_ context.Context, _ *[]models.User) error {
	return nil
}

func (n *NoopHandlers) DeleteUser(_ context.Context, _ *models.User) error {
	return nil
}

func NewNoopEnvironment() *NoopEnvironment {
	return &NoopEnvironment{}
}

func (n *NoopEnvironment) Sis() Sis {
	return nil
}

func (n *NoopEnvironment) Init(_ context.Context) error {
	return nil
}

func (n *NoopEnvironment) Close() error {
	return nil
}

func (n *NoopEnvironment) WithSis(_ Sis) Environment {
	return n
}

func (n *NoopEnvironment) Environment() string {
	return ""
}

func (n *NoopEnvironment) Service() string {
	return ""
}

func (n *NoopEnvironment) Version() string {
	return ""
}

func (n *NoopEnvironment) DebugPprof() bool {
	return false
}

func (n *NoopEnvironment) DatabaseDsn() string {
	return ""
}

func (n *NoopEnvironment) CacheAddress() string {
	return ""
}
