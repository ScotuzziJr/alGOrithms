package fibonacci

import(
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct{
		nthTerm 	  int
		expectedValue int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
	}

	for _, test := range tests {
		if got := Fibonacci(test.nthTerm); got != test.expectedValue {
			t.Errorf("expected %v, got %v", test.expectedValue, got)
		}
	}
}
