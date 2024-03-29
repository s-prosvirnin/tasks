package leetcode

import "sort"

// Используем хэш-таблицу
func maxOperations(nums []int, k int) int {
	// Мапа: ключ - число из массива, значение - сколько раз число встретилось в массиве.
	numCounters := make(map[int]int)
	// Результат - количество исключенных пар из массива
	res := 0
	// n - первое число из пары.
	for _, n := range nums {
		// Если число больше k, то пару для него точно не найти - пропускаем.
		if n >= k {
			continue
		}

		// Необходимое второе число из пары, которое нужно найти.
		diff := k - n
		// Если второе число есть в мапе и его счетчик положительный -
		// уменьшаем счетчик (исключили второе число) и инкрементируем результирующий счетчик
		if count, ok := numCounters[diff]; ok && count > 0 {
			res++
			numCounters[diff]--
		} else { // Не нашли второе число - инкрементируем счетчик числа.
			numCounters[n]++
		}
	}

	return res
}

// Используем сортировку
func maxOperations1(nums []int, k int) int {
	// Сортируем числа.
	sort.Ints(nums)
	// Указатель на левое число (левый бегунок).
	leftP := 0
	// Указатель на правое число (правый бегунок).
	rightP := len(nums) - 1
	// Результат - количество исключенных пар.
	res := 0
	// Пока бегунки не дойдут друг до друга.
	for leftP < rightP {
		// Сумма чисел бегунков.
		sum := nums[leftP] + nums[rightP]
		// Если сумма равна k, то мы нашли искомую пару:
		// - сдвигаем оба бегунка друг к другу.
		// - увеличиваем результирующий счетчик
		if sum == k {
			leftP++
			rightP--
			res++
			continue
		}
		// Если сумма больше k, то нужно передвинуть правый бегунок.
		// Т.к. для правого бегунка нужно число меньшее, чем под левым бегунком.
		if sum > k {
			rightP--
			continue
		}

		// Если сумма меньше k, то нужно передвинуть левый бегунок.
		// Т.к. для левого бегунка нужно число большее, чем под правым бегунком.
		leftP++
	}

	return res
}
