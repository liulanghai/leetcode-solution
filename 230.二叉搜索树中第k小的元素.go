/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 */

// Definition for a binary tree node.

package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findK(root *TreeNode, count *int, k int, val *int) {
	if k <= *count {
		return
	}
	if root.Left != nil {
		findK(root.Left, count, k, val)
	}
	*count = *count + 1
	if *count == k {
		*val = root.Val
		return
	}
	if root.Right != nil {
		findK(root.Right, count, k, val)
	}
}
func kthSmallest(root *TreeNode, k int) int {
	var val int
	var count int
	findK(root, &count, k, &val)
	return val
}
