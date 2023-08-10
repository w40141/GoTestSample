// stringに関する処理をまとめたパッケージ
package util

import (
	"errors"
	"time"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

// 文字列の距離を返す
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// runeのスライスの距離を返す
func Distance(a, b []rune) int {
	dist := 0
	if len(a) != len(b) {
		return -1
	}
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}

// 重い処理
func VeryHeavyFunction(a, b int) int {
	result := a + b
	time.Sleep(3 * time.Second)
	return result
}

// 文字列を反転させる
func ReverseString(s string) (string, error) {
	if !utf8.ValidString(s) {
		return "", ErrInvalidString
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		// ここでswapしている
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil
}
