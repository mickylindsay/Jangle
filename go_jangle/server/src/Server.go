package main

import (
	"fmt"
	"flag"
	"net"
	"container/list"
	"database/sql"
)

type Jangle struct {
	userlist *list.List
	db *sql.DB
	address string
	debug boolean
}

var jangle Jangle;

func main() {
	Init_Server();
	Init_Flags();
	
	fmt.Println("\x1b[0;31mJANGLE GO SERVER");
	fmt.Println("\x1b[0;0mlistening on - " + jangle.address);
	
	listener, e := net.Listen("tcp", jangle.address);
	Check_Error(e);
	defer listener.Close();
	//Read server console input and write that input to every user
	//go write_stdio_to_clients(jangle.userlist);

	//Listen for new client connection
	for {
		conn, err := listener.Accept();
		defer conn.Close();
		fmt.Println("User Connected");
		user := &User{
			c : &conn,
		};
		//Add new connection onto the end of connections list
		elem := jangle.userlist.PushBack(user);
		Check_Error(err);
		//Recieve data packets from clients
		go Listen_To_Clients(user, elem);
	}
}

//Initializes the list of users and makes connection to the database
func Init_Server(){
	//Create new list to store every client connection
	jangle.userlist = list.New();

	//Make connection to Database
	fmt.Println("Connecting to MySQL Database.")
	var e error;
	jangle.db, e = Connect_Database();
	Check_Error(e);
	fmt.Println("Database Connection Successful.")

}

//Creates command line flags and finds their values
func Init_Flags(){
	//Creates address flag (defaults to 'localhost')
	address_flag := flag.String("address", "localhost", "Address of Server");
	//Creates port flag (defaults to '9090')
	port_flag := flag.String("port", "9090", "Port of Server");

	debug_flag := flag.Bool("debug", false, "Puts server in debug mode");

	flag.Parse();

	jangle.address = *address_flag + ":" + *port_flag;
	jangle.debug = *debug_flag;
}
