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
	address string
}

var jangle Jangle;

func main() {
	Init_Server();
	
	fmt.Println("JANGLE GO SERVER");
	fmt.Println("listening on - " + jangle.address);
	
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
		//Read from client and write data to every client
		go Listen_To_Clients(user, elem);
	}
}

func Init_Server(){
	//Create new list to store every client connection
	jangle.userlist = list.New();

	//Make connection to Database
	fmt.Println("Connecting to MySQL Database.")
	var e error;
	jangle.db, e = Connect_Database();
	Check_Error(e);
	fmt.Println("Database Connection Successful.")

	//Address to host server on	
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) 
	dat, err := ioutil.ReadFile(dir + "/../.address")
	//If such file does not exist prompt the user to enter a DSN
	if err != nil{
		jangle.address = "localhost:9090"
	}else{
		jangle.address = string(dat)
	}
}

