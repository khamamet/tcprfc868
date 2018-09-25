package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Not enough parameters. \nusage:    time-client timeserver port\nExample:  time-client time.nist.gov 37")
		os.Exit(1)
	}
	serv := os.Args[1]
	port := os.Args[2]
	conn, err := net.Dial("tcp", serv+":"+port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b := make([]byte, 4)
	n, err := conn.Read(b)
	if err != nil || n != 4 {
		fmt.Println("The response is BAD:", err)
	} else {
		fmt.Println(binary.BigEndian.Uint32(b) - 2208988800)
	}
}
