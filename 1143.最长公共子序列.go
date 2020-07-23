/*
 * @lc app=leetcode.cn id=1143 lang=golang
 *
 * [1143] 最长公共子序列
 */
package solution

// @lc code=start
//longestCommonSubsequence
//solution 1 分治：
// 序列 Text1[0~N] Text2[0~M]
// 如果 Text1[N] != Text2[M] 则规模减少 为求(Text1[0~N] Text2[0~M-1])与 (Text1[0~N-1] Text2[0~M])的最大值
// 如果 Text1[N] == Text2[M] 则规模减少 求 (Text1[0~N-1] Text2[0~M-1])+ 1

// 分治会到来大量的重复计算,将递归转化为递推解决

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func longestCommonSubsequence(text1 string, text2 string) int {

	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	resultArray := make([][]int, len(text1), len(text1))
	for i := 0; i < len(text1); i++ {
		resultArray[i] = make([]int, len(text2), len(text2))
	}
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				last := 0
				if i-1 >= 0 && j-1 >= 0 {
					last = resultArray[i-1][j-1]
				}
				resultArray[i][j] = last + 1
			} else {
				level1 := 0
				if i-1 >= 0 {
					level1 = resultArray[i-1][j]
				}
				level2 := 0
				if j-1 >= 0 {
					level2 = resultArray[i][j-1]
				}
				resultArray[i][j] = max(level2, level1)
			}
		}
	}
	return resultArray[len(text1)-1][len(text2)-1]

}

// @lc code=end
