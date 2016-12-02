// go test -v
// go test -bench=.

package stringer

import "testing"

var text = "abcad"
var right = [2]string{"bcada", "cadab"}
var wrong = [2]string{"bacda", "cbaad"}

// Testing both implementations for correctness

func TestCheckConcat(t *testing.T) {
	for _, elem := range right {
		if CheckConcat(text, elem) == false {
			t.Error("Test failed")
		}
	}
	for _, elem := range wrong {
		if CheckConcat(text, elem) == true {
			t.Error("Test failed")
		}
	}
}

func TestCheckDirect(t *testing.T) {
	for _, elem := range right {
		if CheckDirect(text, elem) == false {
			t.Error("Test failed")
		}
	}
	for _, elem := range wrong {
		if CheckDirect(text, elem) == true {
			t.Error("Test failed")
		}
	}
}

// Benchmarking to compare the direct approach with the concat-based one

var textBench = "abcadacacaca"

func BenchmarkCheckDirect(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckDirect(text, "acacaabcadac")
	}
}

func BenchmarkCheckConcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		CheckConcat(text, "acacaabcadac")
	}
}
