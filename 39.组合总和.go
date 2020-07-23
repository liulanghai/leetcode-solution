/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */
package solution

import "fmt"

func sumSlice(s []int) int {
	sum := 0
	for _, val := range s {
		sum = sum + val
	}
	return sum
}
func combinationCal(pos int, candidates []int, sel []int, target int, result *[][]int) int {
	var sum = sumSlice(sel)
	if sum == target {
		temp := make([]int, len(sel))
		copy(temp, sel)
		*result = append(*result, temp)
		return 0
	}
	//需要回溯回去重新选择
	if sum > target {
		return -1
	}
	for i := pos; i < len(candidates); i++ {

		var tmp []int
		for n := 0; n < len(candidates); n++ {
			if n == i {
				continue
			}
			tmp = append(tmp, candidates[n])
		}
		combinationCal(i, tmp, sel, target, result)

		for j := 1; j <= target/candidates[i]; j++ { //选了i， 选多个

			for m := 0; m < j; m++ {
				sel = append(sel, candidates[i])
			}
			//如果这条路已经选到了，没必要再回溯了dl
			if combinationCal(i, candidates, sel, target, result) == -1 {
				return 0
			}
			sel = sel[:len(sel)-j]

		}
	}
	//路走完了，不需要回溯了
	return 0

}
func combinationSum(candidates []int, target int) [][]int {
	var sel []int
	var result [][]int
	combinationCal(0, candidates, sel, target, &result)
	fmt.Println(result)
	return result
}
