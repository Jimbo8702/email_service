package main

import (
	"context"
	"log"

	"github.com/Jimbo8702/email_service/types"
	"github.com/sendgrid/sendgrid-go"
)

func main() {
	fb, err := NewFirebaseAuthClient(GetAppConfig().FIREBASE_ACCOUNT_ID)
	if err != nil {
		log.Fatal(err)
	}
	sg := sendgrid.NewSendClient(GetAppConfig().SendGridApiKey)

	svc := NewSendGridEmailSerivce(sg, fb)

	testEmail := &types.Email{
		UserID: "O1nm8wcHJdV5x24MkR20GMXVNmO2",
		Username: "MyCoolUsername",
		FullName: "James Sgarella",
		Type: types.WELCOME_EMAIL,
	}

	if err := svc.SendEmail(context.Background(), testEmail); err != nil {
		log.Fatal(err)
	}
}

