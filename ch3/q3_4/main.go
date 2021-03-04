package main

import (
	q3_2 "github.com/wuzehv/the_go_programming_language/ch3/q3_3"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		q3_2.Svg(w)
	})

	log.Fatal(http.ListenAndServe(":9099", nil))
}
