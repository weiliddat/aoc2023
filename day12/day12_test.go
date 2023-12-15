package day12

import (
	"testing"
)

var testInput01 = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
`

func TestPart01(t *testing.T) {
	expected := "21"
	actual, err := Part01(testInput01)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

var testInput02 = ``

func TestPart02(t *testing.T) {
	expected := ""
	actual, err := Part02(testInput02)
	if err != nil {
		t.Error(err)
	}
	if expected != actual {
		t.Errorf("Expected %s got %s", expected, actual)
	}
}

func TestFit(t *testing.T) {
	testCases := []struct {
		record   string
		springs  []int
		expected int
	}{
		// {"???.###", []int{1, 1, 3}, 1},
		// {".??..??...?##.", []int{1, 1, 3}, 4},
		// {"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}, 1},
		// {"????.#...#...", []int{4, 1, 1}, 1},
		// {"????.######..#####.", []int{1, 6, 5}, 4},
		// {"?###????????", []int{3, 2, 1}, 10},
		// {"?###.????????#####??", []int{3, 12}, 3},
		// {"#??????#??.", []int{2, 7}, 1},
		{"?#??", []int{1}, 1},
		// {"???#.????#??", []int{1, 1, 1, 1}, 12},
		// {".?????#??#??", []int{1, 2, 1}, 8},
		// {"#.??#?#?????", []int{1, 2, 1, 1}, 4},
		// {"..????##?????.?????", []int{6, 2, 1}, 46},
		// {"#?##???#?#.??.?#?#", []int{1, 8, 2, 1, 1}, 1},
		// {"??????##??#??#???", []int{12, 1}, 4},
		// {"??..???#??", []int{1, 2, 2}, 2},
		// {"??##..????.?#.?.", []int{3, 1, 1, 1, 1}, 3},
	}
	for _, tC := range testCases {
		t.Run(tC.record, func(t *testing.T) {
			actual := fit(tC.record, tC.springs, "")
			if tC.expected != actual {
				t.Errorf("Expected %v got %v", tC.expected, actual)
			}
		})
	}
}
