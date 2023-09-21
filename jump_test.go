package suanfa

import "testing"

// 青蛙跳台阶
func TestJump(t *testing.T) {
	n := 50
	tmp := make(map[int]int)
	res := Jump(n, tmp)
	t.Log(res)
}

// 递归+优化
func Jump(n int, tmp map[int]int) int {
	if n < 3 {
		return n
	}
	if val, ok := tmp[n]; ok {
		return val
	}
	res := Jump(n-1, tmp) + Jump(n-2, tmp)
	tmp[n] = res
	return res
}
