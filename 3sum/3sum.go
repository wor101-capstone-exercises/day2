package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for index, value := range nums {
		diff := target - value

		_, ok := seen[diff]
		// fmt.Println("valueIndex:", valueIndex)
		// fmt.Println("Diff:", diff)

		if ok {
			return []int{value, diff}
		} else {
			seen[value] = index
		}
	}

	return nil
}

func allTwoSums(nums []int, target int) [][]int {
	seen := map[int]int{}
	var allTwoSums [][]int

	for index, value := range nums {
		diff := target - value

		_, ok := seen[diff]
		// fmt.Println("valueIndex:", valueIndex)
		// fmt.Println("Diff:", diff)

		if ok {
			allTwoSums = append(allTwoSums, []int{value, diff})
		} else {
			seen[value] = index
		}
	}

	if len(allTwoSums) > 0 {
		return allTwoSums
	} else {
		return nil
	}
}

// - model of problem:
//   - USE THE 2SUM SOLUTION
//   - set two pointers, start, and 2sumpointer
//   - iterate over nums
//     - for each num, set your target as the difference between 0 and the current num
//     - use the 2sum solution to find the values you need from 2sumpointer to end of nums

func isEqual(newTriplet []int, triplet []int) bool {
	for i, v := range triplet {
		if newTriplet[i] != v {
			return false
		}
	}
	return true
}

func magicallyCheckforNoDuplicateTriplet(newTriplet []int, triplets [][]int) bool {
	sort.Ints(newTriplet)

	for _, triplet := range triplets {
		if isEqual(newTriplet, triplet) {
			//fmt.Printf("Trip: %v is equal to %v", newTriplet, triplet)
			return false
		}
	}
	return true
}

func threeSum(nums []int) [][]int {
	var triplets [][]int
	start := 0
	twoSumPointer := 1

	for start <= len(nums)-3 {
		target := 0 - nums[start]
		//result := twoSum(nums[twoSumPointer:], target)
		results := allTwoSums(nums[twoSumPointer:], target)

		// fmt.Printf("Target: %v  StartValue: %v  Slice: %v  Result: %v \n", target, nums[start], nums[twoSumPointer:], results)
		// fmt.Println("All2SUMS:", allTwoSums(nums[twoSumPointer:], target))
		if results != nil && len(results) != 0 {
			for _, result := range results {
				newTriplet := []int{nums[start]}
				newTriplet = append(newTriplet, result...)
				if magicallyCheckforNoDuplicateTriplet(newTriplet, triplets) {
					triplets = append(triplets, newTriplet)
				}
			}

		}
		fmt.Printf("2SumPoint: %v  Len-2: %v \n", twoSumPointer, len(nums)-2)
		if twoSumPointer < len(nums)-2 {
			twoSumPointer++
		} else {
			start++
		}
	}
	return triplets
}

func main() {
	fmt.Println(threeSum([]int{0, -1, 1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) // [[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))      // [[-2,0,2],[-2,1,1]]

}
