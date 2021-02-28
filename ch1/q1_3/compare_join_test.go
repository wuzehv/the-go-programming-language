package main

import (
	"github.com/wuzehv/the_go_programming_language/ch1/q1_1"
	"strings"
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	d := appendData()
	for i := 0; i < b.N; i++ {
		q1_1.Echo(d)
	}
}

func BenchmarkJoin(b *testing.B) {
	d := appendData()
	for i := 0; i < b.N; i++ {
		strings.Join(d, " ")
	}
}
