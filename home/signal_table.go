package home

import (
	"errors"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

const (
	NameHomeTable = "t-home-table"
)

var (
	ErrNotFound error = errors.New("not found")
)

type Signal struct {
	Ticker string
	Last   float64
	Change float64
	Volume uint32
	Signal string
}

type SignalTable struct {
	Items []Signal
}

func (s *Signal) assignAttr(t *html.Tokenizer, ix int) (err error) {
	switch ix {
	case 1:
		s.Ticker = string(t.Text())
	case 2:
		var v float64
		v, err = strconv.ParseFloat(strings.ReplaceAll(string(t.Text()), "\n", ""), 64)
		s.Last = v
	case 3:
		var v float64
		val := string(t.Text())
		v, err = strconv.ParseFloat(val[:len(val)-1], 64)
		s.Change = v
	case 4:
		var v int64
		v, err = strconv.ParseInt(string(t.Text()), 10, 64)
		s.Volume = uint32(v)
	case 5:
		s.Signal = string(t.Text())
	}
	return err
}

func ParseSignal(t *html.Tokenizer) (s *Signal) {
	s = new(Signal)
	counter := 0
	for {
		tt := t.Next()
		switch tt {
		case html.TextToken:
			s.assignAttr(t, counter)
			counter++
		case html.EndTagToken:
			v, _ := t.TagName()
			if string(v) == "tr" {

				return
			}
		default:
		}
	}
}

func FindSignalTable(t *html.Tokenizer) error {
	for {
		tt := t.Next()

		switch tt {
		case html.ErrorToken:
			return ErrNotFound
		case html.StartTagToken:
			if b, _ := t.TagName(); string(b) == "table" {

				for {
					if _, v, more := t.TagAttr(); string(v) == NameHomeTable {
						return nil
					} else if !more {
						break
					}
				}
			}
		default:
		}
	}
}

func ParseSignalTable(t *html.Tokenizer) (s *SignalTable) {
	s = new(SignalTable)
	passedHeader := false
	secondTable := false
	for {
		tt := t.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.EndTagToken:
			if v, _ := t.TagName(); string(v) == "table" {
				if secondTable {
					return
				} else {
					secondTable = true
					passedHeader = false
				}
			}
		case html.StartTagToken:
			if v, _ := t.TagName(); string(v) == "tr" {
				if passedHeader {
					s.Items = append(s.Items, *ParseSignal(t))
				} else {
					passedHeader = true
				}
			}
		default:
		}
	}
}
