package home

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"golang.org/x/net/html"
)

func TestSignalTable(t *testing.T) {
	b, _ := ioutil.ReadFile("signal-table.html")
	b = bytes.ReplaceAll(b, []byte("\n"), []byte(""))
	r := bytes.NewReader(b)

	z := html.NewTokenizer(r)
	tic := time.Now()
	out := ParseSignalTable(z)
	ts := time.Since(tic)
	fmt.Println(ts)
	fmt.Println(out)

}
