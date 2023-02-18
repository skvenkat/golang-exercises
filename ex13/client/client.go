package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	srvAddr := ":9040"
	tcpAddr, err := net.ResolveTCPAddr("tcp:", srvAddr)
	if err != nil {
		log.Fatalf("couldn't resolve the server with address %s\n%s", srvAddr, err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("couldn't connect to server listening in the address : %s\n", err.Error())
	}

	valueSet := []string{"set chrome google", "get chrome", "delete chrome"}

	for _, value := range valueSet {
		_, err = conn.Write([]byte(value))
		if err != nil {
			log.Printf("failed to write the value in connection buffer : %s\n%s", value, err.Error())
		}

		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		if err != nil {
			log.Printf("failed to read the reply received from server\n %s", err.Error())
		}

		fmt.Println("Received from Server : ", reply)
	}

}
