/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-postorder-traversal/description/
 *
 * algorithms
 * Hard (71.98%)
 * Likes:    367
 * Dislikes: 0
 * Total Accepted:    106.4K
 * Total Submissions: 147.4K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 后序 遍历。
 *
 * 示例:
 *
 * 输入: [1,null,2,3]
 * ⁠  1
 * ⁠   \
 * ⁠    2
 * ⁠   /
 * ⁠  3
 *
 * 输出: [3,2,1]
 *
 * 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
 *
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stack := []*TreeNode{root}
	for len(stack) != 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root == nil {
			continue
		}
		result = append(result, root.Val)
		stack = append(stack, root.Left)
		stack = append(stack, root.Right)
	}
	reverse(result)
	return result
}
func reverse(res []int) {
	i := 0
	j := len(res) - 1

	for i < j {
		res[i], res[j] = res[j], res[i]

		i++
		j--
	}
}

// @lc code=end

