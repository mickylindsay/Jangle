package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	address := "localhost:9090"
	read_data := make([]byte, 1024)

	fmt.Println("JANGLE GO SERVER")
	fmt.Println("address - " + address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		fmt.Println("User Connected")
		defer conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			for {
				read_len, err := c.Read(read_data)
				if err != nil {
					fmt.Println("User Disconnected")
					break
				}
				read_string := string(read_data[:read_len])
				fmt.Printf("\tRead %d bytes\n", read_len)
				fmt.Println("\t",read_string)
			}
		}(conn)

	}
}
