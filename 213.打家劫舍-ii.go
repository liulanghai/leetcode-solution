/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */
/*
solution:
1.如果0确定不偷的话，那么问题转化为1-len(nums)。最后一家随意偷与不偷，转化为
198问题
2.如果最后一个确定不偷的话，问题转化为 0-len(nums)-1

*/
package solution

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a := rob(nums[1:])
	b := rob(nums[0 : len(nums)-1])
	return max(a, b)
}
