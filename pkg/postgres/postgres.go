package postgres

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Username     string        `mapstructure:"username"`
	Password     string        `mapstructure:"password"`
	Socket       string        `mapstructure:"socket"`
	Database     string        `mapstructure:"database"`
	ConnAttempts int           `mapstructure:"conn_attempts"`
	ConnTimeout  time.Duration `mapstructure:"conn_timeout"`
	PoolSize     int32         `mapstructure:"pool_size"`
}

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewClient(ctx context.Context, cfg *Config) (*Postgres, error) {
	pg := &Postgres{}

	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s", cfg.Username, cfg.Password, cfg.Socket, cfg.Database)

	poolCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: parse config failed: %s", err)
	}

	poolCfg.MaxConns = cfg.PoolSize

	pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, fmt.Errorf("postgres: unable to create pool: %s", err)
	}

	for cfg.ConnAttempts > 0 {

		err = pg.Pool.Ping(ctx)
		if err == nil {
			break
		}
		log.Printf("postgres: ping failed: attempts left %d: %s", cfg.ConnAttempts, err)

		time.Sleep(cfg.ConnTimeout)

		cfg.ConnAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("postgres: unable to create connection pool: %s", err)
	}
	return pg, nil
}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
