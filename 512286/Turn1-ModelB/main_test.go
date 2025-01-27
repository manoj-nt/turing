package emailservice

import (
	"errors"
	"testing"
)

type EmailService struct {
	// Implement your email service logic here
}

func (s *EmailService) SendTransactionalEmail(to string, subject string, body string) error {
	// Simple validation for demonstration purposes
	if to == "" {
		return errors.New("Invalid email address")
	}
	// Send the email (implementation not shown)
	return nil
}

func (s *EmailService) SendPromotionalEmail(to string, subject string, body string) error {
	// Simple validation for demonstration purposes
	if to == "" {
		return errors.New("Invalid email address")
	}
	// Send the email (implementation not shown)
	return nil
}

func TestSendTransactionalEmail(t *testing.T) {
	es := EmailService{}
	err := es.SendTransactionalEmail("test@example.com", "Test Subject", "Test Body")
	if err != nil {
		t.Errorf("Error sending transactional email: %v", err)
	}
}

func TestSendPromotionalEmail(t *testing.T) {
	es := EmailService{}
	err := es.SendPromotionalEmail("test@example.com", "Test Subject", "Test Body")
	if err != nil {
		t.Errorf("Error sending promotional email: %v", err)
	}
}

func TestHandleInvalidEmail(t *testing.T) {
	es := EmailService{}
	err := es.SendTransactionalEmail("", "Test Subject", "Test Body")
	if err == nil {
		t.Error("Expected error for invalid email address")
	}

	err = es.SendPromotionalEmail("", "Test Subject", "Test Body")
	if err == nil {
		t.Error("Expected error for invalid email address")
	}
}
