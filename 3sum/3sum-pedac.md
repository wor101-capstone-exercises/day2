/*
**Problem**
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Constraints:

  0 <= nums.length <= 3000
  -105 <= nums[i] <= 105

Understanding the Problem

- input:
  - an array of integers

- output:
  - a 2dimensional array, with:
    - subarrays containing 3 ints that sum to 0 
    - subarrays must be unique

- model of problem:
  - USE THE 2SUM SOLUTION
  - set two pointers, start, and 2sumpointer
  - iterate over nums
    - for each num, set your target as the difference between 0 and the current num
    - use the 2sum solution to find the values you need from 2sumpointer to end of nums


**Examples / Test Cases**
Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:

Input: nums = []
Output: []
Example 3:

Input: nums = [0]
Output: []

**Data Structures**


**Algorithm**
*/