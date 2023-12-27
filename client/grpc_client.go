package client

import (
	"context"

	"github.com/Jimbo8702/email_service/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	Endpoint string
	client types.EmailSerivceClient
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := types.NewEmailSerivceClient(conn)
	return &GRPCClient{
		Endpoint: endpoint,
		client: c,
	}, nil
}

func (c *GRPCClient) SendEmail(ctx context.Context, req *types.SendEmailRequest) error {
	_, err := c.client.SendEmail(ctx, req)
	return err
}