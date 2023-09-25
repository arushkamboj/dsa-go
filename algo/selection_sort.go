package algo

import (
	"fmt"
	"math/rand"
)

func SelectionSort() {
	// create a slice of random length and random elements
	len := 10
	list := make([]int, len)
	for i := 0; i < len-1; i++ {
		list[i] = rand.Intn(100)
	}
	fmt.Println("Input list before calling Selection Sort...")
	fmt.Println(list)

	// call the sortingAlgo on the list and print the results
	sortedList := ImplSelectionSort(list)
	fmt.Println("Sorted List")
	fmt.Println(sortedList)
}

func ImplSelectionSort(list []int) []int {
	var minIndex int
	for i := 0; i < len(list)-1; i++ {
		minIndex = i
		for j := i + 1; j < len(list); j++ {
			if list[minIndex] > list[j] {
				minIndex = j
			}
		}
		// swap list[i] with the minimum element from the right subarray
		temp := list[i]
		list[i] = list[minIndex]
		list[minIndex] = temp
	}
	return list
}
