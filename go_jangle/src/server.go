package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
	"container/list"
)

type User struct {
	c *net.Conn
}

func main() {
	//Create new list to store every client connection
	connections := list.New()
	//Address to host server on
	address := "localhost:9090"
	//Array to store data read from client
	read_data := make([]byte, 1024)

	fmt.Println("JANGLE GO SERVER")
	fmt.Println("address - " + address)
	
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//Read server console input and write that input to every user
	go func(){
		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			write_to_clients(connections,text)
		}	
	}()
	//Listen for new client connection
	for {
		conn, err := listener.Accept()
		defer conn.Close()
		user := &User{
			c : &conn,
		}
		//Add new connection onto the end of connections list
		elem := connections.PushBack(user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("User Connected")
		//Read from client and write data to every client
		go func(user *User, e *list.Element) {
			for {
				//Read data from client
				read_len, err := (*user.c).Read(read_data)
				//If server fails to read from client,
				//the user has disconnected and can be
				//removed from the lsit fo connections
				if err != nil {
					connections.Remove(e)
					fmt.Println("User Disconnected")
					break
				}
				//Cast read data into a string
				read_string := string(read_data[:read_len])
				fmt.Println("\t",read_string)
				//Write read_string to entire list fo connections
				write_to_clients(connections, read_string)
			}
		}(user, elem)
	}
}
//Writes a string to every connection in the list of client connections
func write_to_clients(connections *list.List, s string){
	//Iterate over every client
	for e := connections.Front(); e != nil; e = e.Next() {
		//Write data to every connection
		fmt.Fprintf(*(e.Value.(*User).c), "%s", s)
	}
}
