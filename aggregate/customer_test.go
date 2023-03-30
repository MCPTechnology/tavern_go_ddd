package aggregate_test

import (
	"testing"

	"github.com/matheuscaputopires/ddd-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	// Build our needed test case data struct
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	// Create test cases
	testCases := []testCase{
		{
			test:        "Empty Name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, {
			test:        "Valid Name validation",
			name:        "Percy Bolmer",
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		// Run tests for each scenario
		t.Run(tc.test, func(t *testing.T) {
			// Create a new Customer
			_, err := aggregate.NewCustomer(tc.name)
			// Check if received error matches expected error
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
