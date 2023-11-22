package basic

import (
	"errors"
	"testing"
)

func TestSecondMax(t *testing.T) {
	a := []int{1, 2, 3, 4}
	t.Log(SecondMax(a))
}

func SecondMax(a []int) (int, error) {
	if len(a) < 2 {
		return 0, errors.New("array length less than 2")
	}
	var m, idx int
	for i := 0; i < len(a); i++ {
		if a[i] > m {
			m = a[i]
			idx = i
		}
	}
	var n int
	for i := 0; i < len(a); i++ {
		if a[i] > n && i != idx {
			n = a[i]
		}
	}
	return n, nil
}
