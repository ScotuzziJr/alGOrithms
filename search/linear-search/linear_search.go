// Linear search
// Time complexity O(n)
// Space complexity O(1)

// Problem:
// 	Given an array o n elements, find if a given key (i.e an specific element we're looking for) is contained
//	in the array

// Solution
//	1. Iterate through the array
//	2. Compare each element with the given key
//	3. If the element matches the key return the element index
//	4. If none of the elements matches the key return -1 

package main

import(
	"fmt"
)

func main()  {
	var num_array = []int {12, 27, 3, 42, 9} // we'll be working only with integers for simplicity
	const key int = 13
	var key_idx int = LinearSearch(num_array, key)

	if key_idx != -1{
		fmt.Printf("The element %d were found in index %d of the array.\n", key, key_idx)
	} else {
		fmt.Printf("The element %d were not found in the array.\n", key)
	}
}

// Classical implementation of linear search
func LinearSearch(num_array []int, key int) int {
	for idx, num := range num_array {
		if num == key {
			return idx
		}
	}

	return -1
}

// Alternative implementation of linear search with sort constraint
// If we add a sort constraint to the array we can stop the execution of the algorithm early by checking
//	if the current element we're looking for is bigger than the given key
// It's well known that one should implement a binary search algorithm instead of linear search for sorted
//	arrays but this implementation is only for study purposes
func LinearSearchSortConstraint(num_array []int, key int) int {
	for idx, num := range num_array {
		if num == key {
			return idx
		} else if num > key {
			break
		}
	}

	return -1
}
