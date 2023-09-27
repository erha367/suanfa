package suanfa

import (
	"log"
	"testing"
)

/*
给定一个范围在 1 ≤ a[i] ≤ n (n = 数组大小) 的整型数组，数组中的一些元素重复出现了两次，而其他元素则完全没有出现过。
找到所有在 [1, n] 范围之间没有出现在数组中的数字。
*/

func TestNoExistNum(t *testing.T) {
	arr := []int{1, 1, 2, 2, 5}
	res := NoExistNum(arr)
	t.Log(res)
}

func NoExistNum(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	var res []int
	for _, v := range nums {
		index := AbsInt(v) - 1
		if nums[index] > 0 {
			nums[index] = 0 - nums[index]
		}
	}
	log.Println(nums)
	for k, v := range nums {
		if v > 0 {
			res = append(res, k+1)
		}
	}
	return res
}

func AbsInt(n int) int {
	if n >= 0 {
		return n
	}
	return 0 - n
}
