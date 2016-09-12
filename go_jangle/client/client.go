package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	for {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	fmt.Fprintf(conn, "%s", text)
	}
}
