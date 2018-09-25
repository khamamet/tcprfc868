package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	var p int //port to start the service

	flag.IntVar(&p, "p", 11037, "Interface port to bind and listen")
	flag.Parse()

	//graceful shutdown
	gracefulStop := make(chan os.Signal, 2)
	signal.Notify(gracefulStop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGALRM)
	go func() {
		for {
			sig := <-gracefulStop
			if sig == syscall.SIGHUP {
				log.Println("got signal SIGHUP. Ignored!")
			} else {
				log.Printf("got signal: %+v. Exiting...", sig)
				os.Exit(0)
			}
		}
	}()

	laddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:"+strconv.Itoa(p))
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	fmt.Println("New incoming connection:", conn.RemoteAddr())
	b := make([]byte, 4)
	now := time.Now().Unix() + 2208988800
	binary.BigEndian.PutUint32(b, uint32(now))
	conn.Write(b)
	conn.Close()
}

/*
printf "%d\n" "0x$(nc time.nist.gov 37 | xxd -p)"
printf "%d\n" "0x$(nc 127.0.0.1 3377 | xxd -p)"
*/
