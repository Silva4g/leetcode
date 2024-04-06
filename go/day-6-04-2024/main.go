/*
Given a string s of '(' , ')' and lowercase English characters.

Your task is to remove the minimum number of parentheses ( '(' or ')',
in any positions ) so that the resulting
parentheses string is valid and return any valid string.

Formally, a parentheses string is valid if and only if:

It is the empty string, contains only lowercase characters, or
It can be written as AB (A concatenated with B), where A and B are valid strings, or
It can be written as (A), where A is a valid string.
*/

package main

import (
	"fmt"
)

func minRemoveToMakeValid(s string) string {
	balance := 0
	result := make([]rune, 0, len(s))
	for _, char := range s {
		if char == 40 {
			balance++
		} else if char == 41 {
			if balance == 0 {
					continue
			}
			balance--
		}
		result = append(result, char)
	}
	balance = 0
	finalResult := make([]rune, 0, len(result))
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] == 41 {
			balance++
		} else if result[i] == 40 {
			if balance == 0 {
				continue
			}
			balance--
		}
		finalResult = append(finalResult, result[i])
	}

	for i, j := 0, len(finalResult)-1; i < j; i, j = i+1, j-1 {
		finalResult[i], finalResult[j] = finalResult[j], finalResult[i]
	}

	return string(finalResult)
}

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)")) // "lee(t(c)o)de"
	fmt.Println(minRemoveToMakeValid("a)b(c)d")) // "ab(c)d"
	fmt.Println(minRemoveToMakeValid("))((")) // ""
}

/*
Resumo da solução:

1. Primeira passagem (da esquerda para a direita):
   - Construímos uma nova string (result) incluindo todos os caracteres da string original, exceto os parênteses fechados ')' que não possuem um parêntese aberto '(' correspondente.
   - Isso é feito mantendo um contador (balance) para parênteses abertos. Incrementamos para '(' e decrementamos para ')'. Se tentarmos decrementar quando balance é 0, significa que encontramos um ')' sem um '(' correspondente, e ele não é adicionado ao result.

2. Segunda passagem (da direita para a esquerda):
   - Invertendo o processo, agora removemos os parênteses abertos '(' extras, ou seja, aqueles sem um ')' correspondente.
   - Novamente, utilizamos um contador (balance), mas desta vez começamos do final da string resultante da primeira passagem. Incrementamos para ')' e decrementamos para '('. Se tentarmos decrementar quando balance é 0, significa que encontramos um '(' sem um ')' correspondente, e ele não é adicionado ao finalResult.

3. Como a segunda passagem é feita de trás para frente, o resultado final é invertido para manter a ordem original dos caracteres.

Essa abordagem assegura que todos os parênteses na string resultante estão balanceados, removendo o número mínimo necessário de parênteses para alcançar uma string válida, conforme definido pelo problema.
*/
