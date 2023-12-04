package aoc_util_test

import (
	"aoc2023/aoc_util"
	"reflect"
	"testing"
)

func TestSplitLines(t *testing.T) {
	input := "a\nb\nc\n"

	expected := []string{
		"a",
		"b",
		"c",
	}

	actual := aoc_util.SplitLines(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v got %+v", expected, actual)
	}
}

func BenchmarkSplitLines(b *testing.B) {
	input := "a\nb\nc\n"

	for i := 0; i < b.N; i++ {
		aoc_util.SplitLines(input)
	}
}
