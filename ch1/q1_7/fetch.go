package q1_7

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Fetch(dst io.Writer, f os.File) {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(dst, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: url: %s, error: %v\n", url, err)
		}
	}
}
