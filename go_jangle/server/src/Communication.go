package main

import (
	"fmt"
	"container/list"
)

//Listens for data in from clients
func Listen_To_Clients (user *User, e *list.Element) {
	//Array to store data read from client
	read_data := make([]byte, 1048576);

	for {
		//Read data from client
		len, err := (*user).Read(read_data);
		//If server fails to read from client,
		//the user has disconnected and can be
		//removed from the lsit fo connections
		if err != nil {
			jangle.userlist.Remove(e);
			Color_Println("orange", "User Disconnected");
			break;
		}
		if (jangle.debug) {
			fmt.Println("In: ", read_data[:len]);
		}
		//Send read array to Message file for parsing and processing
		Parse_Data(user, read_data[:len]);
	}
}

//Sends message to one specific user.
func Send_Message (user *User, message Message) uint {
	write_data := message.Build_Message();
	if (jangle.debug) {
		fmt.Println("OUT: ", write_data);
	}
	user.Write(write_data)
	
	return 0;
}

//Broadcasts a message to all users.
func Send_Broadcast (message Message) {
	write_data := message.Build_Message();
	if (jangle.debug) {
		fmt.Println("OUT: ", write_data);
	}
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		e.Value.(*User).Write(write_data);
	}			
}

//Broadcasts a message to all users in specific server.
func Send_Broadcast_Server (serverid uint, message Message) {
	write_data := message.Build_Message();
	if (jangle.debug) {
		fmt.Println(serverid, ": OUT: ", write_data);
	}
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		if (e.Value.(*User).serverid == serverid) {
			e.Value.(*User).Write(write_data);
		}
	}			
}

//Broadcasts a message to all users in specific server and room.
func Send_Broadcast_Server_Room (serverid uint, roomid uint, message Message) {
	write_data := message.Build_Message();
	if (jangle.debug) {
		fmt.Println(serverid, ".", roomid, ": OUT: ", write_data);
	}
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		if (e.Value.(*User).serverid == serverid && e.Value.(*User).roomid == roomid) {
			e.Value.(*User).Write(write_data);
		}
	}			
}