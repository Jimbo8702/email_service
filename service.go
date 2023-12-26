package main

import (
	"context"
	"errors"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/Jimbo8702/email_service/types"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService interface {
	checkEmailVerified(ctx context.Context, userID string) (string, error)
	SendEmail(ctx context.Context, email *types.Email) error 
}

type SendGridEmailSerivce struct {
	fbAuth   *auth.Client
	sgClient *sendgrid.Client
}

func NewSendGridEmailSerivce(sgClient *sendgrid.Client, fbAuth *auth.Client) EmailService {
	return &SendGridEmailSerivce{
		fbAuth: fbAuth,
		sgClient: sgClient,
	}
}

func (s *SendGridEmailSerivce) SendEmail(ctx context.Context, email *types.Email) error {
	toEmail, err := s.checkEmailVerified(ctx, email.UserID); 
	if err != nil {
		return err
	}

	tempID, err := GetTemplateIDForEmail(email.Type)
	if err != nil {
		return err
	}
	
	var (
		c    = GetAppConfig()
		m 	 = mail.NewV3Mail()
		e    = mail.NewEmail(c.SenderName, c.SenderEmail)
		to   = mail.NewEmail(email.FullName, toEmail)
		tos  = []*mail.Email{to}
		p    = mail.NewPersonalization()
	)

	m.SetFrom(e)	
  	m.SetTemplateID(tempID)
	p.AddTos(tos...)
	p.SetDynamicTemplateData("username", email.Username)
	m.AddPersonalizations(p)

	response, err := s.sgClient.Send(m)
	if err != nil {
		return err
	}

	fmt.Println("res code: ", response.StatusCode)
    fmt.Printf("body: %s\n", response.Body)
    fmt.Printf("headers: %v\n", response.Headers)

	return nil
}

func (s *SendGridEmailSerivce) checkEmailVerified(ctx context.Context, userID string) (string, error) {
	u, err := s.fbAuth.GetUser(ctx, userID)
	if err != nil {
		return "", err
	}
	if !u.EmailVerified {
		return "", errors.New("email is not verified") 
	}
	return u.Email, nil
}