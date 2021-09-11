package finviz

import "time"

type Options struct {
	cacheDuration time.Duration
}

func NewOptions() *Options {
	return &Options{
		cacheDuration: 3 * time.Minute,
	}
}

func (o *Options) WithCacheDuration(t time.Duration) *Options {
	o.cacheDuration = t
	return o
}
