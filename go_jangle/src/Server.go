package main

import (
	"fmt"
	"log"
	"net"
	//"bufio"
	//"os"
	"container/list"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Jangle struct {
	userlist *list.List
	db *sql.DB
}

var jangle Jangle;

func Check_Error(e error){
	if(e != nil){
		log.Fatal(e);
	}
}

func main() {
	//Create new list to store every client connection
	jangle.userlist = list.New();

	//Make connection to Database
	fmt.Println("Connecting to MySQL Database.")
	var e error;
	jangle.db, e = Connect_Database();
	Check_Error(e);
	fmt.Println("Database Connection Successful.")
	
	//Address to host server on
	address := "localhost:9090";

	fmt.Println("JANGLE GO SERVER");
	fmt.Println("listening on - " + address);
	
	listener, e := net.Listen("tcp", address);
	Check_Error(e);
	defer listener.Close();
	//Read server console input and write that input to every user
	//go write_stdio_to_clients(jangle.userlist);
	//Listen for new client connection
	go accept_connections(listener);
}

func accept_connections(listener net.Listener){
	for {
		conn, err := listener.Accept();
		defer conn.Close();
		fmt.Println("User Connected... Waiting for login message.");
		user := &User{
			c : &conn,
		};
		//Add new connection onto the end of connections list
		elem := jangle.userlist.PushBack(user);
		Check_Error(err);
		//Read from client and write data to every client
		go listen_to_clients(user, elem);
	}
}

func listen_to_clients(user *User, e *list.Element){
	//Array to store data read from client
	//read_data := make([]byte, 1024);
	var read_data []byte;
	for {
		//Read data from client
		/*read_len*/ _, err := (*user).Read(read_data);
		//If server fails to read from client,
		//the user has disconnected and can be
		//removed from the lsit fo connections
		if err != nil {
			jangle.userlist.Remove(e);
			fmt.Println("User Disconnected");
			break;
		}
		Parse_data(read_data);
		//Cast read data into a string
		//read_string := string(read_data[:read_len]);
		//fmt.Println("\t",read_string);
		//Write read_string to entire list fo connections
		//write_to_clients(users, read_string);
	}
}
/*
func write_stdio_to_clients(connections *list.List){
	for {
		reader := bufio.NewReader(os.Stdin);
		text, _ := reader.ReadString('\n');
		write_to_clients(connections,text);
	}	
}

//Writes a string to every connection in the list of client connections
func write_to_clients(connections *list.List, s string){
	//Iterate over every client
	for e := connections.Front(); e != nil; e = e.Next() {
		//Write data to every connection
		e.Value.(*User).Printf("%s", s)
	}
}
*/
