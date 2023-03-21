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

package search


// Classical implementation of linear search
func LinearSearch(numArray []int, key int) int {
	for idx, num := range numArray {
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
func LinearSearchSortConstraint(numArray []int, key int) int {
	for idx, num := range numArray {
		if num == key {
			return idx
		} else if num > key {
			break
		}
	}

	return -1
}
