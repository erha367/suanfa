package suanfa

import (
	"testing"
)

func TestQuanSort(t *testing.T) {
	demo := []int{1, 2, 3, 4}
	t.Log(permute(demo))
}

func permute(nums []int) [][]int {
	var result [][]int
	var used = make([]bool, len(nums))
	var current []int
	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			result = append(result, append([]int{}, current...))
			return
		}
		for i, v := range nums {
			if used[i] {
				continue
			}
			used[i] = true
			current = append(current, v)
			backtrack()
			used[i] = false
			current = current[:len(current)-1]
		}
	}
	backtrack()
	return result
}
