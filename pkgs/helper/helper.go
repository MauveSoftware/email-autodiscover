package helper

import "strings"

// ExtractDomainFromEmail extracts the domain part from an email address.
// It returns the domain (everything after the @ symbol) if the email is valid.
// Returns an empty string if the email doesn't contain exactly one @ symbol.
func ExtractDomainFromEmail(email string) string {
	parts := strings.Split(email, "@")

	if len(parts) == 2 {
		return parts[1]
	}

	return ""
}
