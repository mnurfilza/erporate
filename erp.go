package erporate

import (
	"strings"
)

func ReverseString(s string, k int) string {
	var res []string
	count := 0
	st := strings.Split(s, "")
	for i := 0; i <= len(st); i++ {
		count = count + 1
		if count == k {
			res = append(res, st[i])
		}
	}

	for _, it := range res {
		for _, item := range st {
			if it == item {
				continue
			}

			res = append(res, item)
		}

	}

	return strings.Join(res, "")
}
