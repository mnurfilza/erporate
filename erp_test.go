package erporate

import (
	"fmt"
	"testing"
)

func TestErp(t *testing.T) {
	h := "abcdefg"
	l := 2

	st := ReverseString(h, l)
	fmt.Println(st)
}
