package search

import(
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		numArray		[]int
		givenKey		int
		expectedIndex	int
	}{
		{[]int{12, 27, 3, 42, 9}, 42, 3},
		{[]int{52, 15, 17}, 8, -1},
		{[]int{169, 1, 45, 47, 91, 214, 0, 112, 23}, 112, 7},
	}

	for _, test := range tests {
		if got := LinearSearch(test.numArray, test.givenKey); got != test.expectedIndex {
			t.Errorf("expected %v, got %v", test.expectedIndex, got)
		}
	}
}
