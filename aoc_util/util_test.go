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

func TestIntoColumns(t *testing.T) {
	input := []string{
		"123",
		"abc",
	}
	expected := []string{
		"1a",
		"2b",
		"3c",
	}

	actual := aoc_util.IntoColumns(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v got %+v", expected, actual)
	}
}

func BenchmarkIntoColumns(b *testing.B) {
	input := []string{
		"1234567890",
		"abcdefghij",
		"!@#$%^&*()",
	}

	for i := 0; i < b.N; i++ {
		aoc_util.IntoColumns(input)
	}
}

func TestSplitBlocks(t *testing.T) {
	input := `asdf

1234
`

	expected := []string{
		"asdf",
		"1234",
	}

	actual := aoc_util.SplitBlocks(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %+v got %+v", expected, actual)
	}
}
