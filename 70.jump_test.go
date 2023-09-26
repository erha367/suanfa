package suanfa

import "testing"

// 青蛙跳台阶
func TestJump(t *testing.T) {
	n := 5
	tmp := make(map[int]int)
	res := Jump(n, tmp)
	t.Log(res)
	res2 := Jump2(n)
	t.Log(res2)
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

// 递归改循环
func Jump2(n int) int {
	if n < 3 {
		return n
	}
	//保存结果
	res := 0
	//中间过程f(2)
	pre := 2
	//中间过程f(1)
	prepre := 1
	//3~n的数字区间
	for i := 3; i < n+1; i++ {
		res = pre + prepre
		prepre = pre
		pre = res
	}
	return res
}
