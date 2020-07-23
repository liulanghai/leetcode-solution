/*
 * @lc app=leetcode.cn id=154 lang=golang
 *
 * [154] 寻找旋转排序数组中的最小值 II
 */
package solution

func findMin2(nums []int) int {
	l := len(nums)
	start := 0
	end := l - 1
	for (start + 1) < end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] { //在 mid - end 之间
			start = mid
		} else if nums[mid] == nums[end] {
			end = end - 1
		} else {
			end = mid
		}

	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}
