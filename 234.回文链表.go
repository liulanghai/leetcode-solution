/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (42.67%)
 * Likes:    589
 * Dislikes: 0
 * Total Accepted:    112.5K
 * Total Submissions: 262.5K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 *
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 *
 * 输入: 1->2->2->1
 * 输出: true
 *
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */

// @lc code=start
package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reveserLind(head *ListNode) *ListNode {
	tmp := head
	if head == nil {
		return nil
	}
	next := head.Next
	tmp.Next = nil
	for next != nil {
		val := next.Next
		next.Next = tmp
		tmp = next
		next = val
	}
	head = tmp
	return head
}
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	tmp := head
	count := 0
	for tmp != nil {
		tmp = tmp.Next
		count++
	}
	pre := head
	start1 := head
	start2 := head
	for i := 0; i < count/2; i++ {
		pre = head
		head = head.Next
	}
	pre.Next = nil
	start2 = reveserLind(head)

	for start1 != nil {
		if start1.Val != start2.Val {
			return false
		}
		start1 = start1.Next
		start2 = start2.Next
	}
	return true

}

// @lc code=end
