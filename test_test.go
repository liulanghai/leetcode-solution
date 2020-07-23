package solution

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindMin(t *testing.T) {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	resA := findMin(a)
	b := []int{3, 4, 5, 1, 2}
	resB := findMin(b)
	Convey("[153]寻找旋转排序数组中的最小值", t, func() {
		So(resA, ShouldEqual, 0)
		So(resB, ShouldEqual, 1)
		So(findMin([]int{1, 2}), ShouldEqual, 1)
		So(findMin([]int{3, 1, 2}), ShouldEqual, 1)
	})
}

func TestFindMin2(t *testing.T) {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	resA := findMin2(a)
	b := []int{3, 4, 5, 1, 2}
	resB := findMin2(b)
	Convey("[154]寻找旋转排序数组中的最小2", t, func() {
		So(resA, ShouldEqual, 0)
		So(resB, ShouldEqual, 1)
		So(findMin2([]int{1, 2}), ShouldEqual, 1)
		So(findMin2([]int{3, 1, 2}), ShouldEqual, 1)
		So(findMin2([]int{2, 2, 2, 0, 1}), ShouldEqual, 0)
		So(findMin2([]int{2, 2, 2, 1, 1}), ShouldEqual, 1)
		So(findMin2([]int{2, 2, 2, 1, 2}), ShouldEqual, 1)
		So(findMin2([]int{1, 2, 2}), ShouldEqual, 1)
		So(findMin2([]int{10, 1, 10, 10, 10}), ShouldEqual, 1)
	})
}

/*
   3
  / \
 1   4
  \
   2
*/
func TestKthSmallest(t *testing.T) {
	A := TreeNode{Val: 3}
	B := TreeNode{Val: 1}
	C := TreeNode{Val: 4}
	D := TreeNode{Val: 2}
	A.Left = &B
	A.Right = &C
	B.Right = &D
	Convey("[230]二叉搜索树中第K小的元素", t, func() {

		So(kthSmallest(&A, 1), ShouldEqual, 1)

	})

}

func TestIsValidBST(t *testing.T) {
	A := TreeNode{Val: 1}
	B := TreeNode{Val: 1}
	A.Left = &B
	//*[10,5,15,null,null,6,20]*/
	Convey("[230]二叉搜索树中第K小的元素", t, func() {

		So(isValidBST(&A), ShouldEqual, false)
	})
}

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	Convey("[78]子集", t, func() {
		So(len(result), ShouldEqual, 8)
	})
}

func TestMajorityElement(t *testing.T) {
	nums := []int{1, 2, 3, 2, 2}
	result := majorityElement(nums)
	result2 := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	Convey("[169]求众数", t, func() {
		So(result, ShouldEqual, 2)
		So(result2, ShouldEqual, 2)
	})
}

func TestSearchXuanzhuan(t *testing.T) {
	Convey("[33]搜索旋转数组", t, func() {
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), ShouldEqual, 4)
		So(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), ShouldEqual, -1)
		So(search([]int{1, 3}, 3), ShouldEqual, 1)
		So(search([]int{3, 1}, 1), ShouldEqual, 1)
		So(search([]int{5, 1, 3}, 3), ShouldEqual, 2)
		//So(result2, ShouldEqual, 2)
	})
}

func TestCombinationSum(t *testing.T) {
	Convey("[39] 组合总和", t, func() {
		So(len(combinationSum([]int{2, 3, 6, 7}, 7)), ShouldEqual, 2)
		//	So(len(combinationSum([]int{2}, 4)), ShouldEqual, 2)
		So(len(combinationSum([]int{2, 3, 5}, 8)), ShouldEqual, 3)
		So(len(combinationSum([]int{8, 7, 4, 3}, 11)), ShouldEqual, 2)
	})
}

func TestReverseBits(t *testing.T) {
	Convey("[190] 颠倒二进制位\n", t, func() {
		fmt.Println(reverseBits(43261596) & 43261596)
		So(reverseBits(5), ShouldEqual, 0)
	})
}
func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300)
	val1 := cache.Get(1)
	val2 := cache.Get(2)
	val3 := cache.Get(3)

	cache.Put(4, 400)
	cache.Put(3, 310)
	Convey("[146] LRUcache\n", t, func() {
		So(val1, ShouldEqual, -1)
		So(val2, ShouldEqual, 200)
		So(val3, ShouldEqual, 300)
		So(cache.Get(2), ShouldEqual, -1)
		So(cache.Get(3), ShouldEqual, 310)
	})
}
func TestLongestPalindrome(t *testing.T) {
	s1 := longestPalindrome("babad")
	Convey("[146] LRUcache\n", t, func() {
		So(s1, ShouldEqual, "bab")

	})
}
func TestLongestValidParentheses(t *testing.T) {
	Convey("[32] 最长有效括号\n", t, func() {
		So(longestValidParentheses("(())"), ShouldEqual, 4)
		So(longestValidParentheses("()()()"), ShouldEqual, 6)
		So(longestValidParentheses("(()()())"), ShouldEqual, 8)
		So(longestValidParentheses(")()())"), ShouldEqual, 4)
	})
}
