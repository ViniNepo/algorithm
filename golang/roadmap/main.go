package main

import (
	"fmt"
	datastruc "github.com/ViniNepo/algorithm/golang/date_structure/arrays_and_hashing"
	linked_list "github.com/ViniNepo/algorithm/golang/roadmap/6_linked_list"
)

func main() {
	//Array & Hashing

	// Contain duplicate - best form hashing
	//duplicatedArr := []int{1, 2, 3, 4, 4, 5, 6}
	//(result := )array_and_hashing.HasDuplicate(duplicatedArr))

	// Is anagram - best form hashing
	//anagramStr1, anagramStr2 := "racecar", "carrace"
	//println(array_and_hashing.IsAnagram(anagramStr1, anagramStr2))

	// Group anagram - best case hash table O(m * n)
	//groupAnagramList := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	//resultGroupAnagram := array_and_hashing.GroupAnagrams(groupAnagramList)
	//for _, r := range resultGroupAnagram {
	//	fmt.Println(r)
	//}

	// Top K frequent elements
	//result := array_and_hashing.TopKFrequent([]int{3, 3, 3, 1, 1, 2, 2, 2, 3}, 2)
	//fmt.Println(result)

	// Validate sudoku
	//sudoku := [][]string{
	//	[]string{"1", "2", ".", ".", "3", ".", ".", ".", "."},
	//	[]string{"4", ".", ".", "5", ".", ".", ".", ".", "."},
	//	[]string{".", "9", "8", ".", ".", ".", ".", ".", "3"},
	//	[]string{"5", ".", ".", ".", "6", ".", ".", ".", "4"},
	//	[]string{".", ".", ".", "8", ".", "3", ".", ".", "5"},
	//	[]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	//	[]string{".", ".", ".", ".", ".", ".", "2", ".", "."},
	//	[]string{".", ".", ".", "4", "1", "9", ".", ".", "8"},
	//	[]string{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	//}

	//sudokuWrong := [][]string{
	//	[]string{"1", "2", ".", ".", "3", ".", ".", ".", "."},
	//	[]string{"4", ".", ".", "5", ".", ".", ".", ".", "."},
	//	[]string{".", "9", "1", ".", ".", ".", ".", ".", "3"},
	//	[]string{"5", ".", ".", ".", "6", ".", ".", ".", "4"},
	//	[]string{".", ".", ".", "8", ".", "3", ".", ".", "5"},
	//	[]string{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	//	[]string{".", ".", ".", ".", ".", ".", "2", ".", "."},
	//	[]string{".", ".", ".", "4", "1", "9", ".", ".", "8"},
	//	[]string{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	//}
	//result := array_and_hashing.ValidateSudoku(sudoku)
	//fmt.Println(result)
	//resultWrong := array_and_hashing.ValidateSudoku(sudokuWrong)
	//fmt.Println(resultWrong)

	// Longest consecutive sequence
	//sequence := []int{2, 20, 4, 10, 3, 4, 5}
	//fmt.Println(array_and_hashing.LongestConsecutive(sequence))

	// Validate palindrome with question mark
	//str := "arte???"
	//str := "yh??tx"
	//result := array_and_hashing.ValidatePalindromeWithQuestions(str)
	//println(result)

	////////////////////////////////////////////////////////

	// Two Pointers

	// Validate Palindrome
	//strTrue := "Was it a car or a cat I saw?"
	//println(two_poiters.ValidPalindrome(strTrue))

	// Two Sum II
	//nums := []int{1, 2, 3, 4}
	//result := two_poiters.TwoSumII(nums, 3)
	//println(result)

	// Three sum
	//arr := []int{-1, 0, 1, 2, -1, -4}
	//result := two_poiters.ThreeSum(arr)
	//for _, value := range result {
	//	fmt.Println(value)
	//}

	// Container with most water
	//arr := []int{1, 7, 2, 5, 4, 7, 3, 6}
	//result := two_poiters.MaxArea(arr)
	//fmt.Println(result)

	// Trap water
	//height := []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1}
	//fmt.Println(two_poiters.Trap(height))

	////////////////////////////////////////////////////////

	// STACK

	// minimun stack
	//stk := datastruc.NewMinStack()
	//stk.Push(2)
	//stk.Push(1)
	//stk.Pop()
	//stk.Top()
	//stk.Push(1)
	//stk.Push(0)
	//println(stk.GetMin())
	//stk.Pop()
	//println(stk.GetMin())

	// validate parentheses
	//str := "([{}])"
	//result := stack.IsValid(str)
	//println(result)

	// Evaluate Reverse Polish Notation
	//tokens := []string{"1", "2", "+", "3", "*", "4", "-"}
	//result := stack.EvalRPN(tokens)
	//println(result)

	// Generate Parenthesis
	//pairs := 3
	//result := stack.GenerateParenthesis(pairs)
	//for _, res := range result {
	//	fmt.Println(res)
	//}

	// Daily temperatures
	//temperatures := []int{30, 38, 30, 36, 35, 40, 28}
	//result := stack.DailyTemperature(temperatures)
	//fmt.Println(result)

	// Car fleet
	//target := 10
	//positions, speed := []int{4, 1, 0, 7}, []int{2, 2, 1, 1}
	//result := stack.CarFleet(target, positions, speed)
	//fmt.Println(result)

	////////////////////////////////////////////////////////

	// Binary Search

	// Search 2D matrix
	//matrix := [][]int{[]int{1, 2, 4, 8}, []int{10, 11, 12, 13}, []int{14, 20, 30, 40}}
	//fmt.Println(binarysearch.SearchMatrix(matrix, 10))

	// Eating bananas
	//pile := []int{1, 4, 3, 2}
	//fmt.Println(binarysearch.MinEatingSpeed(pile, 9))

	// Find min value rotated array
	//arr := []int{4, 5, 0, 1, 2, 3}
	//fmt.Println(binarysearch.FindMinimumRotatedSortedKey(arr))

	// Search in sorted array
	//arr := []int{3, 4, 5, 6, 1, 2}
	//arr := []int{3, 5, 6, 0, 1, 2}
	//fmt.Println(binarysearch.SearchInRotatedSortedKey(arr, 4))

	////////////////////////////////////////////////////////

	// Sliding Window

	// Best time to buy and sell
	//prices := []int{7, 2, 5, 3, 6, 4, 1, 9}
	//fmt.Println(sliding_window.MaxProfit(prices))

	// Longest substring without repeating chars
	//str := "abcade"
	//fmt.Println(sliding_window.LongestSubstring(str))

	// Longest repeating substring with replacement
	//str := "AAABABB"
	//fmt.Println(sliding_window.CharacterReplacement(str, 2))

	// Permutation string
	//fmt.Println(sliding_window.CheckInclusion("abc", "lecabee"))

	////////////////////////////////////////////////////////

	// Linked list

	//Reverse linked List
	//node := datastruc.ListNode{Val: 0}
	//node.Next = &datastruc.ListNode{Val: 1}
	//node.Next.Next = &datastruc.ListNode{Val: 2}
	//node.Next.Next.Next = &datastruc.ListNode{Val: 3}
	//result := linked_list.ReverseLinkedList(&node)
	//for result != nil {
	//	fmt.Println(result.Val)
	//	result = result.Next
	//}

	// Merge two sorted linked list
	//node1 := datastruc.ListNode{Val: 1}
	//node1.Next = &datastruc.ListNode{Val: 2}
	//node1.Next.Next = &datastruc.ListNode{Val: 4}
	//
	//node2 := datastruc.ListNode{Val: 1}
	//node2.Next = &datastruc.ListNode{Val: 3}
	//node2.Next.Next = &datastruc.ListNode{Val: 5}
	//result := linked_list.MergeTwoLists(&node1, &node2)
	//for result != nil {
	//	fmt.Println(result.Val)
	//	result = result.Next
	//}

	// Reorder List
	//node1 := &datastruc.ListNode{Val: 0}
	//node1.Next = &datastruc.ListNode{Val: 1}
	//node1.Next.Next = &datastruc.ListNode{Val: 2}
	//node1.Next.Next.Next = &datastruc.ListNode{Val: 3}
	//node1.Next.Next.Next.Next = &datastruc.ListNode{Val: 4}
	//node1.Next.Next.Next.Next.Next = &datastruc.ListNode{Val: 5}
	//node1.Next.Next.Next.Next.Next.Next = &datastruc.ListNode{Val: 7}
	//linked_list.ReorderList(node1)

	// Remove Nht node from end of list
	node := &datastruc.ListNode{Val: 1}
	node.Next = &datastruc.ListNode{Val: 2}
	node.Next.Next = &datastruc.ListNode{Val: 3}
	node.Next.Next.Next = &datastruc.ListNode{Val: 4}
	node.Next.Next.Next.Next = &datastruc.ListNode{Val: 5}
	result := linked_list.RemoveNhtFromEnd(node, 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
