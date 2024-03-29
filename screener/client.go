package screener

import (
	"strings"

	"github.com/mrod502/finviz/utils"
)

type Exchange string

const Uri = utils.BaseUri + "screener.ashx"

const (
	AMEX   Exchange = "exch_amex"
	NASDAQ Exchange = "exch_nasd"
	NYSE   Exchange = "exch_nyse"
)

type PriceTarget string

const (
	Above50 PriceTarget = "targetprice_a50"
	Below50 PriceTarget = "targetprice_b50"
)

type Filter struct {
	exchange Exchange
	target   PriceTarget
}

func (f Filter) String() string {
	var fs = make([]string, 0)
	if f.exchange != "" {
		fs = append(fs, string(f.exchange))
	}
	if f.target != "" {
		fs = append(fs, string(f.target))
	}
	val := strings.Join(fs, ",")
	if val != "" {
		val = "&f=" + val
	}
	return val
}

type Client struct {
	filter  Filter
	sorting Sorting
}

func New() *Client {
	return &Client{}
}

func (c *Client) Get() (*Table, error) {

	return &Table{}, nil
}
