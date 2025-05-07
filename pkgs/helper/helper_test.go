package helper

import "testing"

func TestExtractDomainFromEmail(t *testing.T) {
	testCases := []struct {
		email    string
		expected string
	}{
		{"user@example.com", "example.com"},
		{"another.user@test-domain.co.uk", "test-domain.co.uk"},
		{"john.doe123@gmail.com", "gmail.com"},
		{"info@my-company.org", "my-company.org"},
		{"no-reply@subdomain.example.net", "subdomain.example.net"},
		{"user+tag@domain.com", "domain.com"},
		{"", ""},
		{"invalid-email", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.email, func(t *testing.T) {
			result := ExtractDomainFromEmail(tc.email)
			if result != tc.expected {
				t.Errorf("ExtractDomainFromEmail(%q) = %q, want %q", tc.email, result, tc.expected)
			}
		})
	}
}
