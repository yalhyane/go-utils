package url_utils

import (
	"testing"
)

func TestGetUrlDomain(t *testing.T) {
	// Define test cases.
	testCases := []struct {
		url, expected string
	}{
		{"https://www.google.com", "google.com"},
		{"https://example.com", "example.com"},
		{"http://www.subdomain.example.com", "subdomain.example.com"},
		{"https://localhost", "localhost"},
		{"", ""},
	}

	// Run each test case.
	for _, tc := range testCases {
		actual := GetUrlDomain(tc.url)

		// Assert that the expected output is produced.
		if actual != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, actual)
		}
	}
}
