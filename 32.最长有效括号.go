/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 /*

result[i] = 从i开始到结尾的最长有效括号 resulti是第一个(符号
从后往前推导：
if result[i-1] == '(' && result[i]+ index 是否为 ')'  // (())的情况
	result[i-1] = result[i] + 2
	if result[i]+ index + 1 < size     //子串相邻的情况,由于（括号将将 )匹配了，之前的还可以增加。
	如  (())()情况
	result[i-1] = result[result[i]+ index + 1]


*/
package solution

// @lc code=start
func longestValidParentheses(s string) int {
	size := len(s)
	if size < 2 {
		return 0
	}
	result := make([]int, len(s), len(s))
	max := 0
	result[size-1] = 0

	for i := size - 2; i > -1; i-- {
		if s[i] == '(' {
			sym := result[i+1] + i + 1 //可能匹配的 )的位置
			if sym < size && s[sym] == ')' {
				result[i] = 2 + result[i+1] //匹配了

				if sym+1 < size { //是否有连续的
					result[i] = result[i] + result[sym+1]
				}
				if result[i] > max {
					max = result[i]
				}
			}
		}
	}

	return max
}

// @lc code=end
