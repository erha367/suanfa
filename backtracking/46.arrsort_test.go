package backtracking

import (
	"log"
	"testing"
)

func TestQuanSort(t *testing.T) {
	demo := []int{1, 2, 3, 4}
	t.Log(permute(demo))
}

func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, 0, &result)
	return result
}

func backtrack(nums []int, start int, result *[][]int) {
	if start == len(nums)-1 {
		// 将当前排列添加到结果中
		temp := make([]int, len(nums))
		copy(temp, nums)
		log.Println(temp)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		log.Println(i, start)
		// 交换当前元素与后面的元素
		nums[start], nums[i] = nums[i], nums[start]
		// 递归求解下一个位置的排列
		backtrack(nums, start+1, result)
		// 恢复数组原始顺序
		nums[start], nums[i] = nums[i], nums[start]
	}
}
