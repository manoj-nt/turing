// email_service_test.go
package emailservice

import "testing"

func TestSendTransactionalEmail(t *testing.T) {
	service := &BasicEmailService{}

	err := service.SendTransactionalEmail("user@example.com", "This is a transactional email.")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSendPromotionalEmail(t *testing.T) {
	service := &BasicEmailService{}

	err := service.SendPromotionalEmail("user@example.com", "This is a promotional email.")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSendEmailWithInvalidAddress(t *testing.T) {
	service := &BasicEmailService{}

	err := service.SendTransactionalEmail("invalid-email", "This email should fail.")
	if err == nil {
		t.Errorf("Expected error for invalid email address, got none")
	}

	err = service.SendPromotionalEmail("invalid-email", "This email should fail.")
	if err == nil {
		t.Errorf("Expected error for invalid email address, got none")
	}
}
