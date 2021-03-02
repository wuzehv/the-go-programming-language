package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetchWithTimeout(url, 1*time.Second, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchWithTimeout(url string, timeout time.Duration, ch chan<- string) {
	start := time.Now()

	s := make(chan string)
	go func(url string) {
		resp, err := http.Get(url)
		if err != nil {
			s <- fmt.Sprint(err)
			return
		}

		nbytes, err := io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()

		if err != nil {
			s <- fmt.Sprint(err)
			return
		}

		s <- fmt.Sprintf("%7d %s", nbytes, url)
	}(url)

	select {
	case r := <-s:
		ch <- fmt.Sprintf("%.2fs %s", time.Since(start).Seconds(), r)
	case <-time.After(timeout):
		ch <- "timeout"
	}
}
