package main

import(
	"testing"
	"net"
)

var user User;

func Init_Testing_Server(){
	Init_Server();
	listener, e := net.Listen("tcp", "localhost:9090");
	Check_Error(e);
	defer listener.Close();
	conn, _ := listener.Accept();
	defer conn.Close();
	Color_Println("green", "User Connected");
	user = User{
		c : &conn,
	};
	if(jangle.debug){
		user.roomid = 1;
		user.serverid = 1;
	}
	//Add new connection onto the end of connections list
	elem := jangle.userlist.PushBack(&user);
	//Recieve data packets from clients
	go Listen_To_Clients(&user, elem);
}

func TestInit(t *testing.T){
	Init_Testing_Server();
}