package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
	"container/list"
)

func main() {


	//users_map := make(map[int]*User)


	//Create new list to store every client connection
	users := list.New()
	//Address to host server on
	address := "localhost:9090"

	fmt.Println("JANGLE GO SERVER")
	fmt.Println("address - " + address)
	
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	//Read server console input and write that input to every user
	go write_stdio_to_clients(users)
	//Listen for new client connection
	for {
		conn, err := listener.Accept()
		defer conn.Close()
		user := &User{
			c : &conn,
		}
		//Add new connection onto the end of connections list
		elem := users.PushBack(user)
		if err != nil {
			log.Fatal(err)
		}
		//Read from client and write data to every client
		go listen_to_clients(users, user, elem)
	}
}

func listen_to_clients(users *list.List, user *User, e *list.Element){
	//Array to store data read from client
	read_data := make([]byte, 1024)

	for {
		//Read data from client
		read_len, err := (*user).Read(read_data)
		//If server fails to read from client,
		//the user has disconnected and can be
		//removed from the lsit fo connections
		if err != nil {
			users.Remove(e)
			fmt.Println("User Disconnected")
			break
		}
		//Cast read data into a string
		read_string := string(read_data[:read_len])
		fmt.Println("\t",read_string)
		//Write read_string to entire list fo connections
		write_to_clients(users, read_string)
	}
}

func write_stdio_to_clients(connections *list.List){
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		write_to_clients(connections,text)
	}	
}

//Writes a string to every connection in the list of client connections
func write_to_clients(connections *list.List, s string){
	//Iterate over every client
	for e := connections.Front(); e != nil; e = e.Next() {
		//Write data to every connection
		e.Value.(*User).Printf("%s", s)

		//fmt.Fprintf(*(e.Value.(*User).c), "%s", s)
	}
}
