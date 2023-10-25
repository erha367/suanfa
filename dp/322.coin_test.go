package dp

import (
	"log"
	"testing"
)

/*
给你一个整数数组 coins ，表示不同面额的硬币；以及一个整数 amount ，表示总金额。
计算并返回可以凑成总金额所需的 最少的硬币个数 。如果没有任何一种硬币组合能组成总金额，返回 -1 。
你可以认为每种硬币的数量是无限的。
*/

func TestMinCoins(t *testing.T) {
	coins := []int{1, 5, 10, 25}
	amount := 40
	res := minCoins(coins, amount)
	t.Log(res)
}

func minCoins(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	log.Println(dp)
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
