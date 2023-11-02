package greedy

import (
	"log"
	"testing"
)

/*
一群孩子站成一排，每一个孩子有自己的评分。现在需要给这些孩子发糖果，规则是如果一
个孩子的评分比自己身旁的一个孩子要高，那么这个孩子就必须得到比身旁孩子更多的糖果；所
有孩子至少要有一个糖果。求解最少需要多少个糖果。
*/

func TestSendCandies(t *testing.T) {
	a := []int{1, 3, 4, 5, 2}
	res := candy(a)
	t.Log(res)
	b := []int{1, 2, 87, 87, 2, 1}
	res2 := candy(b)
	t.Log(res2)
}

func candy(ratings []int) int {
	n := len(ratings)
	if n <= 1 {
		return n
	}
	candies := make([]int, n)
	for i := 0; i < n; i++ {
		candies[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	log.Println(candies)
	for i := n - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			candies[i-1] = max(candies[i-1], candies[i]+1)
		}
	}
	log.Println(candies)
	sum := 0
	for i := 0; i < n; i++ {
		sum += candies[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
