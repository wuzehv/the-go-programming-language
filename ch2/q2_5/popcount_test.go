package q2_5

import (
	"github.com/wuzehv/the_go_programming_language/ch2/q2_3"
	"github.com/wuzehv/the_go_programming_language/ch2/q2_4"
	"math"
	"testing"
)

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{0, 0},
		{1, 1},
		{255, 8},
		{1255, 7},
		{math.MaxUint64, 64},
	}

	for _, test := range tests {
		if got := q2_3.PopCount(test.input); got != test.want {
			t.Errorf("PopCount(%q) = %v", test.input, got)
		}

		if got := q2_3.PopCount2(test.input); got != test.want {
			t.Errorf("PopCount2(%q) = %v", test.input, got)
		}

		if got := q2_4.PopCount(test.input); got != test.want {
			t.Errorf("PopCount2(%q) = %v", test.input, got)
		}

		if got := PopCount(test.input); got != test.want {
			t.Errorf("PopCount2(%q) = %v", test.input, got)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q2_3.PopCount(math.MaxUint64)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q2_3.PopCount2(math.MaxUint64)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q2_4.PopCount(math.MaxUint64)
	}
}

func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(math.MaxUint64)
	}
}
