package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

// TODO ???
type PostgresClient interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type StorageConfig struct {
	Username, Password, Host, Port, Database string
}

// TODO: defer dbpool.Close()
func NewClient(ctx context.Context, sc StorageConfig) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	pool, err := pgxpool.New(context.TODO(), dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create client: %v")
	}
	return pool, nil
}
