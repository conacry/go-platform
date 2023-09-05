package postgresql

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type Logger interface {
	LogError(ctx context.Context, errs ...error)
	LogInfo(ctx context.Context, messages ...string)
}

type Config struct {
	Url             string
	MaxConnIdleTime time.Duration
	MaxConns        int32
	MinConns        int32
}

type PostgreSQL struct {
	config   *Config
	logger   Logger
	connPool *pgxpool.Pool
}

func (pg *PostgreSQL) Start(ctx context.Context) error {
	pgConfig, err := pgxpool.ParseConfig(pg.config.Url)
	if err != nil {
		pg.logger.LogError(ctx, err)
		return err
	}

	pgConfig.MinConns = pg.config.MinConns
	pgConfig.MaxConns = pg.config.MaxConns
	pgConfig.MaxConnIdleTime = pg.config.MaxConnIdleTime

	connPool, err := pgxpool.NewWithConfig(ctx, pgConfig)
	if err != nil {
		pg.logger.LogError(ctx, err)
		return err
	}

	err = connPool.Ping(ctx)
	if err != nil {
		pg.logger.LogError(ctx, err)
		return err
	}

	pg.connPool = connPool
	pg.logger.LogInfo(ctx, "PostgreSQL connection is initialized")

	return nil
}

func (pg *PostgreSQL) Stop(ctx context.Context) error {
	pg.connPool.Close()
	pg.logger.LogInfo(ctx, "PostgreSQL connection is closed")
	return nil
}

func (pg *PostgreSQL) ConnPool() *pgxpool.Pool {
	return pg.connPool
}
