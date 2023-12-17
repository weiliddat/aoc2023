package day15

import (
	"testing"
)

var testInput = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
`

func TestPart01(t *testing.T) {
	expected := "1320"
	actual, err := Part01(testInput)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestPart02(t *testing.T) {
	expected := "145"
	actual, err := Part02(testInput)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestHash(t *testing.T) {
	expected := 52
	actual := hash("HASH")
	if expected != actual {
		t.Errorf("Expected %d got %d", expected, actual)
	}
}
