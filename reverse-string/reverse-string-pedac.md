/*
**Problem**
Write a function that reverses a string. The input string is given as an array of characters s.

You must do this by modifying the input array in-place with O(1) extra memory.

Constraints:

1 <= s.length <= 105
s[i] is a printable ascii character.

Understanding the Problem

- input:
  - an array of characters

- output:
  - the same array with characters reversed

- model of problem:
  - create two pointers: one at the start, one at the end
  - swap values at both pointers
  - increment start pointer, decrement end pointer until start passes end

**Examples / Test Cases**
Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]

**Data Structures**


**Algorithm**
 - Initalize pointer start with value of 0
 - initialize pointer end with value of length of input - 1
 - loop until start >= end
   - swap value at start with value at end
   - increment start pointer
   - decrement end pointer
- return array
*/