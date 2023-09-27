package suanfa

import (
	"log"
	"testing"
)

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
请注意 ，必须在不复制数组的情况下原地对数组进行操作。
示例 1:
输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:
输入: nums = [0]
输出: [0]
*/

func TestMoveZero(t *testing.T) {
	arr := []int{0, 1, 0, 2, 0, 3, 0, 4}
	t.Log(MoveZero(arr))
	MoveZero2(arr)
	t.Log(arr)
}

// 创建一个数组，把非0的移过去，然后用0补全剩余的
func MoveZero(arr []int) []int {
	var res []int
	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			res = append(res, arr[i])
		}
	}
	for i := len(res); i < len(arr); i++ {
		res = append(res, 0)
	}
	return res
}

// 循环交换0和它后面的元素
func MoveZero2(nums []int) {
	k, length := 0, len(nums)
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			if i != k {
				nums[k], nums[i] = nums[i], nums[k]
				log.Println(nums)
			}
			k++
		}
	}

}
