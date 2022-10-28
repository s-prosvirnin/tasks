package leet_code

import "unsafe"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode.com/problems/linked-list-cycle/
141. Linked List Cycle

Given head, the head of a linked list, determine if the linked list has a cycle in it.
There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
Return true if there is a cycle in the linked list. Otherwise, return false.
*/
func hasCycle(head *ListNode) bool {
	return hasCycleByHaunting(head)
}

//также можно решить через мапу, ключами которой будут указатели. если при итерации указатель уже есть в мапе, то это закольцованный список

/**
2 бегущих указателя. один (убегающий) итерируется на следующий элемент, второй (догоняющий) через элемент
т.к. скорость догоняющего выше убегающего рано или поздно догоняющий догонит убегающего
*/
func hasCycleByHaunting(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	haunting := &(*head)
	victim := &(*head)
	for {
		if haunting.Next == nil || haunting.Next.Next == nil {
			return false
		}
		if victim.Next == nil {
			return false
		}
		haunting = haunting.Next.Next
		victim = victim.Next
		if haunting == victim {
			return true
		}
	}
}

// сомнительное решение через утверждение о последовательном выделении указателей
func hasCycleByPointersAddreses(head *ListNode) bool {
	for {
		if head == nil {
			return false
		}
		if head.Next == nil {
			return false
		}
		next := head.Next
		if uintptr(unsafe.Pointer(next)) <= uintptr(unsafe.Pointer(head)) {
			return true
		}
		head = next
	}
}
