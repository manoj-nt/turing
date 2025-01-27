// email_service.go
package emailservice

import (
	"errors"
	"fmt"
	"regexp"
)

// EmailService interface defines the method for sending emails.
type EmailService interface {
	SendTransactionalEmail(to string, body string) error
	SendPromotionalEmail(to string, body string) error
}

// BasicEmailService is a simple implementation of EmailService.
type BasicEmailService struct{}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// SendTransactionalEmail sends a transactional email.
func (s *BasicEmailService) SendTransactionalEmail(to string, body string) error {
	if !isValidEmail(to) {
		return errors.New("invalid email address")
	}

	// Simulate sending email
	fmt.Printf("Sent transactional email to %s: %s\n", to, body)
	return nil
}

// SendPromotionalEmail sends a promotional email.
func (s *BasicEmailService) SendPromotionalEmail(to string, body string) error {
	if !isValidEmail(to) {
		return errors.New("invalid email address")
	}

	// Simulate sending email
	fmt.Printf("Sent promotional email to %s: %s\n", to, body)
	return nil
}

// isValidEmail validates the email address format.
func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}
