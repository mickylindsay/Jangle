package main

import (
	"container/list"
	"fmt"
)

//Interface used for objects that can be Read and Written to
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
			user.status = uint(offline)
			user.serverid = uint(0)
			user.roomid = uint(0)
			m := Create_Message(recieve_status, Int_Converter(user.id), byte(user.status), byte(user.muted), byte(user.voice))
			Send_Broadcast_Server(uint(1), m)
			m = Create_Message(recieve_location, Int_Converter(user.serverid), Int_Converter(user.roomid), Int_Converter(user.id))
			Send_Broadcast_Server(uint(1), m)
			break
		}
		//If data is recieved by the server with a length of less than 4 bytes, ignore them
		if len < 4 {
			continue
		}
		//Use the 4 bytes read to determine the size of the following message packet
		message_len := Byte_Converter(packet_size[:])
		//Create array to store the entire packet
		read_data := make([]byte, message_len)
		//Read into the new array
		read_len, err := (*user).Read(read_data)
		//If the array has not been filled, the following data needs to be appended to the end
		if uint(read_len) < message_len {
			continue
		}
		
		//Entire message packet has been read then print the debug statment if necessary and parse the message.
		if jangle.debug {
			fmt.Println("In: ", read_data[:])
		}
		//Send read array to Message file for parsing and processing
		Parse_Data(user, read_data[:])
	}
}

//Sends message to one specific user.
func Send_Message(user *User, message Message) uint {
	if user == nil{
		return 0;
	}
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
	//Iterate through all users and write the data to each user
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
	//Iterate through all users and write the data to each user who is in the corresponding server.
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
	//Iterate through all users and write the data to each user who is in the corresponding server and room.
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		if e.Value.(*User).serverid == serverid && e.Value.(*User).roomid == roomid {
			e.Value.(*User).Write(write_data)
		}
	}
}

//Broadcasts a message to all users who are members of a specific server.
func Send_Broadcast_Members(serverid uint, message Message) {
	//Create list of members from a specific server
	ids, err :=  Get_Member_Userid(serverid);
	if err == nil{
		fmt.Println("Unable to find members of server", serverid);
		return;
	}
	//Attempt to send the message to each member
	for i := 0; i < len(ids); i++ {
		Send_Message(Get_User_From_Userid(ids[i]), message);
	}
}

//TODO
func Send_Broadcast_Friends(userid uint, message Message) {

}
