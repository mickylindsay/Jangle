package main

import "fmt"
import "bufio"
import "net"
import "os"
import "log"

func main() {
	address := "000.0.0.0:0000"
	data := make([]byte, 1024)
	conn, err := net.Dial("tcp", address)

	if err != nil {
		log.Fatal(err)
	}

	go fun() {
		for {
			len, _ := conn.Read(data)
			str := string(data[:len])
			parse(str)
		}
	}()
	
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

	}
}