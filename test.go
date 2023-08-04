package main

import (
	"sort"
	"time"
)

func main() {

}

/*
Дан список целых чисел, повторяющихся элементов в списке нет.
Нужно преобразовать это множество в строку,
сворачивая соседние по числовому ряду числа в диапазоны.

Примеры:
- [1, 4, 5, 2, 3, 9, 8, 11, 0] => "0-5,8-9,11"
- [1, 4, 3, 2] => "1-4"
- [1, 4] => "1,4"
*/

func compress(l []int) string {
	// Результирующий массив.
	var res []byte
	// Обрабатываем кейс тут, чтобы сразу назначить startNum.
	if len(l) == 0 {
		return string(res)
	}
	// Сортируем.
	sort.Ints(l)
	// Начальный элемент интервала
	startNum := l[0]
	// 0-й элемент нам не интересен.
	// Так же нам нужно обработать последний элемент, для этого выходим за рамки массива.
	// Так же можно сравнивать текущий элемент с последующим, чтобы не выходить за рамки массива.
	for i := 1; i <= len(l); i++ {
		// Создаем интервал.
		// Для этого либо текущий элемент должен быть началом нового интервала.
		// Либо выходим за рамки массива и обрабатываем последний интервал.
		if i == len(l) || l[i]-l[i-1] > 1 {
			// Интервал состоит из одного числа.
			if l[i-1] == startNum {
				res = append(res, byte(startNum))
			} else { // Интервал состоит из двух чисел.
				res = append(res, byte(startNum), '-', byte(l[i-1]))
			}
			// Если мы не вышли за пределы массива, переустанавливаем начало интервала.
			if i < len(l) {
				startNum = l[i]
			} else {
				// Добавляем запятую, если это не последний интервал.
				res = append(res, ',')
			}
		}
	}

	return string(res)
}

/*Есть последовательность запросов, упорядоченная по времени.
Запросы бывают двух видов:

Пользователь user_id сгенерировал событие (нажал на красную кнопку)
Посчитать количество пользователей, которые за последние 5 минут сгенерировали >= 1000 событий (нажали на красную кнопку >= 1000 раз).

Необходимо реализовать структуру данных, умеющую эффективно обрабатывать данные запросы.

*/

// All methods are guaranteed to be called with non-decreasing (>=) time

// Событие.
type event struct {
	// Время события.
	t time.Time
	// id пользователя.
	u int
}

type UserStatistics struct {
	// Двусторонняя очередь событий.
	events []event
	// Счетчики пользователей. Ключ - id пользователя, значение - счетчик за window время.
	userCounters map[int]int
	// Общий счетчик пользователей, которые превысили количество запросов (счетчик нарушителей).
	usersLimitCounter int
	// Временное окно, в котором нужно считать счетчики.
	window time.Duration
	// Ограничение по количеству событий для каждого пользователя за время window.
	eventsLimit int
}

func NewUserStatistics(window time.Duration, limit int) *UserStatistics {
	return &UserStatistics{
		userCounters: make(map[int]int),
		window:       window,
		eventsLimit:  limit,
	}
}

func (u *UserStatistics) AddEvent(now time.Time, userID int) {
	// Добавляем новое событие в конец очереди.
	u.events = append(u.events, event{t: now, u: userID})
	// Если счетчик пользователя на грани ограничения, то увеличиваем счетчик нарушителей.
	// Т.о. не нужно использовать лишние переменные, чтобы считать до и после.
	if u.userCounters[userID] == u.eventsLimit-1 {
		u.usersLimitCounter++
	}
	// Увеличиваем счетчик пользователя.
	u.userCounters[userID]++
	// Удаляем все события с начала очереди, которые вышли за пределы окна window.
	u.remove(now)
}

// GetRobotCount returns number of robots.
// User is robot at given time now <=> they have >= eventsLimit clicks since (>=) now - window.
func (u *UserStatistics) GetRobotCount(now time.Time) int {
	// Удаляем события, не входящие в окно.
	u.remove(now)
	// Теперь имеем только события входящие в окно и счетчик актуализирован.
	return u.usersLimitCounter
}

func (u *UserStatistics) remove(now time.Time) {
	for _, event := range u.events {
		if now.Sub(event.t) <= u.window {
			break
		}
		u.events = u.events[1:]
		// Если пользователь был нарушителем, но после удаления события перестанет быть нарушителем - декрементируем счетчик нарушителей.
		// Т.о. не нужно использовать лишние переменные, чтобы считать до и после.
		if u.userCounters[event.u] == u.eventsLimit {
			u.usersLimitCounter--
		}
		// Декрементируем счетчик пользователя.
		u.userCounters[event.u]--
		// Удаляем счетчик из мапы, если необходимо.
		if u.userCounters[event.u] == 0 {
			delete(u.userCounters, event.u)
		}
	}
}
