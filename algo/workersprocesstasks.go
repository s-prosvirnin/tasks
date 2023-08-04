package main

import (
	"container/heap"
	"fmt"
)

func main() {
	testWorkersProcessTasks()
}

func testWorkersProcessTasks() {
	type testCase struct {
		jobs     []int
		workers  int
		expected int
	}

	testCases := []testCase{
		{jobs: []int{1, 1, 1, 1, 1}, workers: 1, expected: 5},
		{jobs: []int{3, 2, 3}, workers: 3, expected: 3},
		{jobs: []int{2, 2, 1, 2, 1}, workers: 2, expected: 4},
		{jobs: []int{7, 6, 5, 4, 3, 2, 1, 1, 2, 3, 4, 5, 6, 7}, workers: 5, expected: 14},
		{jobs: []int{1, 2, 4, 7, 8}, workers: 2, expected: 13},
		{jobs: []int{5, 4, 2, 7, 8, 2, 6}, workers: 3, expected: 13},
		{jobs: []int{5}, workers: 2, expected: 5},
		{jobs: []int{}, workers: 2, expected: 0},
	}

	passedCount := 0
	overallCount := len(testCases)
	for _, testCase := range testCases {
		if res := workersProcessTasks(testCase.jobs, testCase.workers); res != testCase.expected {
			fmt.Printf("Expected: %d; result: %d; \n", testCase.expected, res)
		} else {
			passedCount++
		}
	}

	fmt.Printf("Tests done - passed: %d; overall: %d; \n", passedCount, overallCount)
}

func workersProcessTasks(jobs []int, workers int) int {
	// Минимальная куча для хранения текущего затраченного времени по воркерам.
	minHeap := make(minHeap, 0, len(jobs))
	// Минимальное текущее время работы одного из воркеров.
	workersMinTime := 0
	// Результирующее максимальное количество времени.
	res := 0
	// Бежим по джобам.
	for _, jobTime := range jobs {
		// Нам нужно хранить элементов не больше чем воркеров,
		// поэтому вынимаем минимально затраченное время одного из воркеров.
		if minHeap.Len() == workers {
			workersMinTime = heap.Pop(&minHeap).(int)
		}
		// Новое время воркера с учетом текущей джобы.
		workerTime := jobTime + workersMinTime
		heap.Push(&minHeap, workerTime)
		// Если время воркера больше максимального - перезаписываем.
		if workerTime > res {
			res = workerTime
		}
	}

	return res
}

// Реализация минимальной кучи для container/heap.
type minHeap []int

func (w minHeap) Len() int {
	return len(w)
}
func (w minHeap) Less(i, j int) bool {
	return w[i] < w[j]
}
func (w minHeap) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
func (w *minHeap) Pop() interface{} {
	a := *w
	n := len(a) - 1
	v := a[n]
	*w = a[:n]
	return v
}
func (w *minHeap) Push(x interface{}) {
	*w = append(*w, x.(int))
}
