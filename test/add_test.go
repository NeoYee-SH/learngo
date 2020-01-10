package main

import (
	"testing"

	mystring "github.com/yihuaiyuan/learngo/basic/strings"

	mymath "github.com/yihuaiyuan/learngo/basic/math"
)

func TestAdd(t *testing.T) {
	tests := []struct{ a, b, c int }{
		{3, 4, 7},
		{13, 14, 27},
		{23, 24, 47},
		{33, 34, 67},
		{43, 44, 87},
	}
	for _, tt := range tests {
		if ret := mymath.Add(tt.a, tt.b); ret != tt.c {
			t.Errorf("Add(%d, %d) got %d; expected %d;", tt.a, tt.b, ret, tt.c)
		}
	}
}
func BenchmarkStr(b *testing.B) {
	tests := []struct {
		s   string
		ret int
	}{
		{"what are you 弄啥咧?", 8},
		{"一二三二一", 3},
		{"abcancabd", 5},
	}

	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			if ret := mystring.LengthOfLongestSubstring(tt.s); ret != tt.ret {
				b.Errorf("LengthOfLongestSubstring(%s) got %d; expected %d;", tt.s, ret, tt.ret)
			}
		}
	}
}
