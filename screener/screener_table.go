package screener

import "golang.org/x/net/html"

type Table struct {
	Entries []*Entry
}

var titles = map[string]bool{
	"No.":        true,
	"Ticker":     true,
	"Company":    true,
	"Sector":     true,
	"Industry":   true,
	"Country":    true,
	"Market Cap": true,
	"P/E":        true,
	"Price":      true,
	"Change":     true,
	"Volume":     true,
}

type Parser struct {
}

func (p *Parser) Parse(t *html.Tokenizer) (v *Table, err error) {

	return
}

type Entry struct {
	No        uint32
	Ticker    string
	Company   string
	Sector    string
	Industry  string
	Country   string
	MarketCap string
	PE        float64
	Price     float64
	Change    float64
	Volume    uint32
}
