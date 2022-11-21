package leet_code

// https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int) {
	pos1 := m - 1
	pos2 := n - 1
	for pos := len(nums1) - 1; pos >= 0; pos-- {
		if pos2 < 0 && pos1 < 0 {
			break
		}
		if pos2 < 0 {
			nums1[pos] = nums1[pos1]
			pos1--
			continue
		}
		if pos1 < 0 {
			nums1[pos] = nums2[pos2]
			pos2--
			continue
		}
		if nums1[pos1] >= nums2[pos2] {
			nums1[pos] = nums1[pos1]
			pos1--
			continue
		}
		if nums1[pos1] < nums2[pos2] {
			nums1[pos] = nums2[pos2]
			pos2--
			continue
		}
	}
}
