package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	testCases := []struct {
		description   string
		headers       http.Header
		expected      string
		expectedError error
	}{
		{
			description:   "No Authorization Header",
			headers:       http.Header{},
			expected:      "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			description:   "Malformed Authorization Header",
			headers:       http.Header{"Authorization": []string{"Bearer token123"}},
			expected:      "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			description:   "Valid Authorization Header",
			headers:       http.Header{"Authorization": []string{"ApiKey validApiKey123"}},
			expected:      "validApiKey123",
			expectedError: nil,
		},
	}
	// looping through the test cases

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result, err := GetAPIKey(tc.headers)
			if result != tc.expected || (err != nil && err.Error() != tc.expectedError.Error()) {
				t.Errorf("expected result: %v, err: %v; got result: %v, err: %v", tc.expected, tc.expectedError, result, err)
			}
		})
	}
}
