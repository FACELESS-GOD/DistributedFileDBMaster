package Utility

import (
	"DistributedFileDBMaster/Helper"
	Structstore "DistributedFileDBMaster/Helper/StructStore"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"slices"

	"github.com/jasonlvhit/gocron"
)

type TcpServer struct {
}

var SocketConnection net.Listener

var Connections []net.Conn

func InitiateSocketConnection() {

	conn, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}
	SocketConnection = conn
}

func AcceptingIncommingRequest() {
	for {
		conn, err := SocketConnection.Accept()
		if err != nil {
			fmt.Println(err)
		}
		Connections = append(Connections, conn)
		go MessageController(conn)
	}
}

func MessageController(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the incoming data
	fmt.Printf("Received: %s", buf)

	ChunkData := Structstore.ChunkMapping{}
	buf = bytes.Trim(buf, "\x00")
	err = json.Unmarshal(buf, &ChunkData)
	if err != nil {
		fmt.Println(err.Error())
	}
	r := buf[0]
	fmt.Print(r)
	ProcessChunkData(ChunkData)

}
func ProcessChunkData(ChunkData Structstore.ChunkMapping) {

	Helper.AvailableChunk[ChunkData.ServerID] = ChunkData.AvailableSpace

	for _, ChunkName := range ChunkData.ChunkList {

		_, IsPresent := Helper.ChunkMapping[ChunkName]

		if IsPresent == true {
			ServerList, _ := Helper.ChunkMapping[ChunkName]
			if slices.Contains(ServerList, ChunkData.ServerID) != true {
				ServerList = append(ServerList, ChunkData.ServerID)
				Helper.ChunkMapping[ChunkName] = ServerList
			}
		} else {
			var ServerList []string
			ServerList = append(ServerList, ChunkData.ServerID)
			Helper.ChunkMapping[ChunkName] = ServerList
		}

	}

}

func InitiateBroadCast() {
	gocron.Every(1).Minute().Do(BroadCastMessage)
	<-gocron.Start()
}

func BroadCastMessage() {
	for _, client := range Connections {
		var message = []byte("HeartBeat")
		message = bytes.Trim(message, "\x00")
		writer := bufio.NewWriter(client)
		_, err := writer.Write(message)

		if err != nil {
			fmt.Println(err)
		}

		err = writer.Flush()
		if err != nil {
			fmt.Println(err)
		}
		//client.Write(message)
	}
}

func TerminateSocketConnection() {
	SocketConnection.Close()
}
