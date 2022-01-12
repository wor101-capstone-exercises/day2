/*
**Problem**
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Constraints:

0 <= s.length <= 100
0 <= t.length <= 104
s and t consist only of lowercase English letters.

Understanding the Problem

- input:
  - a 'substring' string
  - a 'originalString' string

- output:
  - boolean

- model of problem:
  - iteraete over the originalString characters
  - keep track of index for substring starting at 0
  - if character in originalString matches substring at current index, increase substring index by 1
  - if index for substring equals length - 1 of substring, return true
  - otherwise return false


**Examples / Test Cases**
Example 1:

Input: s = "abc", t = "ahbgdc"
                       [a, h, b, g, d, c]
Output: true

Example 2:

Input: s = "axc", t = "ahbgdc"
Output: false

**Data Structures**


**Algorithm**
  - Initialize variable to store index of the subString and set to 0
  - Initialize variable to store number of matching characters and set to 0
  - iterate over the originalString characters
    - if the current originalString character equals the character in the subString at the current subStringIndex:
      - increment the subStringIndex
      - increment the number of matching characters
    - if the number of mathcing characters equals the length of the substring, return true
  - If you iterate through all characters, return false
*/