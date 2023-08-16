package vo

import (
	"testing"
)

func TestNewUserId(t *testing.T) {
	tests := []struct {
		name           string
		userID         string
		expectedUserId UserId
		expectedErr    error
	}{
		{"ValidUserId", "john_doe123", UserId{value: "john_doe123"}, nil},
		{"InvalidUserIdStartingWithUnderscore", "_user123", UserId{}, InvalidUserIdError},
		{"InvalidUserIdWithNotAllowedCharacters", "j+ohndoe123", UserId{}, InvalidUserIdError},
		{"InvalidUserIdShorter3", "ab", UserId{}, InvalidUserIdError},
		{"InvalidUserIdLonger20", "123456789abcdefghijkl", UserId{}, InvalidUserIdError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := NewUserId(test.userID)
			if err != test.expectedErr {
				t.Errorf("Expected no error, but got %v", err)
			}
			if actual != test.expectedUserId {
				t.Errorf("Expected userID=%v, but got %v", test.expectedUserId, actual)
			}
			if err == nil && actual.Value() != test.userID {
				t.Errorf("Expected userID=%v, but got %v", test.userID, actual.Value())
			}
		})
	}
}

func TestIsValidUserID(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		expected bool
	}{
		{"ValidUserId", "john_doe123", true},
		{"InvalidUserIdStartingWithUnderscore", "_user123", false},
		{"InvalidUserIdWithNotAllowedCharacters", "j+ohndoe123", false},
		{"InvalidUserIdShorter3", "ab", false},
		{"InvalidUserIdLonger20", "123456789abcdefghijkl", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isValidUserID(test.userID)
			if actual != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, actual)
			}
		})
	}
}

func TestCheckUserIdCharacters(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		expected bool
	}{
		{"ValidUserId", "john_doe123", true},
		{"ValidUserIdShorter3", "ab", true},
		{"ValidUserIdLonger20", "123456789abcdefghijkl", true},
		{"InvalidUserIdStartingWithUnderscore", "_user123", false},
		{"InvalidUserIdWithNotAllowedCharacters", "j+ohndoe123", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := checkUserIdCharacters(test.userID)
			if actual != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, actual)
			}
		})
	}
}

func TestCheckUserIdLength(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		expected bool
	}{
		{"ValidUserId", "john_doe123", true},
		{"ValidUserIdStartingWithUnderscore", "_user123", true},
		{"ValidUserIdWithNotAllowedCharacters", "j+ohndoe123", true},
		{"InvalidUserIdShorter3", "ab", false},
		{"InvalidUserIdLonger20", "123456789abcdefghijkl", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := checkUserIdLength(test.userID)
			if actual != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, actual)
			}
		})
	}
}

func TestGerateNewUserId(t *testing.T) {
	t.Run("GenerateNewUserId", func(t *testing.T) {
		actual := GenerateNewUserId()
		if len(actual.Value()) != 10 {
			t.Errorf("Expected: %v, Got: %v", 10, len(actual.Value()))
		}
	})
}

func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"ValidLength", 10},
		{"InvalidZeroLength", 0},
		{"InvalidMinus", -1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := generateRandomString(test.length)
			if test.length < 0 {
				if actual != "" {
					t.Errorf("Expected empty string, but got %v", actual)
				}
			} else {
				if len(actual) != test.length {
					t.Errorf("Expected length=%v, but got %v", test.length, len(actual))
				}
			}
		})
	}
}
