/*
A good string is a string which doesn't have two adjacent characters s[i] and s[i + 1] where:

	0 <= i <= s.length - 2
	s[i] is a lower-case letter and s[i + 1] is the same letter but in upper-case or vice-versa.

Input: s = "leEeetcode"
Output: "leetcode"

Input: s = "abBAcC"
Output: ""
*/
package main

import (
	"fmt"
)

func makeGood(s string) string {
	// var array []string
	// var starIndex int
	for i, char := range s {
		fmt.Println(i)
		fmt.Println(char)
	}
}