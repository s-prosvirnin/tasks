package leetcode

import "math"

// Простое сравнение (эталон решения)

func increasingTriplet(nums []int) bool {
	// Первое число (самое минимальное из 3-х), устанавливаем в максимальное для int.
	first := math.MaxInt
	// Второе число (больше 1-го и меньше 3-го), устанавливаем в максимальное для int.
	second := math.MaxInt
	// Третье число не нужно запоминать, при его нахождении - сразу выходим.

	for _, n := range nums {
		// Если нашли число меньшее, чем первое - подменяем первое.
		// Т.к. мы его установили в максимальное при инициализации,
		// то оно сразу переопределяется.
		if n < first {
			first = n
			continue
		}
		// Если нашли число меньшее, чем второе - подменяем второе.
		// При этом это число точно больше чем первое, т.к. иначе сработала бы проверка выше.
		if n < second && n > first {
			second = n
			continue
		}
		// Если число больше чем второе - выходим.
		if n > second {
			return true
		}
	}

	// Суть решения в том, что если мы находим число меньше, то подменяем числа по порядку.
	// Т.о. мы можем как бы начать отсчет сначала, когда это необходимо.
	// Т.е. если удасться подменить два числа, то отсчет начнется сначала,
	// если только одно, то отсчет продолжиться и ничего не сломается.

	// Кейс: [3,7,4,1,5]. Искомые числа - [3,4,5].
	// Можно первое число подменить число 1 и начать отсчет сначала,
	// но тогда мы вернем false, т.к. после числа 1 всего одно число.
	// Поэтому в решении мы не начинаем сначала, а просто подменяем первое число.

	// Кейс: [3,7,4,1,2,3]. Искомые числа - [1,2,3].
	// При обработке числа 1 нужно начать отсчет сначала, иначе, оставив 3
	// в качестве первого числа, мы получим только два числа - [3,4].
	// Поэтому в решении мы подменяем первое число и второе число,
	// как бы начиная отсчет сначала.

	// Кейс: [1,1,-2,6]. Искомые числа - [-2,6]. Возвращаем false.

	return false
}