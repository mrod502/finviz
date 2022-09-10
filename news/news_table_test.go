package news

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestParseArticle(t *testing.T) {
	f, err := os.Open("nn.html")
	if err != nil {
		t.Fatal(err)
	}
	tok := html.NewTokenizer(f)
	tok.Next()
	v0, err := ParseArticle(tok)
	if err != nil {
		t.Fatal(err)
	}
	v1, err := ParseArticle(tok)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", v0)
	fmt.Printf("%+v\n", v1)
}

func TestParseTable(t *testing.T) {

	f, err := os.Open("news.html")
	if err != nil {
		t.Fatal(err)
	}
	nt := NewsTable{Articles: make(map[string]Article)}
	if err := Parse(f, &nt); err != nil {
		t.Fatal(err)
	}

}
