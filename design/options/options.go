package options

import (
	"time"
)

/**
1. 需要Option选项得对象，里面是每一个配置
2. 需要一个Option得接口，拥有一个设置option得函数
3. 需要一个自定义类型得func，实现了Option接口
4. func要返回一个Option类型得接口，并且参数为要设置得内容
*/

type Connection struct {
	addr    string
	cache   bool
	timeout time.Duration
}

const (
	defaultTimeout = 10
	defaultCaching = false
)

type options struct {
	timeout time.Duration
	caching bool
}

type Option func(*options)

func WithTimeout(t time.Duration) Option {
	// 定义函数具体逻辑
	return func(o *options) {
		o.timeout = t
	}
}

func WithCaching(cache bool) Option {
	return func(o *options) {
		o.caching = cache
	}
}

// Connect creates a connection.
func Connect(addr string, opts ...Option) (*Connection, error) {
	options := options{
		timeout: defaultTimeout,
		caching: defaultCaching,
	}

	for _, o := range opts {
		o(&options)
	}

	return &Connection{
		addr:    addr,
		cache:   options.caching,
		timeout: options.timeout,
	}, nil
}
