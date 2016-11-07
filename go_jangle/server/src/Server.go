package main

import (
	"fmt"
	"flag"
	"net"
	"container/list"
	"database/sql"
	"os"
)

type Jangle struct {
	user map[uint]*User
	userlist *list.List
	db *sql.DB
	log_file *os.File
	address string
	debug bool
	no_database bool
	logging bool
	logging_warn bool
	Messages []func(*User, []byte) Message
	Commands []func([]string)
}

func Get_User_From_Userid(id uint) *User{
	for e := jangle.userlist.Front(); e != nil; e = e.Next() { 
		if (e.Value.(*User).id == id) {
			return e.Value.(*User);
		}
	}		
	return nil;
}

func Remove_User_From_Userlist(id uint){
	for e := jangle.userlist.Front(); e != nil; e = e.Next() { 
		if (e.Value.(*User).id == id) {
			jangle.userlist.Remove(e);
		}
	}		
}

type Server struct {
	rooms []Room
	serverid uint
	name []byte
	members map[uint]*User
}

type Room struct {
	name []byte
	roomid uint
}

var jangle Jangle;

func main() {
	Init_Flags();
	Init_Logger();
	Init_Server();
	Init_Parse();
	Init_Command();
	


	Load_Server();
	
	Color_Println("red", "JANGLE GO SERVER");
	fmt.Println("listening on - " + jangle.address);
	Logln("Hosting server on address:", jangle.address);
	
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
			fmt.Println("From Address:", user.Get_Remote_Address());
			user.roomid = 1;
			user.serverid = 1;
		}
		
		Logln("User Connected from address:", user.Get_Remote_Address());
		//Add new connection onto the end of connections list
		elem := jangle.userlist.PushBack(user);
		//Recieve data packets from clients
		go Listen_To_Clients(user, elem);
	}
}

//Initializes the list of users and makes connection to the database
func Init_Server(){
	Log("Initializing Server.");
	//Create new list to store every client connection
	jangle.userlist = list.New();
	//Make connection to Database
	jangle.db, _ = Connect_Database();
	defer jangle.log_file.Close()
}

//Creates command line flags and finds their values
func Init_Flags(){
	Log("Initializing Flags.");
	//String Flags
	//Creates address flag (defaults to 'localhost')
	address_flag := flag.String("address", "localhost", "Address of Server");
	//Creates port flag (defaults to '9090')
	port_flag := flag.String("port", "9090", "Port of Server");

	//Boolean Flags
	debug_flag := flag.Bool("debug", false, "Puts server in debug mode");
	no_database_flag := flag.Bool("nodb", true, "Turns off connection to database");
	logging := flag.Bool("log", false, "Turns on outputing to logging file");
	logging_warn := flag.Bool("logwarn", false, "Turns on outputing only warnings to logging file");

	flag.Parse();

	jangle.address = *address_flag + ":" + *port_flag;
	jangle.debug = *debug_flag;
	jangle.no_database = !*no_database_flag;
	jangle.logging_warn = *logging_warn;
	jangle.logging = *logging;
	if(jangle.logging_warn){
		jangle.logging = false;
	}
}

func Load_Server(){
	Log("Loading Server.");
}
