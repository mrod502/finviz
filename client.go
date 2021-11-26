package finviz

import (
	"time"

	"github.com/mrod502/finviz/home"
	"github.com/mrod502/finviz/screener"
	gocache "github.com/mrod502/go-cache"
)

type Client struct {
	endpoints *gocache.InterfaceCache
}

func (c *Client) Home() (h *home.Home, err error) {
	if c.endpoints.Exists("home") {
		v, err := c.endpoints.Get("home")
		if err != nil {
			return nil, err
		}
		return v.(*home.Home), nil
	}
	v, err := home.GetHome()
	if err != nil {
		return nil, err
	}
	c.endpoints.Set("home", v)
	return v, nil
}

func (c *Client) News() {}

func (c *Client) Screener() (*screener.Table, error) {
	return &screener.Table{}, nil
}

func (c Client) Forex() {}

func (c Client) Crypto() {}

func NewClient(opts *Options) (c *Client, err error) {
	if opts.cacheDuration < 0 {
		opts.cacheDuration = 3 * time.Minute
	}
	c = &Client{
		endpoints: gocache.NewInterfaceCache().WithExpiration(opts.cacheDuration),
	}

	return
}
