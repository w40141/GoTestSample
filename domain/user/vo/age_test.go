package vo

import (
	"testing"
)

func TestNewUserAge(t *testing.T) {
	tests := []struct {
		name        string
		age         int
		expectedAge UserAge
		expectedErr error
	}{
		{"ValidAge", 20, UserAge{value: 20}, nil},
		{"InvalidAgeOver150", 151, UserAge{}, InvalidAgeError},
		{"InvalidAgeUnder0", -1, UserAge{}, InvalidAgeError},
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
		{"ValidZero", 0, true},
		{"ValidNumber", 17, true},
		{"ValidMaxNumber", 150, true},
		{"InvalidUnder0", -1, false},
		{"InValidOver150", 151, false},
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
