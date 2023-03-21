package search

import(
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct{
		numArray 	  []int
		givenKey 	  int
		expectedIndex int
	}{
		{[]int{3, 7, 9, 12, 17}, 9, 2},
		{[]int{52, 78, 91}, 113, -1},
		{[]int{21, 37, 43, 78, 114, 215, 342}, 78, 3},
	}

	for _, test := range tests {
		if got := BinarySearch(test.numArray, test.givenKey); got != test.expectedIndex {
			t.Errorf("expected %v, got %v", test.expectedIndex, got)
		}
	}
}
