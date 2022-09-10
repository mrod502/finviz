package news

import (
	"io"
	"strings"

	"github.com/mrod502/finviz/client"
	"github.com/mrod502/finviz/utils"
	"golang.org/x/net/html"
)

type Article struct {
	Title string
	Link  string
}

type NewsTable struct {
	Articles map[string]Article
}

type NewsClient struct {
	finviz   *client.Http[NewsTable]
	articles *client.Http[[]byte]
}

func (n *NewsClient) News() (*NewsTable, error) {
	return nil, nil
}
func Parse(r io.Reader, v *NewsTable) error {
	t := html.NewTokenizer(r)
	t.Next()
	for err := nextNn(t); err == nil; err = nextNn(t) {
		a, err := ParseArticle(t)
		if err != nil {
			return err
		}
		v.Articles[a.Link] = a
	}
	return nil
}

func ParseArticle(t *html.Tokenizer) (v Article, err error) {
	tok := t.Token()
	tt := tok.Type

	for {
		switch tt {
		case html.ErrorToken:
			err = utils.ErrNotFound
			return
		case html.EndTagToken:
			if t.Token().String() == "</tr>" {
				t.Next()
				return
			}
		default:
			attrs := getAttrs(tok)
			if val, ok := attrs["href"]; ok {
				v.Link = val
				v.Title = string(t.Text())

			}

		}
		tok = t.Token()
		tt = t.Next()
	}
}

func nextNn(t *html.Tokenizer) error {
	tok := t.Token()
	tt := tok.Type
	for {
		switch tt {
		case html.ErrorToken:
			return utils.ErrNotFound
		default:
			if strings.Contains(tok.String(), "<tr") {
				attrs := getAttrs(tok)
				if attrs["class"] == "nn" {
					return nil
				}
			}
		}
		tt = t.Next()
		tok = t.Token()
	}
}

func getAttrs(t html.Token) (a map[string]string) {
	a = make(map[string]string, 0)

	for _, v := range t.Attr {
		a[v.Key] = v.Val
	}
	return a
}
