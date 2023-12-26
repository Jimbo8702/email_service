package types

import (
	"errors"
	"os"
)

var (
	WELCOME_EMAIL                    = "welcome"
	SUBMIT_RESERVATION_SUCCESS_EMAIL = "user_reservation"
	RESERVATION_APPROVED_EMAIL       = "reservation_approved"
	RESERVATION_ADMIN_EMAIL          = "admin_reservation"
)

type Email struct {
	// use this to get the email from firebase
	UserID   string
	FullName string
	Username string
	Type     string
}

func (e *Email) GetEmailTemplateID() (string, error) {
	switch e.Type {
	case WELCOME_EMAIL:
		return os.Getenv("WELCOME_EMAIL_TEMPLATE_ID"), nil
	case SUBMIT_RESERVATION_SUCCESS_EMAIL:
		return os.Getenv("SUBMIT_RESERVATION_SUCCESS_EMAIL_TEMPLATE_ID"), nil
	case RESERVATION_APPROVED_EMAIL: 
	   return os.Getenv("RESERVATION_APPROVED_EMAIL_TEMPLATE_ID"), nil
	case RESERVATION_ADMIN_EMAIL:
		return os.Getenv("RESERVATION_ADMIN_EMAIL_TEMPLATE_ID"), nil
	default: 
		return "", errors.New("email type not suporrted")
	}
}
