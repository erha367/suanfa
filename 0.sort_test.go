package suanfa

import (
	"log"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{10, 2, 30, 4, 50, 6, 70, 8}
	res := QuickSort(arr)
	t.Log(res)
	arr2 := []int{10, 2, 30, 4, 50, 6, 70, 8}
	bubbleSort(arr2)
	t.Log(arr2)
}

func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	flags := arr[0]
	var l, r, ls, rs []int
	for i := 1; i < len(arr); i++ {
		if arr[i] > flags {
			r = append(r, arr[i])
		} else {
			l = append(l, arr[i])
		}
	}
	ls = QuickSort(l)
	ls = append(ls, flags)
	rs = QuickSort(r)
	return append(ls, rs...)
}

func bubbleSort(arr []int) {
	n := len(arr)
	// 外层循环控制比较的轮数
	for i := 0; i < n-1; i++ {
		// 内层循环进行相邻元素比较并交换
		for j := 0; j < n-i-1; j++ {
			log.Println(i, j)
			if arr[j] > arr[j+1] {
				// 交换 arr[j] 和 arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			log.Println(arr)
		}
	}
}
