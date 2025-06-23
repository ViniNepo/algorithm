package sliding_window

/*
Longest Repeating Character Replacement
You are given a string s consisting of only uppercase english characters and an integer k. You can choose up to k characters of the string and replace them with any other uppercase English character.

After performing at most k replacements, return the length of the longest substring which contains only one distinct character.

Example 1:

Input: s = "XYYX", k = 2

Output: 4
Explanation: Either replace the 'X's with 'Y's, or replace the 'Y's with 'X's.

Example 2:

Input: s = "AAABABB", k = 1

Output: 5
*/

func CharacterReplacement(s string, k int) int {
	count := make(map[byte]int)
	res, l, maxFrequency := 0, 0, 0

	for r := 0; r < len(s); r++ {
		count[s[r]]++
		if count[s[r]] > maxFrequency {
			maxFrequency = count[s[r]]
		}

		for (r-l+1)-maxFrequency > k {
			count[s[l]]--
			l++
		}

		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}
