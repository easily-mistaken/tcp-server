package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println("Processing request...")
	time.Sleep(1 * time.Second)
	fmt.Println(string(buf))

	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\r\n"))

}

func main(){
	listner, err := net.Listen("tcp", ":1729")

	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("Waiting for connection...")
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal(err)
		}	
		fmt.Println(conn)
		fmt.Println("Connection accepted")

		go handleConnection(conn)
	}

	

}