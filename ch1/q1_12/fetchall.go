package main

import (
	"fmt"
	"github.com/wuzehv/the_go_programming_language/ch1/q1_5"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		var cycles int = 5
		if _, ok := r.Form["cycles"]; ok {
			if len(r.Form["cycles"]) > 0 {
				cycles, _ = strconv.Atoi(r.Form["cycles"][0])
			}
		}

		q1_5.Lissajous(w, float64(cycles))
	})

	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "count: %d\n", count)
		mu.Unlock()
	})

	log.Fatal(http.ListenAndServe(":8099", nil))
}
