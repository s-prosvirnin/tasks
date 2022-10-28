package leet_code

import "math"

/**
https://leetcode.com/problems/min-cost-climbing-stairs/
746. Min Cost Climbing Stairs

Условие:
Интерпритация1
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.
Интерпритация2
Мальчик подошел к платной лестнице. Чтобы наступить на любую ступеньку, нужно заплатить указанную на ней сумму.
Мальчик умеет перешагивать на следующую ступеньку, либо перепрыгивать через ступеньку.
Требуется узнать, какая наименьшая сумма понадобится мальчику, чтобы добраться до верхней ступеньки.
Решение:
Очевидно, что сумма, которую мальчик отдаст на N-ой ступеньке, есть сумма, которую он отдал до этого плюс стоимость самой ступеньки.
«Сумма, которую он отдал до этого» зависит от того, с какой ступеньки мальчик шагает на N-ую — с (N-1)-й или с (N-2)-й. Выбирать нужно наименьшую.
*/
func minCostClimbingStairs(cost []int) int {
	count := len(cost) - 1
	for i := 2; i <= count; i++ {
		cost[i] += int(math.Min(float64(cost[i-1]), float64(cost[i-2])))
	}
	return int(math.Min(float64(cost[count]), float64(cost[count-1])))
}
