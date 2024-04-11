/*
Given string num representing a non-negative integer num, and an integer k,
return the smallest possible integer after removing k digits from num.

Example 1:
Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219
which is the smallest.

Example 2:
Input: num = "10200", k = 1
Output: "200"
Explanation: Remove the leading 1 and the number is 200. Note that the output
must not contain leading zeroes.

Example 3:
Input: num = "10", k = 2
Output: "0"
Explanation: Remove all the digits from the number and it is left with nothing which is 0.
*/

package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stack := []byte{} // using a byte slice as a stack

	for i := range num {
		// while stack is not empty and the top of the stack is greater than the current number
		// and we still have k digits to remove
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i]) // push the current number onto the stack
	}

	// If k > 0, remove the remaining digits from the end
	stack = stack[:len(stack)-k]
	// Convert stack to string and remove leading zeros
	result := strings.TrimLeft(string(stack), "0")

	// If result is empty, return "0"
	if result == "" {
		return "0"
	}
	return result
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
}