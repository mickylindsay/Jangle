package main

//Initializes Message type
var m Message

//If create user successful, convert message to code type 4
//If create user fail, convert message to code type 1
func Message0 (user *User, data []byte) {

	m = Username_Password{
		code: data[0],
		username: data[1:20],
		password: data[21:]}

			id, err := User_Create(data[1:20], data[21:])

			if (err == nil) {
				data[0] = login_success
				copy(data[1:4], Int_Converter(id))
				user.id = id
				Message4(user,data)
			} else {
				data[0] = create_user_fail
				Message1(user, data)
			}
}

//Writes to user message code type 1, create user fail
func Message1 (user *User, data []byte) {

	m = Base{
		code: data[0]}

			user.Write(m.Build_Message())
}

//If login successful, convert message to code type 4
//If login fail, convert message to code type 3
func Message2 (user *User, data []byte) {

	m = Username_Password{
		code: data[0],
		username: data[1:20],
		password: data[21:]}

			id, err := User_Login(data[1:20], data[21:])

			if (err == nil) {
				data[0] = login_success
				copy(data[1:4], Int_Converter(id))
				user.id = id
				Message4(user,data)
			} else {
				data[0] = login_fail
				Message3(user, data)
			}
}

//Writes to user message code type 3, login fail
func Message3 (user *User, data []byte) {

	m = Base{
		code: data[0]}

			user.Write(m.Build_Message())
}

//Send message code type 4 to client, login success
func Message4 (user *User, data []byte) {

	m = Userid{
		code: data[0],
		userid: data[1:4]}

			Send_Message(user, m)
}

//Converts message code type 16, message client send
//to message code type 17, message client recieve
func Message16 (user *User, data []byte) {

	m = Message_Send{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		userid: data[9:12],
		text: data[13:]}

			err := Message_Create(user, data[13:])
			Check_Error(err)

			data[0] = message_client_recieve
			data = Time_Stamp(data)
			Message17(user, data)
}

//Sends message code type 17, message client recieve,
//to a chat room on a specific server
func Message17 (user *User, data []byte) {

	m = Message_Recieve{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		userid: data[9:12],
		time: data[13:16],
		text: data[17:]}

		num1 := Byte_Converter(data[1:4])
		num2 := Byte_Converter(data[5:8])
		Send_Broadcast_Server_Room(num1, num2, m)
}

//Requests n message code type 17's, message client recieve, from database 
//dependent on offset value from message code type 32, request n messages
func Message32 (user *User, data []byte) {

	m = Multi_Message{
		code: data[0],
		offset: data[1]}

			num := uint(data[1])
			messages, err := Request_Offset_Messages(user, num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Message17(user, messages[i].Build_Message())
			}
}

//Requests message code type 48's, recieve userid, from database
//which consists of all userids conected to a specific server
func Message33 (user *User, data []byte) {

	m = Base{
		code: data[0]}

			messages,err := Request_Userid_Messages(user.serverid)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Message48(user, messages[i].Build_Message())
			}
}

//Requests a display name from the database from a user connected to a specific server
//from requested userid in message code type 34, requested display name, then
//builds a new byte array in the format of message code type 49, recieve display name
func Message34 (user *User, data []byte) {

	m = Userid{
		code: data[0],
		userid: data[1:4]}

			num := Byte_Converter(data[1:4])
			requested_display_name, err := Request_Display_Name(user.serverid, num)
			Check_Error(err)

			new_data := make([]byte, len(requested_display_name) + 5)
			new_data[0] = recieve_display_name
			copy(new_data[1:4], data[1:4])
			copy(new_data[5:], requested_display_name)

			Message49(user, new_data)				
}

//Requests message code type 50's, recieve serverid, which consists of all the serverids
//that a specific user is connected to from message code type 35, request all serverid
func Message35 (user *User, data []byte) {

	m = Userid{
		code: data[0],
		userid: data[1:4]}

			num := Byte_Converter(data[1:4])
			messages, err := Request_Serverid_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Message50(user, messages[i].Build_Message())
			}
}

//Requests server display name from a specific server from the requested servid in
//message code tye 36, request server display name, then builds
//new byte array in the format of message code type 51, recieve server display name
func Message36 (user *User, data []byte) {

	m = Serverid{
		code: data[0],
		serverid: data[1:4]}

			num := Byte_Converter(data[1:4])
			requested_server_display_name, err := Request_Server_Display_Name(num)
			Check_Error(err)

			new_data := make([]byte, len(requested_server_display_name) + 5)
			new_data[0] = recieve_server_display_name
			copy(new_data[1:4], data[1:4])
			copy(new_data[5:], requested_server_display_name)

			Message51(user, new_data)
}

//Request message code type 52's, recieve roomid, which consist of all the roomids
//on a specific server from the serverid in message code type 37, request all roomid
func Message37 (user *User, data []byte) {

	m = Serverid{
		code: data[0],
		serverid: data[1:4]}

			num := Byte_Converter(data[1:4])
			messages, err := Request_Roomid_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Message52(user, messages[i].Build_Message())
			}
}

//Requests room display name from a specific room on specific server from the requested
//serverid and roomid in message code type 38, request room display name, then builds
//a new byte array in the format of message code type 53, recieve room display name
func Message38 (user *User, data []byte) {

	m = Serverid_Roomid{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8]}

			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			requested_room_display_name, err := Request_Room_Display_Name(num1, num2)
			Check_Error(err)

			new_data := make([]byte, len(requested_room_display_name) + 9)
			new_data[0] = recieve_room_display_name
			copy(new_data[1:8], data[1:8])
			copy(new_data[9:], requested_room_display_name)

			Message53(user, new_data)
}

//Sends message code type 48, recieve userid, to client
func Message48 (user *User, data []byte) {

	m = Userid{
		code: data[0],
		userid: data[1:4]}

			Send_Message(user, m)
}

//Sends message code type 49, recieve display name, to client
func Message49 (user *User, data []byte) {

	m = Display_Name{
		code: data[0],
		userid: data[1:4],
		display_name: data[5:]}

			Send_Message(user, m)
}

//Sends message code type 50, recieve serverid, to client
func Message50 (user *User, data []byte) {

	m = Serverid_Userid{
		code: data[0],
		serverid: data[1:4],
		userid: data[5:8]}

			Send_Message(user, m)
}

//Sends message code type 51, recieve server display name, to client
func Message51 (user *User, data []byte) {

	m = Server_Display_Name{
		code: data[0],
		serverid: data[1:4],
		server_display_name: data[5:]}

			Send_Message(user, m)
}

//Sends message code type 52, recieve roomid, to client
func Message52 (user *User, data []byte) {

	m = Serverid_Roomid{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8]}

			Send_Message(user, m)
}

//Sends message code type 53, recieve room display name, to client
func Message53 (user *User, data []byte) {

	m = Room_Display_Name{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		room_display_name: data[9:]}

			Send_Message(user, m)
}

//Replaces the user's display name with the new display name in message code type 64,
//send new display name, then sends message code type 49, recieve display name, to
//all users on the user's connected server
func Message64 (user *User, data []byte) {

	m = New_Display_Name{
		code: data[0],
		new_display_name: data[1:]}

			err := Set_New_Display_Name(user.serverid, user.id, data[1:])
			Check_Error(err)

			arr := Int_Converter(user.id)
			new_m := Display_Name{
				code: recieve_display_name,
				userid: arr,
				display_name: data[1:]}

				Send_Broadcast_Server(user.serverid, new_m)
}

//TODO
func Message65 (user *User, data []byte) {

	m = New_Server_Display_Name{
		code: data[0],
		serverid: data[1:4],
		new_server_display_name: data[5:]}

			/*num := Byte_Converter(data[1:4])
			err := Set_New_Server_Display_Name(num, data[5:])
			Check_Error(err)

			new_m := Server_Display_Name{
				code: recieve_server_display_name,
				serverid: data[1:4],
				server_display_name: data[5:]}

				Send_Broadcast_Members(num, new_m)*/
}

//TODO
func Message66 (user *User, data []byte) {

	m = New_Room_Display_Name{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		new_room_display_name: data[9:]}

			/*num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			err := Set_New_Room_Display_Name(num1, num2, data[9:])
			Check_Error(err)

			new_m := Room_Display_Name{
				code: recieve_room_display_name,
				serverid: data[1:4],
				roomid: data[5:8],
				room_display_name: data[9:]}

				Send_Broadcast_Server(num1, new_m)*/
}

//TODO
func Message80 (user *User, data []byte) {

	m = Status{
		code: data[0],
		status: data[1]}

			/*num := uint(data[1])
			message, err := Set_Status(user.id, num)
			Check_Error(err)

			Parse_Data(user, message.Build_Message())*/
}

//TODO
func Message81 (user *User, data []byte) {

	m = Userid_Status{
		code: data[0],
		userid: data[1:4],
		status: data[5]}

			//Send_Broadcast(m)
}

//TODO
func Message82 (user *User, data []byte) {

	m = Serverid{
		code: data[0],
		serverid: data[1:4]}

			/*num := Byte_Converter(data[1:4])
			message, err := Set_Server(num, user.id)
			Check_Error(err)

			Parse_Data(user, message.Build_Message())*/
}

//TODO
func Message83 (user *User, data []byte) {

	m = Serverid_Userid{
		code: data[0],
		serverid: data[1:4],
		userid: data[5:8]}

			/*num := Byte_Converter(data[1:4])
			Send_Broadcast_Server(num, m)*/
}

//TODO
func Message84 (user *User, data []byte) {

	m = Roomid{
		code: data[0],
		roomid: data[1:4]}

			/*num := Byte_Converter(data[1:4])
			message, err := Set_Room(user.serverid,num,user.id)
			Check_Error(err)

			Parse_Data(user, message)*/
}

//TODO
func Message85 (user *User, data []byte) {

	m = Roomid_Userid{
		code: data[0],
		roomid: data[1:4],
		userid: data[5:8]}

			//Send_Broadcast(m)
}

//TODO
func Message255 (user *User, data []byte) {
	
	m = Text{
		code: data[0],
		text: data[1:]}
}

//Master function: takes paramaters type User struct and byte array
//byte array is the message that is recieved from the client
//the type User struct is a reference to the connection that represents
//the client side user that is associated with the byte array message
//this function determines what type of message is being recieved
//and calls the appropiate function based off the code type
func Parse_Data (user *User, data []byte) {

	if (data[0] == create_user) {

		Message0(user, data)
	
	} else if (data[0] == create_user_fail) {

		Message1(user, data)
	
	} else if (data[0] == login) {

		Message2(user, data)
		
	} else if (data[0] == login_fail) {

		Message3(user, data)
	
	} else if (data[0] == login_success) {

		Message4(user, data)
	
	} else if (data[0] == message_client_send) {

		Message16(user, data)
	
	} else if (data[0] == message_client_recieve) {

		Message17(user, data)
	
	} else if (data[0] == request_n_messages) {

		Message32(user, data)
	
	} else if (data[0] == request_all_userid) {

		Message33(user, data)
	
	} else if (data[0] == request_display_name) {

		Message34(user, data)

	} else if (data[0] == request_all_serverid) {

		Message35(user, data)
	
	} else if (data[0] == request_server_display_name) {

		Message36(user, data)
	
	} else if (data[0] == request_all_roomid) {

		Message37(user, data)
	
	} else if (data[0] == request_room_display_name) {

		Message38(user, data)
	
	} else if (data[0] == recieve_userid) {

		Message48(user, data)
	
	} else if (data[0] == recieve_display_name) {

		Message49(user, data)
	
	} else if (data[0] == recieve_serverid) {

		Message50(user, data)
	
	} else if (data[0] == recieve_server_display_name) {

		Message51(user, data)
	
	} else if (data[0] == recieve_roomid) {

		Message52(user, data)
	
	} else if (data[0] == recieve_room_display_name) {

		Message53(user, data)
	
	} else if (data[0] == send_new_display_name) {

		Message64(user, data)

	} else if (data[0] == send_new_server_display_name) {

		Message65(user, data)

	} else if (data[0] == send_new_room_display_name) {

		Message66(user, data)
	
	} else if (data[0] == status_change) {

		Message80(user, data)

	} else if (data[0] == status_broadcast) {

		Message81(user, data)
		
	} else if (data[0] == server_change) {

		Message82(user, data)
		
	} else if (data[0] == server_broadcast) {

		Message83(user, data)
		
	} else if (data[0] == room_change) {

		Message84(user, data)
		
	} else if (data[0] == room_broadcast) {

		Message85(user, data)
		
	} else if (data[0] == error_check) {

		Message255(user, data)
		
	}
}