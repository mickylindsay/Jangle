package main

import (
	"container/list"
	"fmt"
)

type Communicator interface {
	Read(read_data []byte) (int, error)
	Write(write_data []byte) (int, error)
}

//Listens for data in from clients
func Listen_To_Clients(user *User, e *list.Element) {
	for {
		//Array to store the size metadata of an incoming packet
		packet_size := make([]byte, 4)
		//Read data from client
		len, err := (*user).Read(packet_size)

		//If server fails to read from client,
		//the user has disconnected and can be
		//removed from the lsit fo connections
		if err != nil {
			jangle.userlist.Remove(e)
			Color_Println("orange", "User Disconnected")

			Logln("User Disconncete from address:", user.Get_Remote_Address())
			break
		}

		if jangle.debug {
			//fmt.Println("Size: ", packet_size[:], "\nConverted: ", Byte_Converter(packet_size[:]))
		}
		if len < 4 {
			continue
		}
		message_len := Byte_Converter(packet_size[:])
		read_data := make([]byte, message_len)
		read_len, err := (*user).Read(read_data)
		if uint(read_len) < message_len {
			continue
		}

		if jangle.debug {
			fmt.Println("In: ", read_data[:])
		}
		//Send read array to Message file for parsing and processing
		Parse_Data(user, read_data[:])
	}
}

//Sends message to one specific user.
func Send_Message(user *User, message Message) uint {
	write_data := message.Build_Message()
	if jangle.debug {
		fmt.Println("OUT: ", write_data)
	}
	user.Write(write_data)
	return 0
}

//Broadcasts a message to all users.
func Send_Broadcast(message Message) {
	write_data := message.Build_Message()
	if jangle.debug {
		fmt.Println("OUT: ", write_data)
	}
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		e.Value.(*User).Write(write_data)
	}
}

//Broadcasts a message to all users in specific server.
func Send_Broadcast_Server(serverid uint, message Message) {
	write_data := message.Build_Message()
	if jangle.debug {
		fmt.Println(serverid, ": OUT: ", write_data)
	}
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		if e.Value.(*User).serverid == serverid {
			e.Value.(*User).Write(write_data)
		}
	}
}

//Broadcasts a message to all users in specific server and room.
func Send_Broadcast_Server_Room(serverid uint, roomid uint, message Message) {
	write_data := message.Build_Message()
	if jangle.debug {
		fmt.Println(serverid, "-", roomid, ": OUT: ", write_data)
	}
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		if e.Value.(*User).serverid == serverid && e.Value.(*User).roomid == roomid {
			e.Value.(*User).Write(write_data)
		}
	}
}

//TODO
func Send_Broadcast_Members(serverid uint, message Message) {

}

//TODO
func Send_Broadcast_Friends(userid uint, message Message) {

}