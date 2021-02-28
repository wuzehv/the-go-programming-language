package main

import (
	"bufio"
	"fmt"
	"os"
)

/*func main() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)

	for l, n := range counts {
		fmt.Printf("%d -- %v\n", n, l)
	}
}
*/

func main() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, f := range files {
			fd, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "open file error: %v\n", err)
				continue
			}

			countLines(fd, counts)
			fd.Close()
		}
	}

	for l, n := range counts {
		fmt.Printf("%d -- %v\n", n, l)
	}
}

func countLines(f *os.File, counts map[string]int) {
	s := bufio.NewScanner(f)
	for s.Scan() {
		counts[s.Text()]++
	}
}
