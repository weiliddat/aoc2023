package day08_test

import (
	"aoc2023/day08"
	"testing"
)

var testInput01 = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`

var testInput02 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
`

func TestPart01(t *testing.T) {
	expected := "2"
	actual, _ := day08.Part01(testInput01)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}

	expected = "6"
	actual, _ = day08.Part01(testInput02)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	testInput := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

	expected := "6"
	actual, _ := day08.Part02(testInput)

	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func BenchmarkPart02(b *testing.B) {
	testInput := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`

	for i := 0; i < b.N; i++ {
		day08.Part02(testInput)
	}
}
