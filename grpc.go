package main

import (
	"context"

	"github.com/Jimbo8702/email_service/types"
)

type GRPCEmailServer struct {
	types.UnimplementedEmailSerivceServer
	svc EmailService
}

func NewGRPCEmailServer(svc EmailService) *GRPCEmailServer {
	return &GRPCEmailServer{
		svc: svc,
	}
}

func (s *GRPCEmailServer) SendEmail(ctx context.Context, req *types.SendEmailRequest) (*types.None, error) {
	email := &types.Email{
		UserID: req.UserID,
		FullName: req.Fullname,
		Username: req.Username,
		Type: req.Emailtype,
		Data: types.EmailReservationData{
			ReservationID: req.ReservationID,
			ProductID: req.ProductID,
			ProductName: req.Productname,
			StartDate: req.Startdate,
			EndDate: req.Enddate,
			MediaURL: req.Mediaurl,
		},
	}
	return &types.None{}, s.svc.SendEmail(ctx, email)
}