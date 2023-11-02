package greedy

import (
	"log"
	"testing"
)

func TestPlantFlower(t *testing.T) {
	a := []int{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0}
	n := 4
	res := CanPlantFlower(a, n)
	t.Log(res)
}

func CanPlantFlower(a []int, n int) bool {
	if len(a) == 0 {
		return false
	}

	b := make([]int, len(a)+2)
	b[0] = 0
	for i := 1; i < len(b)-1; i++ {
		b[i] = a[i-1]
	}
	b[len(b)-1] = 0
	dp := make([]int, len(b))
	for i := 1; i < len(b)-1; i++ {
		if b[i-1]+b[i]+b[i+1] == 0 {
			dp[i] = 1
			b[i] = 1
		}
	}
	log.Println(dp)
	var sum int
	for _, v := range dp {
		sum += v
	}
	return sum > n
}
