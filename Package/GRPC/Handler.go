package GRPCHandler

import (
	//Structstore "DistributedFileDBMaster/Helper/StructStore"

	"DistributedFileDBMaster/Helper"
	"context"
	"log"
	"slices"
	"strconv"
	"strings"
)

type ServerStruct struct {
}

func (serv *ServerStruct) mustEmbedUnimplementedMessageExchangeServiceServer() {}

func (serv *ServerStruct) MessageProcessor(ctx context.Context, RecievedMessage *RecievedMessage) (*RecievedMessage, error) {
	log.Printf("Recieved Message from Client")
	Message := RecievedMessage

	ProcessChunkData(RecievedMessage)
	return Message, nil
}

func ProcessChunkData(RecievedMessage *RecievedMessage) {

	num, err := strconv.ParseInt(RecievedMessage.AvailableSize, 10, 64)
	if err != nil {
		log.Println(err)
		
	}

	Helper.AvailableChunk[RecievedMessage.ServerID] = num

	arr := strings.FieldsFunc(RecievedMessage.ChunkList, func(run rune) bool {
		return run == ','
	})

	for _, ChunkName := range arr {

		_, IsPresent := Helper.ChunkMapping[ChunkName]

		if IsPresent == true {
			ServerList, _ := Helper.ChunkMapping[ChunkName]
			if slices.Contains(ServerList, RecievedMessage.ServerID) != true {
				ServerList = append(ServerList, RecievedMessage.ServerID)
				Helper.ChunkMapping[ChunkName] = ServerList
			}
		} else {
			var ServerList []string
			ServerList = append(ServerList, RecievedMessage.ServerID)
			Helper.ChunkMapping[ChunkName] = ServerList
		}

	}
}
