package leet_code

/**
https://leetcode.com/problems/two-sum/
1. Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
*/
func twoSum(nums []int, target int) []int {
	numbersPositionsMap := map[int][]int{}
	for pos, number := range nums {
		if len(numbersPositionsMap[number]) == 0 {
			numbersPositionsMap[number] = make([]int, 0, 5)
		}
		numbersPositionsMap[number] = append(numbersPositionsMap[number], pos)
	}

	sumNumbersPos := make([]int, 2)
	for number, numberPositions := range numbersPositionsMap {
		secNum := target - number
		secNumPositions, ok := numbersPositionsMap[secNum]
		if ok {
			if number == secNum && len(secNumPositions) <= 1 {
				continue
			}
			sumNumbersPos[0] = numberPositions[0]
			if number != secNum {
				sumNumbersPos[1] = secNumPositions[0]
			} else {
				sumNumbersPos[1] = secNumPositions[1]
			}
			break
		}
	}

	return sumNumbersPos
}
