package leet_code

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode.com/problems/add-two-numbers/
2. Add Two Numbers

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln1 := l1
	ln2 := l2
	var ln3 *ListNode
	var listNodeStart *ListNode
	remains := 0
	for {
		val1 := 0
		val2 := 0
		if ln1 != nil {
			val1 = ln1.Val
		}
		if ln2 != nil {
			val2 = ln2.Val
		}
		val := val1 + val2 + remains
		remains = 0
		if val > 9 {
			remains = val / 10
			val = val % 10
		}

		prevNode := ln3
		l := &ListNode{val, nil}
		ln3 = l
		if prevNode != nil {
			prevNode.Next = ln3
		}
		if listNodeStart == nil {
			listNodeStart = ln3
		}

		if ln1 != nil && ln1.Next != nil {
			ln1 = ln1.Next
		} else {
			ln1 = nil
		}
		if ln2 != nil && ln2.Next != nil {
			ln2 = ln2.Next
		} else {
			ln2 = nil
		}
		if ln1 == nil && ln2 == nil && remains == 0 {
			break
		}
	}

	return listNodeStart
}
