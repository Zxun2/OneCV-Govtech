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

// CreateListOfRandomEmails creates list of random emails 
func CreateListOfRandomEmails(n int) []string {
	var emails []string
	for i := 0; i < n; i++ {
			emails = append(emails, RandomEmail())
	}
	return emails
}

// RemoveDuplicates removes duplicate emails
func RemoveDuplicates	(input []string) []string {
	keys := make(map[string]bool)
	list := []string{} 
	for _, entry := range input {
			if _, value := keys[entry]; !value {
					keys[entry] = true
					list = append(list, entry)
			}
	}
	return list
}