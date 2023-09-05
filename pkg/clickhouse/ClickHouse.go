package clickhouse

import (
	"context"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Logger interface {
	LogError(ctx context.Context, errs ...error)
	LogInfo(ctx context.Context, messages ...string)
}

type Config struct {
	Url          string
	Database     string
	Username     string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
}

type ClickHouse struct {
	config *Config
	logger Logger
	conn   driver.Conn
}

func (ch *ClickHouse) Start(ctx context.Context) error {
	opt := &clickhouse.Options{
		Protocol: clickhouse.Native,
		Addr:     []string{ch.config.Url},
		Auth: clickhouse.Auth{
			Database: ch.config.Database,
			Username: ch.config.Username,
			Password: ch.config.Password,
		},
		MaxOpenConns: ch.config.MaxOpenConns,
		MaxIdleConns: ch.config.MaxIdleConns,
	}

	conn, err := clickhouse.Open(opt)
	if err != nil {
		ch.logger.LogError(ctx, err)
		return err
	}

	if err := conn.Ping(ctx); err != nil {
		var exception *clickhouse.Exception
		if errors.As(err, &exception) {
			msg := fmt.Sprintf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
			wrappedErr := errors.New(msg)
			ch.logger.LogError(ctx, wrappedErr)
			return err
		}

		ch.logger.LogError(ctx, err)
		return err
	}

	ch.conn = conn
	ch.logger.LogInfo(ctx, "ClickHouse connection is initialized")

	return nil
}

func (ch *ClickHouse) Stop(ctx context.Context) error {
	err := ch.conn.Close()
	if err != nil {
		ch.logger.LogError(ctx, err)
		return err
	}

	ch.logger.LogInfo(ctx, "ClickHouse connection is closed")
	return nil
}

func (ch *ClickHouse) Select(ctx context.Context, result any, query string, queryArgs map[string]any) error {
	queryParams := make([]any, 0, len(queryArgs))
	for argName, argValue := range queryArgs {
		queryParam := clickhouse.Named(argName, argValue)
		queryParams = append(queryParams, queryParam)
	}

	err := ch.conn.Select(ctx, result, query, queryParams...)
	if err != nil {
		ch.logger.LogError(ctx, err)
		return err
	}

	return nil
}
