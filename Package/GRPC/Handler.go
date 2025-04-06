package GRPCHandler

import (
	"context"
	"log"
)

type ServerStruct struct {
}

func (serv *ServerStruct) mustEmbedUnimplementedMessageExchangeServiceServer() {}

func (serv *ServerStruct) MessageProcessor(ctx context.Context, RecievedMessage *RecievedMessage) (*RecievedMessage, error) {
	log.Printf("Recieved Message from Client")
	Message := RecievedMessage
	return Message, nil
}
