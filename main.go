// packageの上に文章を書くとdocのOverviewに表示されます
package main

import "time"

// importの書き方
func StringDistance(lhs, rhs string) int {
	return Distance([]rune(lhs), []rune(rhs))
}

// importの書き方
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

func VeryHeavyFunction(a, b int) int {
	// 重い処理
	resutl := a + b
	time.Sleep(3 * time.Second)
	return resutl
}
