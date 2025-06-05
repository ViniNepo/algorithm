package stack

/*
Largest Rectangle In Histogram
You are given an array of integers heights where heights[i] represents the height of a bar. The width of each bar is 1.

Return the area of the largest rectangle that can be formed among the bars.

Note: This chart is known as a histogram.

Example 1:

Input: heights = [7,1,7,2,2,4]

Output: 8
Example 2:

Input: heights = [1,3,7]

Output: 7
*/

func LargestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([][2]int, 0)

	for i, h := range heights {
		start := i
		for len(stack) > 0 && stack[len(stack)-1][1] > h {
			index := stack[len(stack)-1][0]
			height := stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]
			area := height * (i - index)

			if area > maxArea {
				maxArea = area
			}
			start = index
		}
		stack = append(stack, [2]int{start, h})
	}

	n := len(heights)
	for _, pair := range stack {
		area := pair[1] * (n - pair[0])
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}
