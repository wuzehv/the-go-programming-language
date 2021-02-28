package main

import (
	"fmt"
	"github.com/wuzehv/the_go_programming_language/ch1/q1_1"
	"strings"
	"time"
)

var s = "hello, 世界!"

func appendData() []string {
	x := 10000
	var input = make([]string, 0, len(s)*x)
	for i := 0; i < x; i++ {
		input = append(input, s)
	}

	return input
}

func main() {
	d := appendData()

	s := time.Now()
	q1_1.Echo(d)
	e := time.Since(s)
	fmt.Printf("echo cost = %v\n", e)

	s = time.Now()
	strings.Join(d, " ")
	e = time.Since(s)
	fmt.Printf("echo cost = %v\n", e)
}
