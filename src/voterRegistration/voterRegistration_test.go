package voterRegistration

import (
	"errors"
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name              string
		voterRegistration string
		expectedError     error
	}{
		{
			name:              "valid voter registration",
			voterRegistration: "356061030159",
			expectedError:     nil,
		},
		{
			name:              "voter registration with invalid length",
			voterRegistration: "35606103015",
			expectedError:     errors.New("Voter Registration with invalid length"),
		},
		{
			name:              "voter registration with invalid UF",
			voterRegistration: "356061032959",
			expectedError:     errors.New("Invalid UF"),
		},
		{
			name:              "voter registration with invalid check digit 1",
			voterRegistration: "356061030119",
			expectedError:     errors.New("Invalid Voter Registration"),
		},
		{
			name:              "voter registration with invalid check digit 2",
			voterRegistration: "356061030150",
			expectedError:     errors.New("Invalid Voter Registration"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := IsValid(test.voterRegistration)
			if err == nil && test.expectedError == nil {
				return
			}

			if err.Error() != test.expectedError.Error() {
				t.Errorf("Expected error %v but got %v", test.expectedError, err)
			}
		})
	}
}
