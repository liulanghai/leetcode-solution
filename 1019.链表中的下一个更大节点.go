/*
 * @lc app=leetcode.cn id=1019 lang=golang
 *
 * [1019] 链表中的下一个更大节点
 *
 * https://leetcode-cn.com/problems/next-greater-node-in-linked-list/description/
 *
 * algorithms
 * Medium (55.97%)
 * Likes:    96
 * Dislikes: 0
 * Total Accepted:    10.5K
 * Total Submissions: 18.8K
 * Testcase Example:  '[2,1,5]'
 *
 * 给出一个以头节点 head 作为第一个节点的链表。链表中的节点分别编号为：node_1, node_2, node_3, ... 。
 *
 * 每个节点都可能有下一个更大值（next larger value）：对于 node_i，如果其 next_larger(node_i) 是
 * node_j.val，那么就有 j > i 且  node_j.val > node_i.val，而 j 是可能的选项中最小的那个。如果不存在这样的
 * j，那么下一个更大值为 0 。
 *
 * 返回整数答案数组 answer，其中 answer[i] = next_larger(node_{i+1}) 。
 *
 * 注意：在下面的示例中，诸如 [2,1,5] 这样的输入（不是输出）是链表的序列化表示，其头节点的值为 2，第二个节点值为 1，第三个节点值为 5
 * 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：[2,1,5]
 * 输出：[5,5,0]
 *
 *
 * 示例 2：
 *
 * 输入：[2,7,4,3,5]
 * 输出：[7,0,5,5,0]
 *
 *
 * 示例 3：
 *
 * 输入：[1,7,5,1,9,2,5,1]
 * 输出：[7,9,9,9,0,5,0,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 对于链表中的每个节点，1 <= node.val <= 10^9
 * 给定列表的长度在 [0, 10000] 范围内
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package solution

import "container/list"

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value int) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() int {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value.(int)
	}
	return 0
}

func (stack *Stack) Top() int {
	e := stack.list.Back()
	if e != nil {
		return e.Value.(int)
	}

	return 0
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func value(in []int) []int {
	//1, 7, 5, 1, 9, 2, 5, 1
	s := NewStack()
	size := len(in)
	result := make([]int, size, size)
	for i := 0; i < size; i++ {
		for !s.Empty() && in[i] > in[s.Top()] {
			result[s.Top()] = in[i]
			s.Pop()
		}
		s.Push(i)
	}

	return result

}

func nextLargerNodes(head *ListNode) []int {
	var in []int
	for head != nil {
		in = append(in, head.Val)
		head = head.Next
	}
	return value(in)
}

// @lc code=end
