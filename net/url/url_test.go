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

func TestResolveLink(t *testing.T) {
	baseLink := "https://example.com"
	testCases := []struct {
		link        string
		expected    string
		shouldError bool
	}{
		// Absolute links
		{"https://example.com", "https://example.com", false},
		{"https://google.com", "https://google.com", false},
		{"https://stackoverflow.com", "https://stackoverflow.com", false},

		// Relative links
		{"relative/link1", "https://example.com/relative/link1", false},
		{"/relative/link2", "https://example.com/relative/link2", false},

		// Invalid URLs
		{":invalid:", "", true},
	}

	for _, tc := range testCases {
		resolvedLink, err := ResolveLink(baseLink, tc.link)
		if tc.shouldError && err == nil {
			t.Errorf("Expected an error for invalid URL %s", tc.link)
		}
		if !tc.shouldError && err != nil {
			t.Errorf("Unexpected error for URL %s: %v", tc.link, err)
		}
		if resolvedLink != tc.expected {
			t.Errorf("Expected %s, got %s for link %s", tc.expected, resolvedLink, tc.link)
		}
	}
}
