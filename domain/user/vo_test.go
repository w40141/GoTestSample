package user

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
		{"InvalidShortUserId", "ab", UserId{}, InvalidUserIdError},
		{"InvalidLongUserId", "123456789abcdefghijkl", UserId{}, InvalidUserIdError},
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

func TestGerateUserId(t *testing.T) {
	t.Run("GenerateUserId", func(t *testing.T) {
		actual := GenerateUserId()
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
			}
		})
	}
}

func TestNewUserName(t *testing.T) {
	tests := []struct {
		name             string
		firstName        string
		lastName         string
		expectedUserName UserName
		expectedErr      error
	}{
		{"ValidUserName", "John", "Doe", UserName{firstName: "John", lastName: "Doe"}, nil},
		{"InvalidUserNameWithEmptyFirstName", "", "Doe", UserName{}, InvalidNameError},
		{"InvalidUserNameWithEmptyLastName", "John", "", UserName{}, InvalidNameError},
		{"InvalidUserNameWithNumber", "123", "Doe", UserName{}, InvalidNameError},
		{"InvalidUserNameWithNotAllowedCharacter", "_Alice", "Smith", UserName{}, InvalidNameError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := NewUserName(test.firstName, test.lastName)
			if err != test.expectedErr {
				t.Errorf("Expected no error, but got %v", err)
			}
			if actual != test.expectedUserName {
				t.Errorf("Expected userUserName=%v, but got %v", test.expectedUserName, actual)
			}
			if err == nil {
				if actual.FirstName() != test.firstName {
					t.Errorf("Expected value %v but got %v", test.firstName, actual.FirstName())
				}
				if actual.LastName() != test.lastName {
					t.Errorf("Expected value %v but got %v", test.lastName, actual.LastName())
				}
				if actual.FullName() != test.firstName+" "+test.lastName {
					t.Errorf("Expected value '%v %v' but got %v", test.firstName, test.lastName, actual.FullName())
				}
			}
		})
	}
}

func TestIsValidName(t *testing.T) {
	tests := []struct {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isValidName(test.input)
			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}
}

func TestNewUserAge(t *testing.T) {
	tests := []struct {
		name        string
		age         int
		expectedAge UserAge
		expectedErr error
	}{
		{"ValidAge", 20, UserAge{value: 20}, nil},
		{"InvalidAge", 151, UserAge{}, InvalidAgeError},
		{"InvalidAge", -1, UserAge{}, InvalidAgeError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := NewUserAge(test.age)
			if err != test.expectedErr {
				t.Errorf("Expected no error, but got %v", err)
			}
			if actual != test.expectedAge {
				t.Errorf("Expected userAge=%v, but got %v", test.expectedAge, actual)
			}
			if err == nil && actual.Value() != test.age {
				t.Errorf("Expected value %v but got %v", test.age, actual.Value())
			}
		})
	}
}

func TestIsValidAge(t *testing.T) {
	tests := []struct {
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

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := isValidAge(test.input)
			if actual != test.expected {
				t.Errorf("Expected %v but got %v", test.expected, actual)
			}
		})
	}
}
