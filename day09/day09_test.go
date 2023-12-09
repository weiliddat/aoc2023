package day09_test

import (
	"aoc2023/day09"
	"testing"
)

var testInput = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestPart01(t *testing.T) {
	expected := "114"
	actual, _ := day09.Part01(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := "2"
	actual, _ := day09.Part02(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func BenchmarkPart01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day09.Part01(day09.Input)
	}
}

func BenchmarkPart01Lagrange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day09.Part01Lagrange(day09.Input)
	}
}

func TestLagrange(t *testing.T) {
	testInput := []int{0, 3, 6, 9, 12, 15}
	expected := 18
	actual := day09.Lagrange(testInput)

	if expected != actual {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	testInput = []int{1, 3, 6, 10, 15, 21}
	expected = 28
	actual = day09.Lagrange(testInput)

	if expected != actual {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}

	testInput = []int{10, 13, 16, 21, 30, 45}
	expected = 68
	actual = day09.Lagrange(testInput)

	if expected != actual {
		t.Errorf("Expected %#v got %#v", expected, actual)
	}
}
