package suanfa

import "testing"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。

示例 2：
输入：nums = [3,2,4], target = 6
输出：[1,2]
*/

// 两数之和(https://leetcode-cn.com/problems/two-sum/)
func TestTowSum(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	sum := 14
	x := TwoSum(arr, sum)
	t.Log(x)
}

// 空间换时间  key是arr的值，value是坐标，然后判断key是否存在
func TwoSum(arr []int, sum int) []int {
	arrMap := make(map[int]int)
	var res []int
	for k, v := range arr {
		if val, ok := arrMap[sum-v]; ok {
			res = append(res, val)
			res = append(res, k)
			return res
		}
		arrMap[v] = k
	}
	return []int{-1, -1}
}
