package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"os"
)

func isPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if (n%2 == 0) || (n%3 == 0) {
		return false
	}
	i := 5
	stop := int(math.Sqrt(float64(n)))
	for i <= stop {
		if (n%i == 0) || (n%(i+2) == 0) {
			return false
		}
		i += 6
	}
	return true
}

func parseMessage() {

}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Cannot create listener, err:", err)
		os.Exit(1)
	}
	log.Printf("Listening on %s\n", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Printf("Connection closed\n")
				return
			}

			log.Printf("Failed to accept incoming connection\n")
			continue
		}

		go func(c net.Conn) {
			b, err := io.ReadAll

		}(conn)
	}

}
