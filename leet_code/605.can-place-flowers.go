package leet_code

// Перебор со сравнением

// Перебор со сравнением
func canPlaceFlowers(flowerbed []int, n int) bool {
	// Количество (счетчик) доступных мест для посадки.
	filledCount := 0
	for i := 0; i < len(flowerbed); i++ {
		// Если место занято, нечего делать.
		if flowerbed[i] == 1 {
			continue
		}

		// Если место не занято.

		// Высчитываем занятость левого места
		isLeftEmpty := false
		// Если слева пусто либо сейчас первое место (нулевой элемент), то слева пусто
		if i == 0 || flowerbed[i-1] == 0 {
			isLeftEmpty = true
		}
		isRightEmpty := false
		// Если справа пусто либо сейчас последнее место, то справа пусто
		if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
			isRightEmpty = true
		}

		// Если соседние места пустые, то увеличиваем счетчик и заполняем текущее место.
		if isLeftEmpty && isRightEmpty {
			filledCount++
			flowerbed[i] = 1
		}

		// Проверяем, не заполнили ли мы все пустые места.
		if filledCount == n {
			return true
		}
	}

	// Учитываем кейс: n = 0
	return filledCount >= n
}
