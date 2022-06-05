package conf

type (
	DatabaseOptions struct {
		Driver     string
		DataSource string
	}

	RedisOptions struct {
		URL  string
		Port int
	}
)

// DataOptions database and redis config struct ...
type DataOptions struct {
	Database DatabaseOptions
	Redis    RedisOptions
}

// WithDatabase ...
func WithDatabase(db DatabaseOptions) Option {
	return func(opts *Options) {
		opts.Data.Database = db
	}
}

// WithRedis ...
func WithRedis(redis RedisOptions) Option {
	return func(opts *Options) {
		opts.Data.Redis = redis
	}
}
