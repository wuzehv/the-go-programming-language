package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	Fetch()
}

func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(resp.StatusCode)
	}
}
