/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 /*

result[i] = k 代表从i开始最长的有效括号。


*/
package cn

type loc struct {
	start int
	len   int
}

// @lc code=start
func longestValidParentheses(s string) int {
	size := len(s)
	result := make([][]loc, len(s), len(s))
	for i := 0; i < size; i++ {
		result[i] = make([]loc, size, size)
	}
	for i := 0; i < size-1; i++ {
		if s[i] == '(' && s[i+1] == ')' { //（）这种情况
			result[i][i+1].start = i
			result[i][i+1].len = 2
		}
	}
	for i := 2; i < size; i=i+2 {
		for j := 0; j < size-i; j++ {
			if j-1 >= 0 {
				if j+1 < size { //只求前后的最长的，其他的可以来拼接。
					if s[j-1] == '(' && s[j+i+1] == ')' && result[j][j+i-1].len ==i {
						result[j-1][j+1].start = 
					}
				}

			}

		}
	}
}

// @lc code=end
