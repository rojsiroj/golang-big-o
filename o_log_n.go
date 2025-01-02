package big_o

/*
 * binarySearch is a function that has O(log n) time complexity
 */

func binarySearch(arr []int, value int) int {
	start := 0
	end := len(arr) - 1
	middle := (start + end) / 2

	for arr[middle] != value && start <= end {
		if value < arr[middle] {
			end = middle - 1
		} else {
			start = middle + 1
		}
		middle = (start + end) / 2
	}

	if arr[middle] == value {
		return middle
	}
	return -1
}
