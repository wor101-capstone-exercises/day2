package main

import "fmt"

// **Algorithm**
//   - Initialize variable to store index of the subString and set to 0
//   - Initialize variable to store number of matching characters and set to 0
//   - iterate over the originalString characters
//     - if the current originalString character equals the character in the subString at the current subStringIndex:
//       - increment the subStringIndex
//       - increment the number of matching characters
//     - if the number of mathcing characters equals the length of the substring, return true
//   - If you iterate through all characters, return false

func isSubsequence(s string, t string) bool {
	subStringIndex := 0
	matchCount := 0

	if len(s) == 0 {
		return true
	}

	for _, charRune := range t {
		char := string(charRune)
		if char == string(s[subStringIndex]) {
			subStringIndex++
			matchCount++
		}
		if matchCount == len(s) {
			return true
		}
	}
	return false
}

// josh
/*
   Input: Two strings named s and t
   Output: true if s is a subsequence of t, else false

   A subsequence is a new string that can be formed by removing characters
   from another string without changing the relative position of the remaining
   characters

   s and t consist only of lowercase english letters
   s and t may be empty strings

   initialize sp to 0
   initialize tp to 0

   while sp < len(s) and tp < len(t)
       if t[tp] == s[sp]
           increment sp and tp
       else
           increment tp
   return sp == len(s)

   Anchor/Runner Model
*/

func isSubsequence2(s string, t string) bool {
	sp, tp := 0, 0
	for sp < len(s) && tp < len(t) {
		if t[tp] == s[sp] {
			sp++
			tp++
		} else {
			tp++
		}
	}
	return sp == len(s)
}

// wes

/*
Assumptions:
- Arguments: two strings (s and t)
- Returns: bool
- return true iff s is a subsequence of t
- s is a subsequence of t just in case
    - s is formed from t by deleting 0 or more characters from t
    - without disturbing relative positions of remaining characters

Questions:
What should I do if len(s) == 0?
    - Presumed response:
        - return true ("" is, indeed, formed by deleting from t 0 or more characters without distruption)

Examples:

Input: s = "abc", t = "ahbgdc"
Output: true

Input: s = "axc", t = "ahbgdc"
Output: false

Algorithm:
- if len(s) == 0
    - return true
- if len(t) == 0
    - return false
- if s[0] == t[0]
    - return IsSubsequence(s[1:], t[1:])
- return IsSubsequence(s, t[1:])
*/

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	if s[0] == t[0] {
		return IsSubsequence(s[1:], t[1:])
	}
	return IsSubsequence(s, t[1:])
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc")) // true
	fmt.Println(isSubsequence("axc", "ahbgdc")) // false
	fmt.Println(isSubsequence("a", "b"))        // false
}
