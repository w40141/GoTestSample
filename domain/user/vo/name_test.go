package vo

import (
	"testing"
)

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
