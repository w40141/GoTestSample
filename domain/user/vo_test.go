package user

import (
	"testing"
)

func TestIsValidUserID(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		expected bool
	}{
		{"ValidUserId", "john_doe123", true},
		{"InvalidUserIdStartingWithUnderscore", "_user123", false},
		{"InvalidUserIdWithNotAllowedCharacters", "j+ohndoe123", false},
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
		{"InvalidUserIdStartingWithUnderscore", "_user123", false},
		{"InvalidUserIdWithNotAllowedCharacters", "j+ohndoe123", false},
		{"InvalidUserIdWithEmpty", "", false},
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
		{"ValidLengthUserID", "abc", true},
		{"ValidLengthUserID", "123456789abcdefghijk", true},
		{"InvalidLengthUserID", "ab", false},
		{"InvalidLengthUserID", "123456789abcdefghijkl", false},
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

func TestNewUserId(t *testing.T) {
	tests := []struct {
		name     string
		userID   string
		expected bool
	}{
		{"ValidUserId", "john_doe123", true},
		{"InvalidUserIdStartingWithUnderscore", "_user123", false},
		{"InvalidUserIdWithNotAllowedCharacters", "j+ohndoe123", false},
		{"InvalidShortUserId", "ab", false},
		{"InvalidLongUserId", "123456789abcdefghijkl", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			userID, err := NewUserId(test.userID)
			if test.expected {
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
				if userID.Value() != test.userID {
					t.Errorf("Expected userID=%v, but got %v", test.userID, userID.Value())
				}
			} else {
				if err == nil {
					t.Errorf("Expected error, but got no error")
				}
			}
		})
	}
}

func TestNewUserName(t *testing.T) {
	testCases := []struct {
		name          string
		firstName     string
		lastName      string
		expectedError error
	}{
		{"ValidUserName", "John", "Doe", nil},
		{"InvalidUserNameWithEmptyFirstName", "", "Doe", InvalidNameError},
		{"InvalidUserNameWithEmptyLastName", "John", "", InvalidNameError},
		{"InvalidUserNameWithNumber", "123", "Doe", InvalidNameError},
		{"InvalidUserNameWithNotAllowedCharacter", "_Alice", "Smith", InvalidNameError},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := NewUserName(tc.firstName, tc.lastName)
			if err != nil {
				if err != tc.expectedError {
					t.Errorf("Expected no error, but got %v", err)
				}
			} else {
				if actual.FirstName() != tc.firstName {
					t.Errorf("Expected value %v but got %v", tc.firstName, actual.FirstName())
				}
				if actual.LastName() != tc.lastName {
					t.Errorf("Expected value %v but got %v", tc.lastName, actual.LastName())
				}
				if actual.FullName() != tc.firstName+" "+tc.lastName {
					t.Errorf("Expected value '%v %v' but got %v", tc.firstName, tc.lastName, actual.FullName())
				}
			}
		})
	}
}

func TestIsValidName(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"ValidName", "validName", true},
		{"InValidNameWithNotAllowedCharacters", "invalid_name_12345678901", false},
		{"InValidNameWithSpace", " ", false},
		{"InValidNameWithEmpty", "", false},
		{"InValidNameWithNumber", "123", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isValidName(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, actual)
			}
		})
	}
}

func TestGenerateRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"ValidLength", 10},
		{"InvalidLength", -1},
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
				_, err := NewUserId(actual)
				if err != nil {
					t.Errorf("Expected no error, but got %v", err)
				}
			}
		})
	}
}

func TestIsValidAge(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"InvalidMinus", -1, false},
		{"ValidZero", 0, true},
		{"ValidNumber", 17, true},
		{"ValidMaxNumber", 150, true},
		{"InValidOver", 151, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := isValidAge(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %v but got %v", tc.expected, actual)
			}
		})
	}
}
