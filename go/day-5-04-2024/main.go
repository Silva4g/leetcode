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

func isBadPair(a, b rune) bool {
	return a != b && (a == b+32 || a+32 == b)
}

func makeGood(s string) string {
	stack := []rune{}
	for _, char := range s {
		if len(stack) > 0 && isBadPair(stack[len(stack)-1], char) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
	fmt.Println(makeGood("s"))

}

/*
para esta solucao eu crio uma pilha chamada stack e para sempre percorrer ela usamos o len(stack) o tamanho da pilha -1,
apos isso pegamos o char que printa o bite do caractere do momento, se o bite tiver mais 32 bites do que a letra anterior
nos removemos a letra anterior (se for maiuscula tem menos bites, se for minuscula tem mais 32 bites) e assim se nao 
tiver adicionamos na pilha 
*/