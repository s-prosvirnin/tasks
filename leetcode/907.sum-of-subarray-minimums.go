package leetcode

// https://leetcode.com/problems/sum-of-subarray-minimums/

const limit = 1000000007

func sumSubarrayMins(arr []int) int {
	return sumSubarrayMins2(arr)
}

/**
* есть более быстрое решение, вроде, монотонная очередь
 */
func sumSubarrayMins1(arr []int) int {
	sum := 0
	totalMin := 0
	for pos1, num := range arr {
		// запоминаем минимум
		currentMin := num
		// бежим по оставшимся числам справа, т.к. для всех чисел слева мы уже просуммировали минимумы
		for pos2 := pos1; pos2 < len(arr); pos2++ {
			// с каждой итерацией приращиваем массив для текущего числа (num)
			// формируем минимум на основе текущего минимума и очередного числа
			// потому как левая часть массива не меняется, минимум для нее уже известен
			// таким образом мы не вычисляем минимум для всего массива каждый раз
			if arr[pos2] < currentMin {
				currentMin = arr[pos2]
			}
			// суммируем минимум
			sum += currentMin
			sum %= 1000000007
		}
		// вычисляем минимум для всех элементов массива отдельно
		if num < totalMin {
			totalMin = num
		}
	}
	sum += totalMin
	sum %= 1000000007

	return sum
}

func sumSubarrayMins2(A []int) int {
	var sol int
	mmap := make([]int, len(A))
	stack := []int{}
	for i := range A {
		var cur int
		for len(stack) != 0 && A[i] < A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			cur = (i + 1) * A[i]
		} else {
			cur = mmap[stack[len(stack)-1]] + (i-stack[len(stack)-1])*A[i]
		}
		cur %= limit
		stack = append(stack, i)
		mmap[i] = cur
		sol += cur
		sol %= limit
	}

	return sol
}
