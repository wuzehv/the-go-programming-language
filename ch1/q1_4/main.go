package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dup2()
}

func dup() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)

	for l, n := range counts {
		fmt.Printf("%d -- %v\n", n, l)
	}
}

func dup2() {
	counts := make(map[string]int)

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

func dup3() {
	counts := make(map[string]int)
	for _, f := range os.Args[1:] {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dump3 :%v\n", err)
			continue
		}

		for _, line := range strings.Split(string(content), "\n") {
			counts[line]++
		}

		for k, v := range counts {
			fmt.Printf("%v -- %v\n", v, k)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	s := bufio.NewScanner(f)
	for s.Scan() {
		counts[s.Text()]++
	}
}
