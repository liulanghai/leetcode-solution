/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package solution

import "math"

//
func searchArray(root *TreeNode, pre *int64) bool {
	if root == nil {
		return true
	}
	var res1 = true
	var res2 = true
	if root.Left != nil {
		res1 = searchArray(root.Left, pre)
	}

	if *pre >= int64(root.Val) {
		return false
	}
	*pre = int64(root.Val)

	res2 = searchArray(root.Right, pre)
	return res1 && res2
}
func isValidBST(root *TreeNode) bool {
	var a int64
	a = math.MinInt64
	return searchArray(root, &a)
}
