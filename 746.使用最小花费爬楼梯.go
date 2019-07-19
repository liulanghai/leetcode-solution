/*
 * @lc app=leetcode.cn id=746 lang=golang
 *
 * [746] 使用最小花费爬楼梯
 */
package solution

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func minCostClimbingStairs(cost []int) int {
	if len(cost) < 2 {
		return 0
	}
	//踏上cost[i]的最小值
	cost1 := cost[0]
	cost2 := cost[1]
	for i := 2; i < len(cost); i++ {
		tmp := cost2
		cost2 = min(cost1, cost2) + cost[i]
		cost1 = tmp
	}
	return min(cost2, cost1)
}
