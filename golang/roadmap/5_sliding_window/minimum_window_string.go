package sliding_window

import "math"

/*
Minimum Window Substring
Given two strings s and t, return the shortest substring of s such that every character in t, including duplicates, is present in the substring. If such a substring does not exist, return an empty string "".

You may assume that the correct output is always unique.

Example 1:

Input: s = "OUZODYXAZV", t = "XYZ"

Output: "YXAZ"
Explanation: "YXAZ" is the shortest substring that includes "X", "Y", and "Z" from string t.

Example 2:

Input: s = "xyz", t = "xyz"

Output: "xyz"
Example 3:

Input: s = "x", t = "xy"

Output: ""
*/

func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}

	countT := make(map[rune]int)
	for _, c := range t {
		countT[c]++
	}

	have, need := 0, len(countT)
	res := []int{-1, -1}
	resLen := math.MaxInt32
	l := 0
	window := make(map[rune]int)

	for r := 0; r < len(s); r++ {
		c := rune(s[r])
		window[c]++

		if countT[c] > 0 && window[c] == countT[c] {
			have++
		}

		for have == need {
			if (r - l + 1) < resLen {
				res = []int{l, r}
				resLen = r - l + 1
			}

			window[rune(s[l])]--
			if countT[rune(s[l])] > 0 && window[rune(s[l])] < countT[rune(s[l])] {
				have--
			}
			l++
		}
	}

	if res[0] == -1 {
		return ""
	}
	return s[res[0] : res[1]+1]
}
