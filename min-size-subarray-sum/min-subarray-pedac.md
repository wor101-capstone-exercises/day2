/*
**Problem**
Given an array of positive integers nums and a positive integer target, return the minimal length of a contiguous subarray [numsl, numsl+1, ..., numsr-1, numsr] of which the sum is greater than or equal to target. If there is no such subarray, return 0 instead.

Understanding the Problem

- input:
  - integer for target number
  - array of integers

- output:
  - integer for length of shortest subarray

- model of problem
1. iterate over entire array
	- if any value is >= target, return 1
 - create variable called length and set to 2
2. set pointer s1 at 0 and pointer s2 at s1 + length - 1
3. loop while s1 < length of array - 1
3. use k-window slide method to see if any slice from s1 to s2 is greater than or equal than target
	- if yes, return length
	- if no, increase length by 1 and repeat

**Examples / Test Cases**


**Data Structures**


**Algorithm**
