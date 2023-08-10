package user

import (
	"testing"
)

func TestIsValidUserId(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		expected bool
	}{
		{"正しいId", "john_doe123", true},
		{"先頭が_で正しくないID", "_user123", false},
		{"許可されていない記号を使ったID", "j+ohndoe123", false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result := isValidUserID(test.userID)
			if result != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, result)
			}
		})
	}
}

func TestIsValidLengthUserID(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		expected bool
	}{
		{"3文字のID", "abc", true},
		{"20文字のID", "123456789abcdefghijk", true},
		{"2文字のID", "ab", false},
		{"21文字のID", "123456789abcdefghijkl", false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			result := isValidLengthUserID(test.userID)
			if result != test.expected {
				t.Errorf("Expected: %v, Got: %v", test.expected, result)
			}
		})
	}
}

func TestUserId(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		expected bool
	}{
		{"正しいId", "john_doe123", true},
		{"3文字のID", "abc", true},
		{"20文字のID", "123456789abcdefghijk", true},
		{"先頭が_で正しくないID", "_user123", false},
		{"許可されていない記号を使ったID", "j+ohndoe123", false},
		{"2文字のID", "ab", false},
		{"21文字のID", "123456789abcdefghijkl", false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			_, e := NewUserId(test.userID)
			isValid := e == nil
			if isValid != test.expected {
				t.Errorf("Expected isValid=%v, but got isValid=%v", test.expected, isValid)
			}
		})
	}
}

func TestUserName(t *testing.T) {
	tests := []struct {
		name      string
		firstName string
		lastName  string
		isValid   bool
		fullName  string
	}{
		{"ValidName", "John", "Doe", true, "John Doe"},
		{"InvalidName", "John123", "Doe", false, ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userName, err := NewUserName(test.firstName, test.lastName)
			if test.isValid {
				assert.NoError(t, err)
				assert.Equal(t, test.firstName, userName.FirstName())
				assert.Equal(t, test.lastName, userName.LastName())
				assert.Equal(t, test.fullName, userName.FullName())
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestAge(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected bool
	}{
		{"ValidAge", 30, true},
		{"InvalidNegativeAge", -10, false},
		{"InvalidHighAge", 150, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NewAge(test.age)
		})
	}
}

func TestGenerateUserID(t *testing.T) {
	userID := GenerateUserID()
	assert.True(t, isValidUserID(userID.Value()))
	assert.True(t, isValidLengthUserID(userID.Value()))
}
