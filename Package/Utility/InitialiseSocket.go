package Utility

import (
	"fmt"
	"net"

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

func InitiateBroadCast() {
	gocron.Every(1).Minute().Do(BroadCastMessage)
	<-gocron.Start()
}

func BroadCastMessage() {
	for _, client := range Connections {
		var message = []byte("HeartBeat")
		client.Write(message)
	}
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
}

func TerminateSocketConnection() {
	SocketConnection.Close()
}
