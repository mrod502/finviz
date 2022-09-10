package utils

import (
	"errors"
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

var (
	ErrNotFound error = errors.New("element not found")
)
