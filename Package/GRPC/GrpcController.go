package GRPCHandler

import (
	
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func InitiateGRPCConnection() {

	ListenConnectionInstance, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println(err)
	}

	GRPCServer := grpc.NewServer()

	server :=  ServerStruct{}

	RegisterMessageExchangeServiceServer(GRPCServer, &server)

	err = GRPCServer.Serve(ListenConnectionInstance)

	if err != nil {
		fmt.Println(err)
	}

}
