package search

import(
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		num_array		[]int
		givenKey		int
		expectedIndex	int
	}{
		{[]int{12, 27, 3, 42, 9}, 42, 3},
		{[]int{52, 15, 17}, 8, -1},
		{[]int{169, 1, 45, 47, 91, 214, 0, 112, 23}, 112, 7},
	}

	for _, test := range tests {
		if got := LinearSearch(test.num_array, test.givenKey); got != test.expectedIndex {
			t.Errorf("expected %q, got %q", test.expectedIndex, got)
		}
	}
}
