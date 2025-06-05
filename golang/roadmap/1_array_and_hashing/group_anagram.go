package array_and_hashing

/*
Group Anagrams
Given an array of strings strs, group all anagrams together into sublists. You may return the output in any order.

An anagram is a string that contains the exact same characters as another string, but the order of the characters can be different.

Example 1:

Input: strs = ["act","pots","tops","cat","stop","hat"]

Output: [["hat"],["act", "cat"],["stop", "pots", "tops"]]
Example 2:

Input: strs = ["x"]

Output: [["x"]]
Example 3:

Input: strs = [""]

Output: [[""]]
Constraints:

1 <= strs.length <= 1000.
0 <= strs[i].length <= 100
strs[i] is made up of lowercase English letters.

*/

func GroupAnagrams(strs []string) [][]string {
	anagrams := make(map[[26]int][]string)

	for _, str := range strs {
		var count [26]int

		for _, letter := range str {
			count[letter-'a']++
		}

		anagrams[count] = append(anagrams[count], str)
	}

	result := make([][]string, 0)
	for _, anagram := range anagrams {
		result = append(result, anagram)
	}

	return result
}
