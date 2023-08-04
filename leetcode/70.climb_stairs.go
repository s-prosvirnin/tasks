package leetcode

/*
*
https://leetcode.com/problems/climbing-stairs/submissions/
70. Climbing Stairs

You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*/
func climbStairs(n int) int {
	return climbStairsByArray(n)
}

func climbStairsByArray(n int) int {
	stairs := make([]int, 2)
	stairs[0] = 1
	stairs[1] = 2
	for i := 2; i < n; i++ {
		stairs[i%2] = stairs[0] + stairs[1]
	}

	return stairs[(n-1)%2]
}

// very slowly
func climbStairsByRecursion(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 2
	}

	return climbStairsByRecursion(n-1) + climbStairsByRecursion(n-2)
}
