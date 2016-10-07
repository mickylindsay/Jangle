package main

import (
	"fmt"
	"net"
	"os"	
	"io/ioutil"
	"path/filepath"
	"container/list"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Jangle struct {
	userlist *list.List
	db *sql.DB
}

var jangle Jangle;

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
	var address string
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) 
	dat, err := ioutil.ReadFile(dir + "/../.address")
	//If such file does not exist prompt the user to enter a DSN
	if err != nil{
		address = "localhost:9090"
	}else{
		address = string(dat)
	}

	fmt.Println("JANGLE GO SERVER");
	fmt.Println("listening on - " + address);
	
	listener, e := net.Listen("tcp", address);
	Check_Error(e);
	defer listener.Close();
	//Read server console input and write that input to every user
	//go write_stdio_to_clients(jangle.userlist);
	//Listen for new client connection
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
	read_data := make([]byte, 1024);

	for {
		//Read data from client
		len, err := (*user).Read(read_data);
		//If server fails to read from client,
		//the user has disconnected and can be
		//removed from the lsit fo connections
		if err != nil {
			jangle.userlist.Remove(e);
			fmt.Println("User Disconnected");
			break;
		}
		//Send read array to Message file for parsing and processing
		Parse_data(user, read_data[:len]);
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
