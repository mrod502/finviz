package home

import (
	"io"
	"time"

	"github.com/mrod502/finviz/client"
	"github.com/mrod502/finviz/utils"
	"golang.org/x/net/html"
)

type Home struct {
	Signals SignalTable
}

func Parse(reader io.Reader, h *Home) (err error) {

	t := html.NewTokenizer(reader)

	err = FindSignalTable(t)
	if err != nil {
		return err
	}
	h.Signals = *ParseSignalTable(t)
	return
}

type HomeClient struct {
	*client.Http[Home]
}

func NewClient(cacheDuration time.Duration) *HomeClient {
	return &HomeClient{
		Http: client.NewHttp(Parse).WithCache(cacheDuration),
	}
}

func (c *HomeClient) Home() (h *Home, err error) {
	return c.Get(utils.BaseUri)
}
