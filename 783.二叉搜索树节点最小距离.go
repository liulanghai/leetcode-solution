/*
 * @lc app=leetcode.cn id=783 lang=golang
 *
 * [783] 二叉搜索树节点最小距离
 *
 * https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/description/
 *
 * algorithms
 * Easy (53.30%)
 * Likes:    67
 * Dislikes: 0
 * Total Accepted:    15.5K
 * Total Submissions: 29.1K
 * Testcase Example:  '[4,2,6,1,3,null,null]'
 *
 * 给定一个二叉搜索树的根节点 root，返回树中任意两节点的差的最小值。
 *
 *
 *
 * 示例：
 *
 * 输入: root = [4,2,6,1,3,null,null]
 * 输出: 1
 * 解释:
 * 注意，root是树节点对象(TreeNode object)，而不是数组。
 *
 * 给定的树 [4,2,6,1,3,null,null] 可表示为下图:
 *
 * ⁠         4
 * ⁠       /   \
 * ⁠     2      6
 * ⁠    / \
 * ⁠   1   3
 *
 * 最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
 *
 *
 *
 * 注意：
 *
 *
 * 二叉树的大小范围在 2 到 100。
 * 二叉树总是有效的，每个节点的值都是整数，且不重复。
 * 本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
 * 相同
 *
 *
 */

// @lc code=start
/*
二叉搜索树中序遍历为有序的数组
*/

package solution

//TODO 可以优化内存使用，不用数组保存
func scan(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		scan(root.Left, result)
	}
	*result = append(*result, root.Val)
	if root.Right != nil {
		scan(root.Right, result)
	}

}

func minDiffInBST(root *TreeNode) int {
	var result []int
	scan(root, &result)
	minVal := result[1] - result[0]
	for i := 1; i < len(result); i++ {
		if (result[i] - result[i-1]) < minVal {
			minVal = result[i] - result[i-1]
		}
	}
	return minVal
}

// @lc code=end
