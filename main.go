package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Jimbo8702/email_service/types"
	"github.com/sendgrid/sendgrid-go"
	"google.golang.org/grpc"
)

func main() {
	fb, err := NewFirebaseAuthClient(GetAppConfig().FIREBASE_ACCOUNT_ID)
	if err != nil {
		log.Fatal(err)
	}
	var (
		sg 	   = sendgrid.NewSendClient(GetAppConfig().SendGridApiKey)
		svc    = NewSendGridEmailSerivce(sg, fb)
		svcwl  = NewLogMiddleware(svc)
	)
	log.Fatal(makeGRPCTransport(GetAppConfig().GRPC_LISTEN_ADDR, svcwl))
}

func makeGRPCTransport(listenAddr string, svc EmailService) error {
	fmt.Println("GRPC transport running on port:", listenAddr)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	server := grpc.NewServer([]grpc.ServerOption{}...)

	types.RegisterEmailSerivceServer(server, NewGRPCEmailServer(svc))
	return server.Serve(ln)
}
