package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for index, value := range nums {
		diff := target - value

		valueIndex, ok := seen[diff]

		if ok {
			return []int{valueIndex, index}
		} else {
			seen[value] = index
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{0, -1, 1}, 0))
}
