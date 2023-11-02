package donggui

import (
	"log"
	"testing"
)

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。


示例 1：

输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。
*/

func TestLongestIncreasingSubsequence(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101}
	res := LongestIncreasingSubsequence(nums)
	t.Log(res)
}

// 求最长递增子序列长度
func LongestIncreasingSubsequence(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	var big int
	for i := 1; i < len(nums); i++ {
		//默认长度为1
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
				if dp[i] > big {
					big = dp[i]
				}
			}
		}
	}
	log.Println(dp)
	return big
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
