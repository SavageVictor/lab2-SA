package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComputeHandler(t *testing.T) {
	testCases := []struct {
		name          string
		input         string
		expected      string
		expectedError string
	}{
		{
			name:     "Valid expression",
			input:    "3 4 + 2 * 1 + 2 ^",
			expected: "The result of the postfix expression '3 4 + 2 * 1 + 2 ^' is: 225.00\n",
		},
		{
			name:          "Invalid expression",
			input:         "3 4 + 2 * 1 + 2 ^ x",
			expectedError: "error evaluating expression: invalid postfix expression",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputReader := strings.NewReader(tc.input)
			outputBuffer := &bytes.Buffer{}

			handler := &ComputeHandler{
				Input:  inputReader,
				Output: outputBuffer,
			}

			err := handler.Compute()

			if tc.expectedError != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, outputBuffer.String())
			}
		})
	}
}
