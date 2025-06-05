package array_and_hashing

/*
Contains Duplicate
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

Example 1:

Input: nums = [1, 2, 3, 3]

Output: true

Example 2:

Input: nums = [1, 2, 3, 4]

Output: false
*/

func HasDuplicate(nums []int) bool {
	hash := make(map[int]bool)

	for _, num := range nums {
		if hash[num] {
			return true
		}

		hash[num] = true
	}

	return false
}
