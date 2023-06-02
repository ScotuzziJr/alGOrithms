package stackslicespointers

// make a new empty stack
func StackBuilder() []int {
	var stack []int
  
  return stack
}

// check if the stack is empty
func IsEmpty(stack *[]int) bool {
  return len(*stack) == 0 
}

// add a new item to the top (the end) of stack
func Push(stack *[]int, item int) {
  // append returns a new slice, that's why we need to assign it to our stack
  *stack = append(*stack, item)
}

// remove the top item (the last item pushed) of the stack
func Pop(stack *[]int) {
  var stackSize int = len(*stack) -1
  *stack = (*stack)[:stackSize] // reassign the stack to itself but left out the last element
}

// prints the top element of the stack (but doesn't remove it)
func Peek(stack *[]int) int {
  var stackSize int = len(*stack) - 1
  return (*stack)[stackSize]
}
