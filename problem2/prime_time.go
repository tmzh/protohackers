package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net"
	"os"
)

func isPrime(f float64) bool {
	n := int(f)
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

func isMalformed(message *Message) bool {
	if message.Method == nil || *message.Method != "isPrime" || message.Number == nil {
		b, _ := json.Marshal(message)
		log.Printf("Message is: %s\n", string(b))
		log.Printf("Returning response: %s\n", "malformed")
		return true
	}
	return false
}

type Message struct {
	Method *string  `json:"method"`
	Number *float64 `json:"number"`
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		var message *Message

		err := json.Unmarshal(scanner.Bytes(), &message)
		if err != nil || isMalformed(message) {
			log.Printf("Received line: %s\n", scanner.Text())
			log.Printf("Error in JSON: %s\n", err)
			log.Printf("Returning response: %s\n", "malformed")
			conn.Write([]byte("malformed"))
			conn.Close()
		} else {
			response := Response{
				Method: "isPrime",
				Prime:  isPrime(*message.Number),
			}

			b, _ := json.Marshal(response)
			log.Printf("Returning response: %s\n", string(b))
			conn.Write(b)
			conn.Write([]byte{10})

		}
	}

}

func main() {
	ln, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("Cannot create listener, err:", err)
		os.Exit(1)
	}
	log.Printf("Listening on %s\n", ln.Addr())

	for {
		conn, err := ln.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Printf("Connection closed\n")
				return
			}
			log.Printf("Failed to accept incoming connection\n")
			continue
		}

		go handleConnection(conn)

	}

}
