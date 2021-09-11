package finviz

import (
	"github.com/mrod502/finviz/home"
	"github.com/mrod502/finviz/screener"
	gocache "github.com/mrod502/go-cache"
)

type Client struct {
	endpoints *gocache.InterfaceCache
}

func (c Client) Home() (h *home.Home, err error) {

	b, err := c.getPage(BaseUri)
	if err != nil {
		return nil, err
	}

	return home.ParseHome(b)
}
func (c Client) News() {}
func (c *Client) Screener() (*screener.Table, error) {

	return &screener.Table{}, nil
}
func (c Client) Forex()  {}
func (c Client) Crypto() {}

func NewClient() (c *Client, err error) {

	c = &Client{}

	return
}
