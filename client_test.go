package finviz

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {

	cli, err := New(NewOptions().WithCacheDuration(5 * time.Minute))

	if err != nil {
		t.Fatal(err)
	}
	home, err := cli.Home()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range home.Signals.Items {
		fmt.Printf("%+v\n", v)
	}

	tic := time.Now()

	homeCached, err := cli.Home()

	if err != nil {
		t.Fatal(err)
	}
	since := time.Since(tic)
	fmt.Println(since)

	for _, v := range homeCached.Signals.Items {
		fmt.Printf("%+v\n", v)
	}

}
