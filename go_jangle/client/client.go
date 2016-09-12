package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
)

func main(){
	read_data := make([]byte, 1024)
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	go func(){
		for {
			read_len, _ := conn.Read(read_data)
			fmt.Printf("%s", string(read_data[:read_len]))
		}
	}()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		send_message(conn,text)
	}
}

func send_message(conn net.Conn, text string){
	fmt.Fprintf(conn, "%s", text)
}
