package client

import (
	"context"

	"github.com/Jimbo8702/email_service/types"
)

type Client interface {
	SendEmail(context.Context, *types.SendEmailRequest) error
}