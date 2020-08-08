/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 *
 * https://leetcode-cn.com/problems/binary-tree-inorder-traversal/description/
 *
 * algorithms
 * Medium (72.19%)
 * Likes:    613
 * Dislikes: 0
 * Total Accepted:    209.8K
 * Total Submissions: 289.7K
 * Testcase Example:  '[1,null,2,3]'
 *
 * 给定一个二叉树，返回它的中序 遍历。
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
 * 输出: [1,3,2]
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

import (
	"container/list"
)

type ListStack struct {
	list *list.List
}

func NewListStack() *ListStack {
	list := list.New()
	return &ListStack{list}
}

func (ListStack *ListStack) Push(value *TreeNode) {
	ListStack.list.PushBack(value)
}

func (ListStack *ListStack) Pop() *TreeNode {
	e := ListStack.list.Back()
	if e != nil {
		ListStack.list.Remove(e)
		return e.Value.(*TreeNode)
	}
	return nil
}

func (ListStack *ListStack) Top() *TreeNode {
	e := ListStack.list.Back()
	if e != nil {
		return e.Value.(*TreeNode)
	}

	return nil
}

func (ListStack *ListStack) Len() int {
	return ListStack.list.Len()
}

func (ListStack *ListStack) Empty() bool {
	return ListStack.list.Len() == 0
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	stack := NewListStack()
	tmp := root
	for stack.Empty() == false || tmp != nil {
		for tmp != nil {
			stack.Push(tmp)
			tmp = tmp.Left
		}
		tmp = stack.Pop()
		result = append(result, tmp.Val)
		tmp = tmp.Right
	}
	return result
}

// @lc code=end
