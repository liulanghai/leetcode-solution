/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */
package cn

import "fmt"

//从nums中获取n个数，
func getNumber(nums []int, tmp []int, start, n int, result *[][]int) {
	if len(tmp) == n {
		res := make([]int, n)
		copy(res, tmp)
		*result = append(*result, res)
		return
	}
	for i := start; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		getNumber(nums, tmp, i+1, n, result)
		tmp = tmp[0 : len(tmp)-1]
	}
}
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	var tmp []int
	for i := 0; i <= len(nums); i++ {
		getNumber(nums, tmp, 0, i, &result)
	}
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
	return result
}
