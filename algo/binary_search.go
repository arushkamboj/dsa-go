package algo

import "fmt"

func BinarySearch() {
	fmt.Println("Implementing Binary Search in Go.")
	fmt.Println("=================================")

	inputArray := []int{4, 9, 13, 20, 30, 45, 67, 78}
	find := 78

	//==============================
	if ImplBinarySearch(inputArray, find) {
		fmt.Println("Found the element", find)
	} else {
		// low and high have crossed, element not present
		fmt.Println("Did not find element", find, "in the list")
	}
}

func ImplBinarySearch(list []int, find int) bool {
	// define mid = (low + high) / 2
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		if find == list[mid] {
			return true
		}

		if find < list[mid] {
			// element is not in the right subarray, ignore it
			high = mid - 1
		} else {
			// element is not in the left subarray, ignore it
			low = mid + 1
		}
	}
	return false
}
