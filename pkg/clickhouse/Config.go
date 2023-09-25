package clickhouse

type Config struct {
	URL          string
	Database     string
	Username     string
	Password     string
	MaxOpenConns int
	MaxIdleConns int
}
