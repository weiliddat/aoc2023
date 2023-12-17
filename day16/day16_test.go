package day16

import (
	"testing"
)

var testInput = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
`

func TestPart01(t *testing.T) {
	expected := "46"
	actual, err := Part01(testInput)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}

	expected = "7608"
	actual, _ = Part01(Input)
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func BenchmarkPart01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Part01(Input)
	}
}

func TestPart02(t *testing.T) {
	expected := "51"
	actual, err := Part02(testInput)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}
