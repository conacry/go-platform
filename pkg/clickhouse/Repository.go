package clickhouse

import "context"

type Repository interface {
	Select(ctx context.Context, result any, query string, queryArgs map[string]any) error
}
