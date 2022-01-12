package main

import "fmt"

func reverseString(s []byte) {
	start := 0
	end := len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func main() {
	byteArray := []byte("hello")
	hannaArray := []byte("Hannah")
	reverseString(byteArray)
	reverseString(hannaArray)
	fmt.Println(string(byteArray))
	fmt.Println(string(hannaArray))
}
