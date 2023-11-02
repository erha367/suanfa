package donggui

import (
	"log"
	"testing"
)

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级……它也可以跳上n级。求该青蛙跳上一个n级的台阶总共有多少种跳法。
求斐波那契数列第N项
*/

// 青蛙跳台阶
func TestJump(t *testing.T) {
	n := 5
	res := Jump(n)
	t.Log(res)
}

// 动态规划
func Jump(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	if n <= 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	log.Println(dp)
	return dp[n]
}
