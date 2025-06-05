package stack

/*
Valid Parentheses
Solved
You are given a string s consisting of the following characters: '(', ')', '{', '}', '[' and ']'.

The input string s is valid if and only if:

Every open bracket is closed by the same type of close bracket.
Open brackets are closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
Return true if s is a valid string, and false otherwise.

Example 1:

Input: s = "[]"

Output: true
Example 2:

Input: s = "([{}])"

Output: true
Example 3:

Input: s = "[(])"

Output: false
Explanation: The brackets are not closed in the correct order.

Constraints:

1 <= s.length <= 1000
*/

func IsValid(s string) bool {
	var stack []rune
	parentheses := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, letter := range s {
		if l, found := parentheses[letter]; found {
			if stack[len(stack)-1] == l {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, letter)
		}
	}

	return true
}
