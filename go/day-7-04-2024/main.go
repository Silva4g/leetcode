/*
Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.

The following rules define a valid string:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".
*/

package main

import (
	"fmt"
)

func checkValidString(s string) bool {
	low, high := 0, 0
	for _, char := range s {
		if char == '(' {
			low++
			high++
		} else if char == ')' {
			if low > 0 {
				low--
			}
			high--
		} else {
			if low > 0 {
				low--
			}
			high++
		}

		if high < 0 {
			return false
		}
	}

    return low == 0
}

func main() {
	fmt.Println(checkValidString("()")) // true
	fmt.Println(checkValidString("(*)")) // true
	fmt.Println(checkValidString("(*))")) // true
	fmt.Println(checkValidString("(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())")) // false
}