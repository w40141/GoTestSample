package util

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	f.Add("hello")
	f.Fuzz(func(t *testing.T, org string) {
		rev, e1 := ReverseString(org)
		// 正しくエラーが返ってきた場合は何もしない
		if e1 != nil {
			return
		}

		doubleRev, e2 := ReverseString(rev)
		// 正しくエラーが返ってきた場合は何もしない
		if e2 != nil {
			return
		}

		// 一致しない場合はエラーを出力する
		if org != doubleRev {
			t.Errorf("expected: %q, got %q", org, doubleRev)
		}
		if utf8.ValidString(org) && !utf8.ValidString(rev) {
			t.Errorf("invalid string: %q", rev)
		}
	})
}
