package leet_code

// Итеративно меняем ссылку следующего узла на предыдущий узел

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Итеративно разворачиваем текущий список.
// Меняем ссылку следующего узла на предыдущий узел.
func reverseList(head *ListNode) *ListNode {
	// Пустой список легче обработать сразу
	if head == nil {
		return nil
	}

	// Текущий обрабатываемый узел.
	cur := head.Next
	// Предыдущий узел.
	prev := head
	// Начинаем сразу со второго узла в качестве текущего,
	// поэтому зануляем у предыдущего (корневого) узла ссылку на второй узел.
	prev.Next = nil
	// Итерируемся, пока не достигнем конца списка.
	for cur != nil {
		// Нам нужен следующий узел, чтобы потом на него сдвинуться.
		next := cur.Next
		// У текущего узла меняем ссылку на предыдущий.
		cur.Next = prev
		// Смещаем предыдущий узел на текущий.
		prev = cur
		// Смещаем текущий узел на следующий.
		cur = next
	}

	return prev
}

// Итеративно строим новый список

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Итеративно разворачиваем текущий список и строим новый.
// Тратим лишнее место на новый список. Можно переворачивать список на лету.
func reverseList1(head *ListNode) *ListNode {
	// Пустой список легче обработать сразу
	if head == nil {
		return nil
	}

	// Текущий узел из старого списка (старый узел).
	oldCurNode := head
	// Текущий узел из нового списка (новый узел).
	var newCurNode *ListNode
	for {
		// Создаем текущий новый узел на основе значения текущего старого узла
		// и ссылки на предыдущий новый узел.
		newCurNode = &ListNode{
			Val:  oldCurNode.Val,
			Next: newCurNode,
		}
		// Переходим к следующему старому узлу.
		oldCurNode = oldCurNode.Next

		// Уперлись в конец списка.
		if oldCurNode == nil {
			break
		}
	}

	return newCurNode
}
