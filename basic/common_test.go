package basic

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

// 判断两个字符串中每个字母出现的次数是否相同
func TestCheckSameNumMap(t *testing.T) {
	x := CheckLetterFrequency(`aabbbcccdd`, `ddcccaabbb`)
	t.Log(x)
}

func CheckLetterFrequency(str1, str2 string) bool {
	// 创建一个map用于存储每个字母的出现次数
	letterCount1 := make(map[rune]int)
	letterCount2 := make(map[rune]int)

	// 遍历第一个字符串，统计每个字母的出现次数
	for _, char := range str1 {
		letterCount1[char]++
	}
	log.Println(letterCount1)
	// 遍历第二个字符串，统计每个字母的出现次数
	for _, char := range str2 {
		letterCount2[char]++
	}
	log.Println(letterCount2)
	// 检查字母出现次数的映射是否相同
	return reflect.DeepEqual(letterCount1, letterCount2)
}

// Go 函数，用于统计字符串中每个字母出现的次数，并按照 a-z 排序返回结果

func TestCountLetters(t *testing.T) {
	str := "zzbbbcccaa"
	result := CountLetters(str)
	fmt.Println(result)
}

func CountLetters(str string) string {
	// 创建一个长度为26的切片，用于存储每个字母的出现次数
	letterCount := make([]int, 26)

	// 统计每个字母的出现次数
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			letterCount[char-'a']++
		}
	}
	log.Println(letterCount)
	// 构建结果字符串
	result := ""
	for i := 0; i < 26; i++ {
		if letterCount[i] > 0 {
			result += fmt.Sprintf("%c%d", 'a'+i, letterCount[i])
		}
	}

	return result
}

/*
reflect.DeepEqual() 是 Go 语言中的一个函数，用于比较两个值是否深度相等。它的函数签名如下：
func DeepEqual(a1, a2 interface{}) bool
DeepEqual() 函数接受两个参数 a1 和 a2，这两个参数可以是任意类型的值。函数会递归地比较这两个值及其内部的成员，以确定它们是否深度相等。
*/

func TestDeepEqual(t *testing.T) {
	// 示例1
	str1 := "hello"
	str2 := "hello"
	result1 := reflect.DeepEqual(str1, str2)
	fmt.Println(result1) // 输出: true

	// 示例2
	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	result2 := reflect.DeepEqual(slice1, slice2)
	fmt.Println(result2) // 输出: true

	// 示例3
	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 2, "a": 1}
	result3 := reflect.DeepEqual(map1, map2)
	fmt.Println(result3) // 输出: true

	// 示例4
	type Person struct {
		Name string
		Age  int
	}
	person1 := Person{Name: "Alice", Age: 30}
	person2 := Person{Name: "Alice", Age: 30}
	result4 := reflect.DeepEqual(person1, person2)
	fmt.Println(result4) // 输出: true

}

func TestIsSubArray(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	b := []int{3, 4, 5}
	res := isSubArray(a, b)
	t.Log(res)
}

func isSubArray(nums []int, target []int) bool {
	var i, j int
	for i < len(nums) && j < len(target) {
		if nums[i] != target[j] {
			i++
		} else {
			j++
			i++
		}

	}
	if j == len(target) {
		return true
	}
	return false
}
