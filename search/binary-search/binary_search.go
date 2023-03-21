// Binary search
// Time complexity O(log n)
// Space complexity O(1)

// Problem:
// 	Given an sorted array o n elements, find if a given key (i.e an specific element we're looking for) is
//	contained in the array

// Solution
//	1. Get the middle element of the array (idx_middle = (idx_first + idx_last) / 2)
//	2. Compare the middle element with the given key
//	3. If the element matches the key return the element index
//	4. If the key is greater than the given key, discard the lower half of the array
// 	5. If the key is lesser than the given key, discard the superior half of the array
//	6. Do it until remains a single element in the array (i.e while the first element index is less or equal
//		the last element index)
//	7. Return -1 if the element were not found 

package search

func BinarySearch(numArray []int, key int) int {
	idxFirstElement := 0
	idxLastElement := len(numArray) - 1
	var idxMiddleElement int

	for idxFirstElement <= idxLastElement {
		idxMiddleElement = (idxFirstElement + idxLastElement) / 2

		if numArray[idxMiddleElement] > key {
			idxLastElement = idxMiddleElement - 1
		} else if numArray[idxMiddleElement] < key {
			idxFirstElement = idxMiddleElement + 1
		} else {
			return idxMiddleElement
		} 
	}
	return -1
}
