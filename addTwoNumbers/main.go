package main

import "fmt"

/**
* You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

/**
 * Definition for singly-linked list.
 **/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// carryVal := 0
	returnLst := ListNode{}
	// currentNode := returnLst
	// for *l1.Next != nil && *l2.Next != nil {
	// 	currentNode.Val = (l1.Val + l2.Val) + carryVal
	// 	currentNode.Next = ListNode{}
	// 	carryVal = l1.Val % l2.Val

	// 	l1 = *l1.Next
	// 	l2 = *l2.Next
	// 	currentNode = *currentNode.Next
	// }

	return &returnLst
}

func createLinkedList(nums []int) *ListNode {
	returnNode := &ListNode{}
	currentNode := returnNode
	for _, num := range nums {
		currentNode.Val = num
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}

	return returnNode
}

func main() {
	firstNode := createLinkedList([]int{1, 3, 2, 5, 4, 6})
	fmt.Println(firstNode)
}
