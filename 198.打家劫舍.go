/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */
package solution

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func rob(nums []int) int {
	var result1, result2 int
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	result1 = nums[0]
	result2 = max(nums[0], nums[1])
	var tmp int
	for i := 2; i < len(nums); i++ {
		tmp = result2
		if nums[i]+result1 > result2 {
			result2 = nums[i] + result1

		} else {
			result2 = result2
		}
		result1 = tmp
	}
	return result2
}
