package fibonacci

func Fibonacci(n int) int {
	firstElement := 0
	secondElement := 1

	for nthTerm := 0; nthTerm < n; nthTerm++ {
		currentElement := firstElement
		firstElement = secondElement
		secondElement = firstElement + currentElement
	} 
	return firstElement
}
