package main

import "fmt"

func sumArray(array []int) int {
	sum := 0

	for _, value := range array {
		sum += value
	}
	return sum
}

func kWindowSlide(target int, array []int, s1 int, s2 int) bool {
	sumOfSlice := sumArray(array[s1 : s2+1])

	for s2 < len(array) {
		if sumOfSlice >= target {
			return true
		} else if s2 >= len(array)-1 {
			return false
		}
		sumOfSlice -= array[s1]
		s1++
		s2++
		sumOfSlice += array[s2]
	}
	return false
}

func minSubArrayLen(target int, array []int) int {
	for _, value := range array {
		if value >= target {
			return 1
		}
	}

	length := 2

	for length <= len(array) {
		s1 := 0
		s2 := s1 + length - 1

		if kWindowSlide(target, array, s1, s2) {
			return length
		}
		length++
	}

	return 0
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))                                 // 2
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))                                          // 1
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))                          // 0
	fmt.Println(minSubArrayLen(80, []int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8})) // 6

}
