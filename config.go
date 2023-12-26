package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// type EmailConfig struct {
// 	SenderName 		 string
// 	SenderEmail 	 string
// 	ApiKey           string
// }

type ApplicationConfig struct {
	SenderName 		 	string
	SenderEmail 	 	string
	SendGridApiKey      string
	GRPC_LISTEN_ADDR 	string
	FIREBASE_ACCOUNT_ID string
}

type TemplateIDs struct {
	WELCOME_EMAIL 					 string
    SUBMIT_RESERVATION_SUCCESS_EMAIL string
    RESERVATION_APPROVED_EMAIL 		 string
    RESERVATION_ADMIN_EMAIL    		 string
}

func GetAppConfig() *ApplicationConfig {
	return &ApplicationConfig{
		SenderName: os.Getenv("SENDGRID_SENDER_NAME"),
		SenderEmail: os.Getenv("SENDGRID_SENDER_EMAIL"),
		SendGridApiKey: os.Getenv("SENDGRID_API_KEY"),
		GRPC_LISTEN_ADDR: os.Getenv("GRPC_LISTEN_ADDR"),
		FIREBASE_ACCOUNT_ID: os.Getenv("FIREBASE_SERVICE_ACCOUNT_ID"),
	}
}

func GetTemplateIDs() *TemplateIDs {
	return &TemplateIDs{
		WELCOME_EMAIL:  os.Getenv("WELCOME_EMAIL_TEMPLATE_ID"),
		SUBMIT_RESERVATION_SUCCESS_EMAIL: os.Getenv("SUBMIT_RESERVATION_SUCCESS_EMAIL_TEMPLATE_ID"),
		RESERVATION_APPROVED_EMAIL: os.Getenv("RESERVATION_APPROVED_EMAIL_TEMPLATE_ID"),
		RESERVATION_ADMIN_EMAIL: os.Getenv("RESERVATION_ADMIN_EMAIL_TEMPLATE_ID"),
	}
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
}