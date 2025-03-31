package Controller

import (
	"fmt"
	"net"
)

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
