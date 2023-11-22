package binsearch

import "testing"

func TestBinSearch(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	t.Log(BinSearch(a, 2))
}

func BinSearch(a []int, target int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == target {
			return mid
		} else if a[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
