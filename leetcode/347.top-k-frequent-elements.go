package leetcode

// https://leetcode.com/problems/top-k-frequent-elements

// кажется это решение O(n). Также можно решить с помощью хэш-таблицы и кучи или хэш-таблицы и сортировки
func topKFrequent(nums []int, k int) []int {
	numsFreq := make(map[int]int, len(nums))
	for _, num := range nums {
		numsFreq[num]++
	}

	freqNums := make(map[int][]int, k)

	maxFreq := 0
	for num, freq := range numsFreq {
		_, ok := freqNums[freq]
		if !ok {
			freqNums[freq] = make([]int, 0)
		}
		freqNums[freq] = append(freqNums[freq], num)

		if freq > maxFreq {
			maxFreq = freq
		}
	}

	numsRes := make([]int, 0, k)
	for i := maxFreq; i >= 0; i-- {
		numsByFreq, ok := freqNums[i]
		if !ok {
			continue
		}
		for j := 0; j < len(numsByFreq); j++ {
			if len(numsRes) == k {
				break
			}
			numsRes = append(numsRes, numsByFreq[j])
		}
		if len(numsRes) == k {
			break
		}
	}

	return numsRes
}
