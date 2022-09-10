package home

import (
	"fmt"
	"testing"
)

func TestGetHome(t *testing.T) {
	cli := NewClient(0)

	table, err := cli.Home()

	if err != nil {
		t.Fatal(err)
	}

	for _, v := range table.Signals.Items {
		fmt.Printf("%+v\n", v)
	}
}
