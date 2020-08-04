/*
 * @lc app=leetcode.cn id=300 lang=golang
 *
 * [300] 最长上升子序列
 *
 * https://leetcode-cn.com/problems/longest-increasing-subsequence/description/
 *
 * algorithms
 * Medium (44.88%)
 * Likes:    868
 * Dislikes: 0
 * Total Accepted:    124.2K
 * Total Submissions: 276.1K
 * Testcase Example:  '[10,9,2,5,3,7,101,18]'
 *
 * 给定一个无序的整数数组，找到其中最长上升子序列的长度。
 *
 * 示例:
 *
 * 输入: [10,9,2,5,3,7,101,18]
 * 输出: 4
 * 解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
 *
 * 说明:
 *
 *
 * 可能会有多种最长上升子序列的组合，你只需要输出对应的长度即可。
 * 你算法的时间复杂度应该为 O(n^2) 。
 *
 *
 * 进阶: 你能将算法的时间复杂度降低到 O(n log n) 吗?
 *
 */

// @lc code=start
package solution

func lengthOfLIS(nums []int) int {

	state := make([]int, len(nums), len(nums))
	//state[i] 代表到i时最长，并且选取了的。
	if len(nums) == 0 {
		return 0
	}

	var result = 1
	state[0] = 1
	for i := 1; i < len(nums); i++ {
		var maxValue int
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxValue = max(maxValue, state[j])
			}
		}
		state[i] = maxValue + 1
		result = max(result, state[i])
	}
	return result

}

// @lc code=end
