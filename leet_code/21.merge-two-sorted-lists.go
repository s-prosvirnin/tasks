package leet_code

// https://leetcode.com/problems/merge-two-sorted-lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// фиктивный нулевой элемент, чтобы не было лишних ифчиков внутри цикла
	list3Cur := &ListNode{Val: -1000}
	list3 := list3Cur
	for {
		if list1 == nil && list2 == nil {
			break
		}

		if (list2 == nil) || (list1 != nil && list2 != nil && list1.Val <= list2.Val) {
			list3Cur.Next = list1
			list3Cur = list1
			list1 = list1.Next
			continue
		}
		if (list1 == nil) || (list1 != nil && list2 != nil && list2.Val < list1.Val) {
			list3Cur.Next = list2
			list3Cur = list2
			list2 = list2.Next
			continue
		}
		break
	}
	if list3.Next == nil {
		return nil
	} else {
		list3 = list3.Next
	}

	return list3
}
