package main

import (
	"log"
	"os"
	"reflect"
	"testing"
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
		got := StringDistance(tt.lhs, tt.rhs)
		if !reflect.DeepEqual(tt.want, got) {
			t.Fatalf("%s: expected: %v, got %v", tt.name, tt.want, got)
		}
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
		test := tt
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := VeryHeavyFunction(test.a, test.b)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("%s: expected: %v, got %v", test.name, test.want, got)
			}
		})
	}
}
