/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */
package solution

//搜索旋转数组
func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	result := -1
	for start <= end {
		if nums[start] == target {
			return start
		}
		if nums[end] == target {
			return end
		}

		if nums[start] > nums[end] {
			start = start + 1
			end = end - 1

		} else {
			mid := (start + end) / 2
			if target == nums[mid] {
				return mid
			}
			if target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}

		}
	}
	return result
}
