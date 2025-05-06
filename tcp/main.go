package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn){
	defer conn.Close()
	reader := bufio.NewReader(conn)
	message,err := reader.ReadString('\n')
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Received message:", message)
	conn.Write([]byte("Message received\n"))
}
func main(){
	ln,err := net.Listen("tcp",":8080")
	if err != nil{
		log.Fatal(err)
	}
	defer ln.Close()
	for {
		conn,err := ln.Accept()
		if err != nil{
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
	
}