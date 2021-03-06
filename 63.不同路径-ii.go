/*
 * @lc app=leetcode.cn id=63 lang=golang
 *
 * [63] 不同路径 II
 *
 * https://leetcode-cn.com/problems/unique-paths-ii/description/
 *
 * algorithms
 * Medium (36.31%)
 * Likes:    385
 * Dislikes: 0
 * Total Accepted:    89.8K
 * Total Submissions: 247.2K
 * Testcase Example:  '[[0,0,0],[0,1,0],[0,0,0]]'
 *
 * 一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
 *
 * 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
 *
 * 现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？
 *
 *
 *
 * 网格中的障碍物和空位置分别用 1 和 0 来表示。
 *
 * 说明：m 和 n 的值均不超过 100。
 *
 * 示例 1:
 *
 * 输入:
 * [
 * [0,0,0],
 * [0,1,0],
 * [0,0,0]
 * ]
 * 输出: 2
 * 解释:
 * 3x3 网格的正中间有一个障碍物。
 * 从左上角到右下角一共有 2 条不同的路径：
 * 1. 向右 -> 向右 -> 向下 -> 向下
 * 2. 向下 -> 向下 -> 向右 -> 向右
 *
 *
 */

// @lc code=start
package solution

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	result := make([][]int, len(obstacleGrid))

	for index := range result {
		result[index] = make([]int, len(obstacleGrid[index]))
	}
	if len(obstacleGrid) == 0 {
		return 0
	}
	if len(obstacleGrid[0]) == 0 {
		return 0
	}
	for i := 0; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 0 {
			result[i][0] = 1
		} else {
			break
		}
	}
	for i := 0; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 0 {
			result[0][i] = 1
		} else {
			break
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
				continue
			}
			var left int
			var high int
			if obstacleGrid[i-1][j] == 0 {
				high = result[i-1][j]
			}
			if obstacleGrid[i][j-1] == 0 {
				left = result[i][j-1]
			}

			result[i][j] = left + high
		}
	}

	return result[len(obstacleGrid)-1][len(obstacleGrid[0])-1]

}

// @lc code=end
