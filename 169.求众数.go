/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 *
 */
package cn

func majorityElement(nums []int) int {
	var count int
	var result int
	result = nums[0]
	for i := 0; i < len(nums); i++ {
		if result == nums[i] {
			count++
		} else {
			count--
			if count == -1 {
				result = nums[i]
				count = 1
			}
		}
	}
	return result
}
