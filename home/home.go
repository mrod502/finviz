package home

import (
	"net/http"

	"github.com/mrod502/finviz/utils"
	"golang.org/x/net/html"
)

type Home struct {
	Signals SignalTable
}

func GetHome() (h *Home, err error) {
	h = new(Home)
	req, err := http.NewRequest(http.MethodGet, utils.BaseUri, nil)

	if err != nil {
		return nil, err
	}
	utils.SetHeaders(req)
	reader, err := utils.GetReader(req)

	if err != nil {
		return nil, err
	}

	t := html.NewTokenizer(reader)

	err = FindSignalTable(t)
	if err != nil {
		return nil, err
	}
	h.Signals = *ParseSignalTable(t)
	return
}
