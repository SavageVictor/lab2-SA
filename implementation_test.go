package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvaluatePostfix(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expected      float64
		expectedError string
	}{
		{
			name:     "3 operations",
			input:    "3 9 + 54 * 32 +",
			expected: 680,
		},
		{
			name:     "10 operations",
			input:    "10 17 23 * 3 / + 2 + 1 - 4 3 ^ + 1 6 * 9 * -",
			expected: 151.3,
		},
		{
			name:          "empty input",
			input:         "",
			expectedError: "invalid postfix expression",
		},
		{
			name:          "wrong symbols",
			input:         "3 4 + x 2 * 1 +",
			expectedError: "invalid postfix expression",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := EvaluatePostfix(tc.input)

			if tc.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
				assert.InDelta(t, tc.expected, result, 0.1) // Use InDelta for comparing float values
			}
		})
	}
}
