package main

import (
	"fmt"
	"time"
	"log"
)

//Initializes all code cases

	//Login type codes
	var create_user byte = 0
	var create_user_fail byte = 1
	var login byte = 2
	var login_fail byte = 3
	var login_success byte = 4

	//Message type codes
	var message_client_send byte = 16
	var message_client_recieve byte = 17

	//Request from client type codes
	var request_n_messages byte = 32
	var request_all_userid byte = 33
	var request_display_name byte = 34
	var request_all_serverid byte = 35
	var request_server_display_name byte = 36
	var request_all_roomid byte = 37
	var request_room_display_name byte = 38
	var request_master_display_name byte = 39
	var request_status byte = 40

	//Client recieve type codes
	var recieve_userid byte = 48
	var recieve_display_name byte = 49
	var recieve_serverid byte = 50
	var recieve_server_display_name byte = 51
	var recieve_roomid byte = 52
	var recieve_room_display_name byte = 53
	var recieve_master_display_name byte = 54
	var recieve_status byte = 55

	//Client send type codes
	var send_new_display_name byte = 64
	var send_new_server_display_name byte = 65
	var send_new_room_display_name byte = 66
	var send_new_master_display_name byte = 67

	//Status change of client type codes
	var change_status byte = 80
	var change_server byte = 81
	var change_room byte = 82
	
	//Broadcast type codes
	var broadcast_status byte = 96
	var broadcast_server byte = 97
	var broadcast_room byte = 98
	var broadcast_display_name byte = 99
	var broadcast_server_display_name byte = 100
	var broadcast_room_display_name byte  = 101
	var broadcast_master_display_name byte = 102

	//Error type codes
	var error_check byte = 255

	//Status types
	var online byte = 1
	var away byte = 2
	var offline byte = 3

//Converts byte array to unsigned int 
func Byte_Converter(data []byte) uint {
	var i uint
	var sum uint

	for i = 0; int(i) < len(data); i++ {
		//Preforms little endian bit shifting and adds int value to sum for each byte
		sum += uint(data[i]) << (8 * i)
	}

	return sum
}

//Converts unsigned int to byte array
func Int_Converter(num uint) []byte {
	data := make([]byte, 4)

	for i := 0; i < 4; i++ {
		mod := num % 256
		data[i] = byte(mod)
		num /= 256
	}

	return data
}

//Returns Current Millisecond time as unsigned int
func Milli_Time() uint {
	return uint(time.Now().UnixNano() / 1000000000)
}

//Used for time stamping code type 16 messages
//Takes in a byte array and creates a 4 byte space from byte 13 to 16
//Places 4 byte time stamp in space
func Time_Stamp (data []byte) []byte {
	new_data := make([]byte, len(data) + 4)
	copy(new_data[0:12], data[0:12])

	for i := 13; i < len(data); i++ {
		new_data[i + 4] = data[i]
	}

	copy(new_data[13:16], Int_Converter(Milli_Time()))
	return new_data
}

//Checks if error has occured and ends program after logging. 
//Only use for Fatal errors
func Check_Error(e error){
	if(e != nil){
		log.Fatal(e)
	}
}

func Color_Println (c string, text string) {
	var s string;
	if (c == "red") {
		s = "\x1b[0;31m"

	} else if (c == "green") {
		s = "\x1b[0;32m"

	} else if (c == "orange") {
		s = "\x1b[0;33m"

	} else if (c == "blue") {
		s = "\x1b[0;34m"

	} else if (c == "purple") {
		s = "\x1b[0;35m"

	} else if (c == "cyan") {
		s = "\x1b[0;36m"

	} else {
		//Default
		s = "\x1b[0;0m"
	}
	
	s += text
	s +="\x1b[0;0m"

	fmt.Println(s)
}