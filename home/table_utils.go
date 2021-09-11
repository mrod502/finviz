package home

import "golang.org/x/net/html"

type TableParser struct {
	columns []string
	v       map[string]string
}

func NewTableParser(cols []string) *TableParser {
	t := &TableParser{
		columns: cols,
		v:       make(map[string]string),
	}

	return t
}

func (p *TableParser) Parse(t *html.Tokenizer) error {

	return nil
}
