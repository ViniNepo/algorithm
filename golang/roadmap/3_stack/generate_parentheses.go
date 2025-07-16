package stack

import "strings"

/*
Generate Parentheses
You are given an integer n. Return all well-formed parentheses strings that you can generate with n pairs of parentheses.

Example 1:

Input: n = 1

Output: ["()"]
Example 2:

Input: n = 3

Output: ["((()))","(()())","(())()","()(())","()()()"]
You may return the answer in any order.

Constraints:

1 <= n <= 7
*/

func GenerateParenthesis(n int) []string {
	stack := make([]string, 0)
	result := make([]string, 0)

	var backtrack func(int, int)
	backtrack = func(openN, closedN int) {
		if openN == n && closedN == n {
			result = append(result, strings.Join(stack, ""))
			return
		}

		if openN < n {
			stack = append(stack, "(")
			backtrack(openN+1, closedN)
			stack = stack[:len(stack)-1]
		}

		if closedN < openN {
			stack = append(stack, ")")
			backtrack(openN, closedN+1)
			stack = stack[:len(stack)-1]
		}
	}

	backtrack(0, 0)
	return result
}

func GenerateParenthesis_2(n int) []string {
	return backtrack(n, 0, 0, []string{}, []string{})
}

func backtrack(
	n, openN, closedN int,
	stack []string,
	result []string,
) []string {
	if openN == n && closedN == n {
		return append(result, strings.Join(stack, ""))
	}

	if openN < n {
		result = backtrack(n, openN+1, closedN, append(stack, "("), result)
	}
	if closedN < openN {
		result = backtrack(n, openN, closedN+1, append(stack, ")"), result)
	}
	return result
}
