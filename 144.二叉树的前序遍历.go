/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-preorder-traversal/description/
 *
 * algorithms
 * Medium (66.20%)
 * Likes:    336
 * Dislikes: 0
 * Total Accepted:    148.3K
 * Total Submissions: 223.1K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的 前序 遍历。
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
 * 输出: [1,2,3]
 *
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

package solution

func preorderTraversal(root *TreeNode) []int {
	var result []int
	stack := NewListStack()
	tmp := root
	for tmp != nil || stack.Empty() == false {
		for tmp != nil {
			stack.Push(tmp)
			result = append(result, tmp.Val)
			tmp = tmp.Left
		}

		tmp = stack.Pop()
		tmp = tmp.Right
	}
	return result
}

// @lc code=end
