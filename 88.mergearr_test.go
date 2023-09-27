package suanfa

import (
	"testing"
)

//给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

func TestMergeArray(t *testing.T) {
	a := []int{1, 3, 5, 7, 10, 20}
	b := []int{2, 4, 60}
	res := MergeArray(a, b)
	t.Log(res)
}

func MergeArray(a []int, b []int) []int {
	var res []int //结果集
	var i, j int  //双指针
	//备注：数组长度相同的话，可以完全合并
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	//如果a有剩余，补充到res里
	for i < len(a) {
		res = append(res, a[i])
		i++
	}
	for j < len(b) {
		res = append(res, b[j])
		j++
	}
	return res
}
