package two_poiters

import (
	"strings"
	"unicode"
)

/*
 Valid Palindrome
 Given a string s, return true if it is a palindrome, otherwise return false.

 A palindrome is a string that reads the same forward and backward. It is also case-insensitive and ignores all non-alphanumeric characters.

 Example 1:

 Input: s = "Was it a car or a cat I saw?"

 Output: true
 Explanation: After considering only alphanumerical characters we have "wasitacaroracatisaw", which is a palindrome.

 Example 2:

 Input: s = "tab a cat"

 Output: false
 Explanation: "tabacat" is not a palindrome.

 Constraints:

 1 <= s.length <= 1000
 s is made up of only printable ASCII characters.
*/

func ValidPalindrome(s string) bool {
	formatedStr := ""
	for _, str := range strings.ToLower(s) {
		if str >= 'a' && str <= 'z' {
			formatedStr += string(str)
		}
	}

	left, right := 0, len(formatedStr)-1

	for left < right {
		if formatedStr[left] != formatedStr[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNum(rune(s[l])) {
			l++
		}
		for r > l && !isAlphaNum(rune(s[r])) {
			r--
		}
		if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlphaNum(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}
