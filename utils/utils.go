package utils

import (
	"log"
	"strings"
)

// ParseEmails parses a string of emails and returns a slice of emails
func ParseEmails(input string) []string {
	var emails []string
	parts := strings.Split(input, " ")
	log.Println(parts)
	for _, part := range parts {
			if strings.HasPrefix(part, "@") {
					email := strings.TrimPrefix(part, "@")
					emails = append(emails, email)
			}
	}
	return emails
}