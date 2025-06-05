package array_and_hashing

/*
Valid Anagram
Given two strings s and t, return true if the two strings are anagrams of each other, otherwise return false.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

Input: s = "racecar", t = "carrace"

Output: true
Example 2:

Input: s = "jar", t = "jam"

Output: false
Constraints:

s and t consist of lowercase English letters.
*/
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countS, countT := make(map[rune]int), make(map[rune]int)
	for i, letter := range s {
		countS[letter]++
		countT[rune(t[i])]++
	}

	for key, value := range countS {
		if countT[key] != value {
			return false
		}
	}

	return true
}
