/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package cn

//按层次遍历
func largestValues(root *TreeNode) []int {
	var res []int
	var tmp []*TreeNode
	if root == nil {
		return res
	}
	tmp = append(tmp, root)
	for len(tmp) > 0 {
		var next []*TreeNode
		max := tmp[0].Val
		for _, val := range tmp {
			if val.Val > max {
				max = val.Val
			}
			if val.Left != nil {
				next = append(next, val.Left)
			}
			if val.Right != nil {
				next = append(next, val.Right)
			}
		}
		res = append(res, max)
		tmp = next
	}
	return res
}
