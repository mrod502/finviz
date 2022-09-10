package news

import (
	"io"
	"strings"
	"time"

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

func NewClient(cacheDuration time.Duration) *NewsClient {
	return &NewsClient{
		finviz:   client.NewHttp(Parse).WithCache(cacheDuration),
		articles: client.NewHttp(copyBytes).WithCache(cacheDuration),
	}
}

type NewsClient struct {
	finviz   *client.Http[NewsTable]
	articles *client.Http[[]byte]
}

func copyBytes(r io.Reader, v *[]byte) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	*v = b

	return nil
}

func (n *NewsClient) News() (*NewsTable, error) {
	return n.finviz.Get(`https://finviz.com/news.ashx`)
}

func (n *NewsClient) Article(url string) (*[]byte, error) {
	return n.articles.Get(url)
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
