package conf

import (
	"time"
)

// Option alias func as option
type Option func(opts *Options)

// defaultOptions set default options
var defaultOptions = Options{
	Server: ServerOptions{
		HTTP: HTTPOptions{
			Addr:    ":80",
			Timeout: time.Second,
		},
		GRPC: GRPCOptions{
			Addr:    ":81",
			Timeout: time.Second,
		},
	},
	Data: DataOptions{
		Database: DatabaseOptions{
			Driver:     "mysql",
			DataSource: "",
		},
		Redis: RedisOptions{
			URL:  "127.0.0.1",
			Port: 6379,
		},
	},
	Mode: "debug",
}

// Options ...
type Options struct {
	Server ServerOptions
	Data   DataOptions
	Mode   string
}

// New return new options
func New(opts ...Option) Options {
	options := defaultOptions
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

// WithServer ...
func WithServer(server ServerOptions) Option {
	return func(opts *Options) {
		opts.Server = server
	}
}

// WithData ...
func WithData(data DataOptions) Option {
	return func(opts *Options) {
		opts.Data = data
	}
}

// WithMode ...
func WithMode(mode string) Option {
	return func(opts *Options) {
		opts.Mode = mode
	}
}
