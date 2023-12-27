package emailservice

import (
	"github.com/Jimbo8702/email_service/client"
	"github.com/Jimbo8702/email_service/types"
)

// request to send an email to the grpc email service
type SendEmailRequest types.SendEmailRequest

// client interface for the email service
type Client client.Client

// grpc client for the email service
type GRPCClient client.GRPCClient

// makes a new grpc email service
func NewGRPCClient(endpoint string) (*client.GRPCClient, error) {
	return client.NewGRPCClient(endpoint)
}

