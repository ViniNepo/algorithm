package stack

/*
Daily Temperatures
You are given an array of integers temperatures where temperatures[i] represents the daily temperatures on the ith day.

Return an array result where result[i] is the number of days after the ith day before a warmer temperature appears on a future day. If there is no day in the future where a warmer temperature will appear for the ith day, set result[i] to 0 instead.

Example 1:

Input: temperatures = [30,38,30,36,35,40,28]

Output: [1,4,1,2,1,0,0]
Example 2:

Input: temperatures = [22,21,20]

Output: [0,0,0]
*/
func DailyTemperature(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := []int{}

	for i, t := range temperatures {
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			tt := temperatures[stack[len(stack)-1]]
			stackInd := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[stackInd] = i - stackInd
			println(tt)
		}
		stack = append(stack, i)
	}

	return res
}
