package kmp

import (
	"strings"
	"testing"
)

/*
给定两个字符串, A 和 B。
A 的旋转操作就是将 A 最左边的字符移动到最右边。 例如, 若 A = ‘abcde’，在移动一次之后结果就是’bcdea’ 。
如果在若干次旋转操作之后，A 能变成B，那么返回True。
输入: A = 'abcde', B = 'cdeab'
输出: true
*/

func TestRoateString(t *testing.T) {
	a := `abcde`
	b := `cdeab`
	res := RoateString(a, b)
	t.Log(res)
}

/*
将str1扩大一倍，如果str2是newStr1的子集，那么
str2必然是str1旋转后的子集
*/
func RoateString(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	newStr1 := str1 + str1
	return strings.Contains(newStr1, str2)
}
