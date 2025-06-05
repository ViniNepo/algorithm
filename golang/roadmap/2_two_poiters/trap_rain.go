package two_poiters

/*
You are given an array non-negative integers height which represent an elevation map. Each value height[i] represents the height of a bar, which has a width of 1.

Return the maximum area of water that can be trapped between the bars.

Example 1:

Input: height = [0,2,0,3,1,0,1,3,2,1]

Output: 9
Constraints:

1 <= height.length <= 1000
0 <= height[i] <= 1000
*/

func Trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := height[left], height[right]
	res := 0

	for left < right {
		if maxLeft < maxRight {
			left++
			maxLeft = max(maxLeft, height[left])
			res += maxLeft - height[left]
		} else {
			right--
			maxRight = max(maxRight, height[right])
			res += maxRight - height[right]
		}
	}

	return res
}
