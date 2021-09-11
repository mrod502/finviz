package utils

import (
	"io"
	"net/http"
)

func GetReader(r *http.Request) (b io.ReadCloser, err error) {
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return
	}

	return res.Body, nil
}
func SetHeaders(r *http.Request) {
	r.Header.Set("accept", `text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9`)
	r.Header.Set("user-agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36`)

}
