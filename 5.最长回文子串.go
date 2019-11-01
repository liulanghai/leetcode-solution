/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
/*
 */

/*
定义 result[i][j] 如果result[i][j] == true 则说明 i-j 是回文

递推关系是： result[i][j] = true
为长度为N的回文子串，则如果 s[i-1] = s[j+1] 则result[i-1][j+1] 为N+1的回文串。

1. 长度为1 的回文串单个字符本身即 ：result[i][i] = true
2. 长度为2 的回文串为 自己与自己相等即
   if s[i] == s[i+1] 则result[i][i+1] = true
3. 递推得到长度为3的回文串（如果存在）
。。。。。。
4. 递推得到长度为N的回文串


*/
package cn

func longestPalindrome(s string) string {
	size := len(s)
	result := make([][]bool, size, size)
	for i := 0; i < size; i++ {
		result[i] = make([]bool, size, size)
	}
	var index int
	var len int

	for i := 0; i < size; i++ {
		result[i][i] = true
		if result[i][i] && len < 1 { //长度为1的回文串
			len = 1
			index = i
		}
		if i+1 < size {
			result[i][i+1] = (s[i] == s[i+1]) //长度为2的回文串
			if result[i][i+1] {
				len = 2
				index = i
			}
		}
	}

	for l := 3; l <= size; l++ { //l为回文串长度
		for start := 0; start <= size-l; start++ { //递推判断回文串
			if s[start] == s[start+l-1] && result[start+1][start+l-2] {
				result[start][start+l-1] = true
				if l > len {
					len = l
					index = start
				}
			}
		}

	}
	return s[index : index+len]
}
