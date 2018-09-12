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
	carryVal := 0
	returnLst := &ListNode{Val: 0}
	currentNode := returnLst
	for l1 != nil || l2 != nil {
		p := l1
		if l1 == nil {
			p = &ListNode{Val: 0}
		}
		q := l2
		if l2 == nil {
			q = &ListNode{Val: 0}
		}
		sum := p.Val + q.Val + carryVal

		carryVal = sum / 10
		currentNode.Next = &ListNode{Val: sum % 10}
		currentNode = currentNode.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carryVal > 0 {
		currentNode.Next = &ListNode{Val: carryVal}
	}

	return returnLst.Next
}

func createLinkedList(nums []int) *ListNode {
	returnNode := &ListNode{}
	currentNode := returnNode
	for i, num := range nums {
		currentNode.Val = num
		if i < len(nums)-1 {
			currentNode.Next = &ListNode{}
			currentNode = currentNode.Next
		}
	}

	return returnNode
}

func main() {
	firstNode := createLinkedList([]int{5})
	secondNode := createLinkedList([]int{5})
	res := addTwoNumbers(firstNode, secondNode)
	for res != nil {
		fmt.Println(res)
		res = res.Next
	}
}
