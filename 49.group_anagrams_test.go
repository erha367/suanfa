package suanfa

import (
	"sort"
	"strings"
	"testing"
)

/*
字母异位词分组
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
举个例子，比如给定的数组是[eat, ate, tea, tan, nat, bat]。其中eat，ate，tea这三个单词用到的字母都是e，t和a各一个。
tan和nat用到的都是a，n和t，最后剩下bat，所以分组结果就是：[eat, ate, tea]，[tan, nat]和[bat]。
*/

func Test(t *testing.T) {
	x := `eta`
	t.Log(SortStr(x))
	a := []string{`eat`, `ate`, `tea`, `tan`, `nat`, `bat`}
	t.Log(GroupAna(a))
}

func GroupAna(arr []string) [][]string {
	mp := make(map[string][]string)
	for _, v := range arr {
		key := SortStr(v)
		mp[key] = append(mp[key], v)
	}
	var res [][]string
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func SortStr(a string) string {
	arr := strings.Split(a, "")
	sort.Strings(arr)
	return strings.Join(arr, "")
}
