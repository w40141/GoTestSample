package util

import (
	"log"
	"os"
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestMain(m *testing.M) {
	// if err := setup(); err != nil {
	// 	log.Fatalf("setup failed: %v", err)
	// }
	log.Println("before all...")
	ret := m.Run()
	log.Println("after all...")
	// if err := teardown(); err != nil {
	// 	log.Fatalf("teardown failed: %v", err)
	// }
	os.Exit(ret)
}

// -shortオプションを付けてテストを実行すると、テストがスキップされる
func TestA(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	log.Println("testA")
}

func TestReverse(t *testing.T) {
	reverseTests := []struct {
		name     string
		input    string
		expected string
		isValid  bool
	}{
		{"アルファベット", "hello", "olleh", true},
		{"日本語", "こんにちは", "はちにんこ", true},
		{"空文字", "", "", true},
		{"絵文字", "🌟🚀", "🚀🌟", true},
		{"不正な値", "invalid\x80string", "", false},
	}

	for _, tt := range reverseTests {
		t.Run(tt.input, func(t *testing.T) {
			reversed, err := ReverseString(tt.input)

			if err != nil && tt.isValid {
				t.Errorf("Unexpected error: %v", err)
			}

			if err == nil && !tt.isValid {
				t.Errorf("Expected an error, but got none")
			}

			if reversed != tt.expected {
				t.Errorf("Expected: %s, Got: %s", tt.expected, reversed)
			}
		})
	}
}

func TestStringDistance(t *testing.T) {
	tests := []struct {
		name string
		lhs  string
		rhs  string
		want int
	}{
		{"lhsとrhsが同じ", "foo", "foo", 0},
		{"lhsがrhsより長い", "foo0", "foo", -1},
		{"rhsがlhsより長い", "foo", "fooo", -1},
		{"lhsとrhsが1文字異なる", "foo", "foa", 1},
		{"lhsとrhsが2文字異なる", "foo", "faa", 2},
		{"lhsとrhsが3文字異なる", "foo", "aaa", 3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := StringDistance(tt.lhs, tt.rhs)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("%s: expected: %v, got %v", tt.name, tt.want, got)
			}
		})
	}
}

func TestVeryHeavyFunction(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"aとbが同じ", 1, 1, 2},
		{"aがbより大きい", 2, 1, 3},
		{"bがaより大きい", 1, 2, 3},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := VeryHeavyFunction(tt.a, tt.b)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("%s: expected: %v, got %v", tt.name, tt.want, got)
			}
		})
	}
}

func FuzzReverse(f *testing.F) {
	f.Add("hello")
	f.Fuzz(func(t *testing.T, s string) {
		r, e1 := ReverseString(s)
		if !utf8.ValidString(s) && e1 == nil {
			// utf8.ValidString()がfalseを返す場合はエラーがあるはず
			t.Errorf("expected error")
		} else if e1 != nil {
			// 正しくエラーが返ってきた場合は何もしない
			return
		} else {
			o, _ := ReverseString(r)
			if o != s {
				t.Errorf("expected: %q, got %q", s, o)
			}
		}
	})
}
