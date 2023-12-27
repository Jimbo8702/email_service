package main

import (
	"context"
	"errors"

	"firebase.google.com/go/v4/auth"
	"github.com/Jimbo8702/email_service/types"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailService interface {
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
	var (
		c    = GetAppConfig()
		m 	 = mail.NewV3Mail()
		e    = mail.NewEmail(c.SenderName, c.SenderEmail)
		to   = mail.NewEmail(email.FullName, toEmail)
	)
	m.SetFrom(e)	
	
	switch email.Type {
	case types.WELCOME_EMAIL:
		m.SetTemplateID(c.Templates.WELCOME_EMAIL)
		m = s.makeWelcomeEmail(ctx, email.Username, to, m)

	case types.SUBMISSION_SUCCESS_EMAIL:
		m.SetTemplateID(c.Templates.SUBMISSION_SUCCESS_EMAIL)
		m = s.makeSubmissionSuccessEmail(ctx, to, m, email)

	case types.RESERVATION_ADMIN_EMAIL:
		m.SetTemplateID(c.Templates.RESERVATION_ADMIN_EMAIL)
		m = s.makeAdminReservationEmail(ctx, to, m, email)

	case types.RESERVATION_APPROVED_EMAIL:
		m.SetTemplateID(c.Templates.RESERVATION_ADMIN_EMAIL)
		m = s.makeReservationApprovedEmail(ctx, to, m, email)
		
	default:
		return errors.New("email type not supported")
	}
	
	_, err = s.sgClient.Send(m)
	if err != nil {
		return err
	}

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

func (s *SendGridEmailSerivce) makeWelcomeEmail(ctx context.Context, username string, to *mail.Email, m *mail.SGMailV3) *mail.SGMailV3 {
	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetDynamicTemplateData("username", username)
	m.AddPersonalizations(p)
	return m
}

func (s *SendGridEmailSerivce) makeSubmissionSuccessEmail(ctx context.Context, to *mail.Email, m *mail.SGMailV3, email *types.Email) *mail.SGMailV3 {
	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetDynamicTemplateData("reservation_id", email.Data.ReservationID)
	p.SetDynamicTemplateData("product_name",  email.Data.ProductName)
	p.SetDynamicTemplateData("start_date",  email.Data.StartDate)
	p.SetDynamicTemplateData("end_date",  email.Data.EndDate)
	p.SetDynamicTemplateData("product_media_url", email.Data.MediaURL)
	m.AddPersonalizations(p)
	return m
}

func (s *SendGridEmailSerivce) makeAdminReservationEmail(ctx context.Context, to *mail.Email, m *mail.SGMailV3, email *types.Email) *mail.SGMailV3 {
	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetDynamicTemplateData("user_id", email.UserID)
	p.SetDynamicTemplateData("user_fullname", email.FullName)
	p.SetDynamicTemplateData("user_username", email.Username)
	p.SetDynamicTemplateData("reservation_id", email.Data.ReservationID)
	p.SetDynamicTemplateData("product_id",  email.Data.ProductID)
	p.SetDynamicTemplateData("product_name",  email.Data.ProductName)
	p.SetDynamicTemplateData("start_date",  email.Data.StartDate)
	p.SetDynamicTemplateData("end_date",  email.Data.EndDate)
	p.SetDynamicTemplateData("product_media_url", email.Data.MediaURL)
	return m
}

func (s *SendGridEmailSerivce) makeReservationApprovedEmail(ctx context.Context, to *mail.Email, m *mail.SGMailV3, email *types.Email) *mail.SGMailV3 {
	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetDynamicTemplateData("product_name",  email.Data.ProductName)
	p.SetDynamicTemplateData("start_date",  email.Data.StartDate)
	p.SetDynamicTemplateData("end_date",  email.Data.EndDate)
	p.SetDynamicTemplateData("product_media_url", email.Data.MediaURL)
	return m
}