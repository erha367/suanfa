package twopointer

import (
	"testing"
)

//在一个增序的整数数组里找到两个数，使它们的和为给定值。已知有且只有一对解。

func TestTwoSum2(t *testing.T) {
	t.Log(twoSum2([]int{2, 7, 11, 15}, 9))
}

func twoSum2(nums []int, target int) (int, int) {
	//定义两个指针，指向数组的首尾
	i, j := 0, len(nums)-1
	for i != j {
		if nums[i]+nums[j] == target {
			return i, j
		}
		//和比目标小，左指针右移
		if nums[i]+nums[j] < target {
			i++
		}
		//和比目标大，右指针左移
		if nums[i]+nums[j] > target {
			j--
		}
	}
	return -1, -1
}
