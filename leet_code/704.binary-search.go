package leet_code

// https://leetcode.com/problems/binary-search/

func search(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}
	if len(nums) == 1 {
		return -1
	}
	middlePos := int(len(nums) / 2)
	var searchIndex int
	if target >= nums[middlePos] {
		searchIndex = search(nums[middlePos:], target)
		if searchIndex == -1 {
			return -1
		}
		return middlePos + searchIndex
	} else {
		searchIndex = search(nums[0:middlePos], target)
		if searchIndex == -1 {
			return -1
		}
		return searchIndex
	}
}
