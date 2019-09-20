/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */
package cn

//采用二分查找，注意边界条件
func findMin(nums []int) int {
	l := len(nums)
	start := 0
	end := l - 1
	for (start + 1) < end {
		mid := (start + end) / 2
		if nums[mid] > nums[end] { //在 mid - end 之间
			start = mid
		} else {
			end = mid
		}
	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]

}
