package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const port = ":42069"

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatalf("Failed to connect to %s: %s", conn.RemoteAddr(), err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Invalid input: %s", err)
			continue
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error writing to %s: %s", conn.RemoteAddr(), err)
			continue
		}
	}
}
