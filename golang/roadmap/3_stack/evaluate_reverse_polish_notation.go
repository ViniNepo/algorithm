package stack

import "strconv"

/*
 Evaluate Reverse Polish Notation
 You are given an array of strings tokens that represents a valid arithmetic expression in Reverse Polish Notation.

 Return the integer that represents the evaluation of the expression.

 The operands may be integers or the results of other operations.
 The operators include '+', '-', '*', and '/'.
 Assume that division between integers always truncates toward zero.
 Example 1:

 Input: tokens = ["1","2","+","3","*","4","-"]

 Output: 5

 Explanation: ((1 + 2) * 3) - 4 = 5
 Constraints:

 1 <= tokens.length <= 1000.
 tokens[i] is "+", "-", "*", or "/", or a string representing an integer in the range [-100, 100].

*/

func EvalRPN(tokens []string) int {
	stack := make([]int, 0)

	for _, token := range tokens {
		if token == "+" {
			val1, val2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, sum(val1, val2))
		} else if token == "-" {
			val1, val2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, sub(val1, val2))
		} else if token == "*" {
			val1, val2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, mul(val1, val2))
		} else if token == "/" {
			val1, val2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, div(val1, val2))
		} else {
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}

	return stack[0]
}

func sum(a, b int) int {
	return a + b
}

func div(a, b int) int {
	return a / b
}

func mul(a, b int) int {
	return a * b
}

func sub(a, b int) int {
	return a - b
}
