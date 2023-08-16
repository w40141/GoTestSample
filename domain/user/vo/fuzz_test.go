package vo

import (
	"testing"
)

func FuzzGenerateUserId(f *testing.F) {
	f.Fuzz(func(t *testing.T, i int) {
		if i < 3 || i > 20 {
			return
		}
		s := generateRandomString(i)
		if len(s) != i {
			t.Errorf("Expected %d, got %d", i, len(s))
		}
		actural, err := NewUserId(s)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if actural.Value() != s {
			t.Errorf("Expected %v, got %v", s, actural.Value())
		}
	})
}
