package finviz

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type endpoint string

const (
	BaseUri    endpoint = "https://finviz.com/"
	NewsUri    endpoint = BaseUri + "news.ashx"
	InsiderUri endpoint = BaseUri + "insidertrading.ashx"
	FuturesUri endpoint = BaseUri + "futures.ashx"
	ForexUri   endpoint = BaseUri + "forex.ashx"
	CryptoUri  endpoint = BaseUri + "crypto.ashx"
)

var headers map[string]string = map[string]string{
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"accept-encoding":           "gzip, deflate, br",
	"accept-language":           "en-US,en;q=0.9,ru-RU;q=0.8,ru;q=0.7",
	"referer":                   "https://finviz.com/",
	"sec-ch-ua":                 `Chromium";v="92", " Not A;Brand";v="99", "Google Chrome";v="92"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-fetch-dest":            "document",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-site":            "same-origin",
	"sec-fetch-user":            "?1",
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36",
	//"cookie":                    `insiderTradingUrl=insidertrading.ashx%3Ftc%3D7; pv_date=Thu Sep 02 2021 19:15:11 GMT-0400 (Eastern Daylight Time); screenerUrl=screener.ashx%3Fv%3D111; LoginRedirectFromUrl=https%3A%2F%2Ffinviz.com%2Fportfolio.ashx; pv_count=18`,
}

func setHeaders(r *http.Request) {
	for k, v := range headers {
		r.Header.Set(k, v)
	}
}

func (c Client) getPage(e endpoint) ([]byte, error) {
	req, err := http.NewRequest("GET", string(e), nil)
	if err != nil {
		return nil, err
	}
	setHeaders(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}

func parseFloat(b []byte) (float64, error) {
	return strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(string(b), "%", ""), ",", ""), 64)
}
