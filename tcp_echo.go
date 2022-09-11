package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("Cannot create listener, err:", err)
		os.Exit(1)
	}
	fmt.Printf("Listening on %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept in coming connection\n")
			continue
		}

		go func(c net.Conn) {
			// Echo all incoming data.
			_, err := io.Copy(c, c)
			if err != nil {
				return
			}
			// Shut down the connection.
			c.Close()
		}(conn)
	}



}
