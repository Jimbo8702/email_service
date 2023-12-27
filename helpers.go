package main

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)


func NewFirebaseAuthClient(fbAccountID string) (*auth.Client, error) {
	conf := &firebase.Config{
		ServiceAccountID: fbAccountID,
	}
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		return nil, err
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return client, err
}

// func GetTemplateIDForEmail(t string) (string, error) {
// 	ids := GetTemplateIDs()

// 	switch t {
// 	case types.WELCOME_EMAIL:
// 		return ids.WELCOME_EMAIL, nil

// 	case types.SUBMIT_RESERVATION_SUCCESS_EMAIL:
// 		return ids.SUBMIT_RESERVATION_SUCCESS_EMAIL, nil

// 	case types.RESERVATION_APPROVED_EMAIL: 
// 	   return ids.RESERVATION_APPROVED_EMAIL, nil

// 	case types.RESERVATION_ADMIN_EMAIL:
// 		return ids.RESERVATION_ADMIN_EMAIL, nil
		
// 	default: 
// 		return "", errors.New("email type not suporrted")
// 	}
// }