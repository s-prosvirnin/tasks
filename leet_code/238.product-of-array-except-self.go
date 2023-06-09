package leet_code

// Префиксная сумма

// Префиксная сумма, точнее произведение
func productExceptSelf(nums []int) []int {
	// Результирующий массив.
	res := make([]int, len(nums))
	// Обозначим элементы на входе:
	// [a, b, c, d]

	multiplier := 1
	// Сдвигаем произведение элементов на 1 элемент вправо.
	// Получим, что в последнем произведении не будет хватать последнего элемента,
	// в предпоследнем - последнего и предпоследнего и т.д:
	// [1, a, ab, abc]
	for i, n := range nums {
		res[i] = multiplier
		multiplier *= n
	}

	multiplier = 1
	// Сдвигаем произведение элементов на 1 элемент влево и начинаем с конца.
	// Получим, что в первом произведении не будет хватать первого элемента,
	// во втором - первого и второго и т.д:
	// [bcd, cd, d, 1]
	for i := len(res) - 1; i >= 0; i-- {
		res[i] *= multiplier
		multiplier *= nums[i]
	}
	// Если перемножить первое префиксное произведение со вторым, получим искомый массив:
	// [bcd, acd, abd, abc]

	return res
}
