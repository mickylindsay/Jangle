package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
)
func main() {
	connections := make([]net.Conn, 6)
	num_connections := 0
	address := "localhost:9090"
	read_data := make([]byte, 1024)

	fmt.Println("JANGLE GO SERVER")
	fmt.Println("address - " + address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	go func(){
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			write_to_clients(connections,text)
		}	
	}()
	for {
		connections[num_connections], err = listener.Accept()
		fmt.Println("User Connected")
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			num_connections++
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
			c.Close()
		}(connections[num_connections])
	}
}
func write_to_clients(connections []net.Conn, s string){
	for _,c := range connections {
		if(c != nil){
			fmt.Fprintf(c, "SERVER: %s\n", s)
		}
	}
	
}
