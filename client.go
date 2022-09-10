package finviz

import (
	"github.com/mrod502/finviz/home"
	"github.com/mrod502/finviz/news"
)

type Client struct {
	*home.HomeClient
	*news.NewsClient
}

func New(opts *Options) (c *Client, err error) {
	c = &Client{
		HomeClient: home.NewClient(opts.cacheDuration),
		NewsClient: news.NewClient(opts.cacheDuration),
	}
	return
}
