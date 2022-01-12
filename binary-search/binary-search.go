package main

import "fmt"

func binarySearch(array []int, searchValue int) int {
	middleIndex := len(array) / 2

	fmt.Println(array)
	if array[middleIndex] > searchValue {
		return binarySearch(array[:middleIndex], searchValue)
	} else if array[middleIndex] < searchValue {
		return binarySearch(array[middleIndex+1:], searchValue)
	} else {
		return array[middleIndex]
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12}, 1))

}
