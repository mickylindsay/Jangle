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
	debug bool
	no_database bool
	Messages []func(*User, []byte)
}

var jangle Jangle;

func main() {
	Init_Flags();
	Init_Server();
	
	Color_Println("red", "JANGLE GO SERVER");
	fmt.Println("listening on - " + jangle.address);
	
	listener, e := net.Listen("tcp", jangle.address);
	Check_Error(e);
	defer listener.Close();
	//Read server console input and write that input to every user
	//go write_stdio_to_clients(jangle.userlist);

	//Listen for new client connection
	for {
		conn, _ := listener.Accept();
		defer conn.Close();
		Color_Println("green", "User Connected");
		user := &User{
			c : &conn,
		};
		if(jangle.debug){
			user.roomid = 1;
			user.serverid = 1;
		}
		//Add new connection onto the end of connections list
		elem := jangle.userlist.PushBack(user);
		//Recieve data packets from clients
		go Listen_To_Clients(user, elem);
	}
}

//Initializes the list of users and makes connection to the database
func Init_Server(){
	//Create new list to store every client connection
	jangle.userlist = list.New();
	//Make connection to Database
	jangle.db, _ = Connect_Database();
}

//Creates command line flags and finds their values
func Init_Flags(){
	//String Flags
	//Creates address flag (defaults to 'localhost')
	address_flag := flag.String("address", "localhost", "Address of Server");
	//Creates port flag (defaults to '9090')
	port_flag := flag.String("port", "9090", "Port of Server");

	//Boolean Flags
	debug_flag := flag.Bool("debug", false, "Puts server in debug mode");
	no_database_flag := flag.Bool("nodb", false, "Turns off connection to database");

	flag.Parse();

	jangle.address = *address_flag + ":" + *port_flag;
	jangle.debug = *debug_flag;
	jangle.no_database = *no_database_flag;
}
