package main

import (
	"context"
	"time"

	"github.com/Jimbo8702/email_service/types"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next EmailService
}

func NewLogMiddleware(svc EmailService) EmailService {
	return &LogMiddleware{
		next: svc,
	}
}

func (l *LogMiddleware) SendEmail(ctx context.Context, email *types.Email) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err": err,
			"user_id": email.UserID,
			"email_type": email.Type,
		}).Info("sending email")
	}(time.Now())
	return l.next.SendEmail(ctx, email)
}

