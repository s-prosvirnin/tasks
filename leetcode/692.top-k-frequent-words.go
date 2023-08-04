package leetcode

import (
	"container/heap"
	"sort"
)

// https://leetcode.com/problems/top-k-frequent-words

func topKFrequent(words []string, k int) []string {
	return hashMapAndHeap(words, k)
}

type PqItem struct {
	Word string
	Freq int
}

type PQ []PqItem

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	if pq[i].Freq > pq[j].Freq {
		return true
	}
	if pq[i].Freq == pq[j].Freq {
		return pq[i].Word < pq[j].Word
	}

	return false
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(PqItem))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// хэш-таблица и куча (приоритетная очередь). O(n log(k)) time, O(n) memory
func hashMapAndHeap(words []string, k int) []string {
	wordsFreq := make(map[string]int, len(words))
	for _, word := range words {
		wordsFreq[word]++
	}

	wordsCount := len(wordsFreq)
	pqv := make(PQ, 0, wordsCount)
	pq := &pqv
	for word, freq := range wordsFreq {
		heap.Push(pq, PqItem{word, freq})
	}
	wordsSorted := make([]string, 0, wordsCount)
	for i := 0; i < wordsCount; i++ {
		wordsSorted = append(wordsSorted, heap.Pop(pq).(PqItem).Word)
	}

	return wordsSorted[0:k]
}

// хэш-таблица и сортировка по частотности. O(n log(n)) time, O(n) memory
func hashMapAndSort(words []string, k int) []string {
	wordsFreq := make(map[string]int, len(words))
	for _, word := range words {
		wordsFreq[word]++
	}

	wordsRes := make([]string, 0, len(words))
	for word := range wordsFreq {
		wordsRes = append(wordsRes, word)
	}

	sort.Slice(
		wordsRes, func(i, j int) bool {
			if wordsFreq[wordsRes[i]] > wordsFreq[wordsRes[j]] {
				return true
			}
			if wordsFreq[wordsRes[i]] == wordsFreq[wordsRes[j]] {
				return wordsRes[i] < wordsRes[j]
			}

			return false
		},
	)

	return wordsRes[:k]
}
