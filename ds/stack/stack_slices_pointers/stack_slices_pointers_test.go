package stackslicespointers

import(
	"testing"
	"reflect"
)

func TestStackSlicePointerPush(t *testing.T) {
	tests := []struct{
		stack 	  	  *[]int
		itemToPush 	  int
		updatedStack  *[]int
	}{
		{&[]int{}, 9, &[]int{9}},
		{&[]int{52, 7, 91}, 113, &[]int{52, 7, 91, 113}},
		{&[]int{21, 37, 43, 147}, 78, &[]int{21, 37, 43, 147, 78}},
	}

	for _, test := range tests {
		if Push(test.stack, test.itemToPush); !reflect.DeepEqual(test.stack, test.updatedStack) {
			t.Errorf("expected %v, got %v", test.updatedStack, test.stack)
		}
	}
}

func TestStackSlicePointerPop(t *testing.T) {
	tests := []struct{
		stack 	  	  *[]int
		updatedStack  *[]int
	}{
		{&[]int{12, 21, 15}, &[]int{12, 21}},
		{&[]int{52, 7, 91}, &[]int{52, 7}},
		{&[]int{21, 37, 43, 147}, &[]int{21, 37, 43}},
	}

	for _, test := range tests {
		if Pop(test.stack); !reflect.DeepEqual(test.stack, test.updatedStack) {
			t.Errorf("expected %v, got %v", test.updatedStack, test.stack)
		}
	}
}

func TestStackSlicePeek(t *testing.T) {
	tests := []struct{
		stack 	  	  *[]int
		topItem 	  int
	}{
		{&[]int{12, 21, 15}, 15},
		{&[]int{52, 7, 91}, 91},
		{&[]int{21, 37, 43, 147}, 147},
	}

	for _, test := range tests {
		if got := Peek(test.stack); !reflect.DeepEqual(got, test.topItem) {
			t.Errorf("expected %v, got %v", test.topItem, got)
		}
	}
}

func TestStackSliceIsEmpty(t *testing.T) {
	tests := []struct{
		stack 	  	  *[]int
		isEmpty 	  bool
	}{
		{&[]int{12, 21, 15}, false},
		{&[]int{}, true},
	}

	for _, test := range tests {
		if got := IsEmpty(test.stack); !reflect.DeepEqual(got, test.isEmpty) {
			t.Errorf("expected %v, got %v", test.isEmpty, got)
		}
	}
}
