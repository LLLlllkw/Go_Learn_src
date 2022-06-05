package conf

import (
	"time"
)

type (
	HTTPOptions struct {
		Addr    string
		Timeout time.Duration
	}
	GRPCOptions struct {
		Addr    string
		Timeout time.Duration
	}
)

// server.HttpServer and server.GRPCServer config struct ...
type ServerOptions struct {
	HTTP HTTPOptions
	GRPC GRPCOptions
}

// WithHTTP ...
func WithHTTP(http HTTPOptions) Option {
	return func(opts *Options) {
		opts.Server.HTTP = http
	}
}

// WithGRPC ...
func WithGRPC(grpc GRPCOptions) Option {
	return func(opts *Options) {
		opts.Server.GRPC = grpc
	}
}

// WithHTTPTimeOut ...
func WithHTTPTimeOut(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Server.HTTP.Timeout = timeout
	}
}

// WithHTTPAddr ...
func WithHTTPAddr(addr string) Option {
	return func(opts *Options) {
		opts.Server.HTTP.Addr = addr
	}
}

// WithGRPCAddr ...
func WithGRPCAddr(addr string) Option {
	return func(opts *Options) {
		opts.Server.GRPC.Addr = addr
	}
}

// WithGRPCTimeOut ...
func WithGRPCTimeOut(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Server.GRPC.Timeout = timeout
	}
}
