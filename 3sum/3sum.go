package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for index, value := range nums {
		diff := target - value

		valueIndex, ok := seen[diff]

		if ok {
			return []int{value, valueIndex}
		} else {
			seen[value] = index
		}
	}

	return nil
}

// - model of problem:
//   - USE THE 2SUM SOLUTION
//   - set two pointers, start, and 2sumpointer
//   - iterate over nums
//     - for each num, set your target as the difference between 0 and the current num
//     - use the 2sum solution to find the values you need from 2sumpointer to end of nums
func threeSum(nums []int) [][]int {
	var triplets [][]int
	start := 0
	twoSumPointer := 1

	for start < len(nums)-3 {
		target := 0 - nums[start]
		result := twoSum(nums[twoSumPointer:], target)
		if result != nil {
			newTriplet := []int{nums[start]}
			newTriplet = append(newTriplet, result...)
			// magicallyCheckforDuplicateTriplet
			triplets = append(triplets, newTriplet)
		}
		if twoSumPointer < len(nums)-2 {
			twoSumPointer++
		} else {
			start++
		}
	}
	return triplets
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]

}
