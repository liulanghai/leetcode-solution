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
	result := make([]int, len(s), len(s))
	max := 0
	for i := 0; i < size-1; i++ {
		if s[i] == '(' && s[i+1] == ')' { //（）这种情况
			result[i] = 2
			if result[i] > max {
				max = result[i]
			}
		}
	}
	flag := true
	for flag {
		flag = false
		for j := 0; j < size; j++ {
			if result[j] > 0 {
				l := result[j]
				if j-1 >= 0 && l+j < size {
					if s[j-1] == '(' && s[l+j] == ')' && result[j-1] == 0 {
						result[j-1] = 2 + result[j]
						flag = true
						if result[j-1] > max {
							max = result[j-1]
						}
					}
				}
				if j+l < size && result[j+l] > 0 {
					result[j] = result[j] + result[j+l]
					flag = true
					if result[j] > max {
						max = result[j]
					}
				}
			}
		}
	}

	return max
}

// @lc code=end
