package string

import "strings"

type StrUtils struct {
}

func (*StrUtils) CapitalizeFirst(s string) string {
	if s == "" {
		return s
	}
	// 将首字母转为大写，其余保持不变
	return strings.ToUpper(string(s[0])) + s[1:]
}
