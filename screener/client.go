package screener

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mrod502/finviz/constants"
	"golang.org/x/net/html"
)

type Exchange string

const Uri = constants.BaseUri + "screener.ashx"

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

func NewClient() *Client {
	return &Client{}
}

func (c *Client) Get() (*Table, error) {

	return &Table{}, nil
}

func buildRequest(f Filter, s Sorting, page uint) (r *http.Request, err error) {
	var uri string = Uri

	r, err = http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return
	}

	reader, err := getReader(r)
	if err != nil {
		return
	}

	t := html.NewTokenizer(reader)

	for {
		fmt.Println(t)
		break
	}

	return
}

func getReader(r *http.Request) (b io.ReadCloser, err error) {
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return
	}

	return res.Body, nil
}
func setHeaders(r *http.Request) {
	r.Header.Set("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`)

}
