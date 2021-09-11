package home

import (
	"fmt"
	"testing"
)

func TestGetHome(t *testing.T) {

	table, err := GetHome()

	if err != nil {
		t.Fatal(err)
	}

	for _, v := range table.Signals.Items {
		fmt.Printf("%+v\n", v)
	}
}
