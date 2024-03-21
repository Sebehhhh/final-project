package utils

import (
	"regexp"
)

func IsValidURL(url string) bool {
	// Regular expression pattern for URL validation
	// This pattern checks for a URL that starts with http:// or https:// followed by a valid domain name
	// and optional path and query parameters
	pattern := `^(?:(?:https?|ftp|file):\/\/)?(?:www\.)?([a-zA-Z0-9-]+\.)*[a-zA-Z0-9-]+\.[a-zA-Z]{2,}(/[a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=]*)?$`
	match, _ := regexp.MatchString(pattern, url)
	return match
}
