package main

// **Algorithm**
//   - initialize two pointers: start = 0, end = length - 1
//   - when start + end is less than target, increment start
//   - when start + end is greater than target, decrement end
//   - when target is matched return array of [start + 1, end + 1]

import "fmt"

func twoSum(numbers []int, target int) []int {
	start := 0
	end := len(numbers) - 1

	for start < end {
		if numbers[start]+numbers[end] < target {
			start++
		} else if numbers[start]+numbers[end] > target {
			end--
		} else {
			return []int{start + 1, end + 1}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [1, 2]
	fmt.Println(twoSum([]int{2, 3, 4}, 6))      // [1, 3]
	fmt.Println(twoSum([]int{-1, 0}, -1))       // [1, 2]

}
