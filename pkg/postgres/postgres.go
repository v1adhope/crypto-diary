package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/v1adhope/crypto-diary/internal/config"
)

// TODO: ???
// type Client interface {
// 	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
// 	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
// 	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
// }

type Postgres struct {
	connAttempts int
	connTimeout  time.Duration
	poolsize     int32

	Pool *pgxpool.Pool
}

// TODO: Separate configure
func NewClient(cfg *config.Config) (*Postgres, error) {
	pg := &Postgres{
		connAttempts: cfg.Storage.ConnAttempts,
		connTimeout:  cfg.Storage.ConnTimeout,
		poolsize:     cfg.Storage.PoolSize,
	}
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Storage.Username, cfg.Storage.Password, cfg.Storage.Host, cfg.Storage.Port, cfg.Storage.Database)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: parse config failed: %s", err)
	}

	poolCfg.MaxConns = pg.poolsize

	pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, fmt.Errorf("unable to create pool: %s", err)
	}

	for pg.connAttempts > 0 {

		err = pg.Pool.Ping(context.Background())
		if err == nil {
			break
		}
		log.Printf("ping failed: attempts left %d: %s", pg.connAttempts, err)

		time.Sleep(pg.connTimeout * time.Second)

		pg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %s", err)
	}
	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}