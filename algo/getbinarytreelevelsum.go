package main

// TODO: добавить тест

// Обход в ширину итеративно по уровням.
func getBinaryTreeLevelSum(root *TreeNode, level int) int {
	// TODO: добавить проверку на пустое дерево и невалидный уровень.
	// Текущий уровень при обходе.
	curLevel := 0
	// Максимальная сумма узлов на уровне.
	resSum := 0
	// Очередь для обработки узлов.
	queue := []*TreeNode{root}
	// Цикл для обработки по уровням.
	for len(queue) > 0 {
		// Увеличиваем текущий уровень.
		curLevel++
		// Сумма узлов на текущем уровне.
		levelSum := 0
		// Количество узлов на уровне для обработки.
		levelCount := len(queue)
		// Обрабатываем узлы текущего уровня.
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			// Увеличиваем сумму на уровне.
			levelSum += node.Val
			// Добавляем дочерние узлы в слайс следующего уровня.
			// Здесь мы можем регулировать порядок обхода - слева направо или справа налево.
			// Добавляем сначала левый узел - обходим слева направо.
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelCount:]
		// Если это искомый уровень.
		if curLevel > level {
			// Определяем результирующую сумму и выходим.
			resSum = levelSum
			break
		}
	}

	return resSum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
