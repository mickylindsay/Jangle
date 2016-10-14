package main

//Initializes Message type
var m Message

func Message0 (user *User, data []byte) {
	m = Username_Password{
			code: data[0],
			username: data[1:20],
			password: data[21:]}

			//Calls User_Create to check if success or fail
			err := User_Create(data[1:20], data[21:])

			if (err == nil) {
				data[0] = login_success
			} else {
				data[0] = create_user_fail
			}

			Parse_Data(user, data)
}

func Message1 (user *User, data []byte) {
	m = Base{
			code: data[0]}

			//Calls Write to send message to a user that does not have a userid
			user.Write(m.Build_Message())
}

func Message2 (user *User, data []byte) {
	m = Username_Password{
			code: data[0],
			username: data[1:20],
			password: data[21:]}

			//Calls User_Login to check if success or fail
			id, err := User_Login(data[1:20], data[21:])

			if (err == nil) {
				data[0] = login_success
				copy(data[1:4], Int_Converter(id))
			} else {
				data[0] = login_fail
			}

			user.id = id
			Parse_Data(user, data)
}

func Message3 (user *User, data []byte) {
	m = Base{
			code: data[0]}

			//Calls Write to send message to a user that does not have a userid
			user.Write(m.Build_Message())
}

func Message4 (user *User, data []byte) {
	m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Sends login success to user
			Send_Message(user, m)
}

func Message16 (user *User, data []byte) {
	m = Message_Send{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			text: data[13:]}

			//Sends message to database
			err := Message_Create(user, data[13:])
			Check_Error(err)

			//Converts message to code type 17
			data[0] = message_client_recieve
			data = Time_Stamp(data)
			Parse_Data(user, data)
}

func Message17 (user *User, data []byte) {
	m = Message_Recieve{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			time: data[13:16],
			text: data[17:]}

			//Send all users code type 17 message
			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			Send_Broadcast_Server_Room(num1, num2, m)
}

func Message32 (user *User, data []byte) {
	m = Multi_Message{
			code: data[0],
			offset: data[1]}

			//Send user multiple code type 17 messages depending on the offset value
			num := uint(data[1])
			messages, err := Request_Offset_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}
}

func Message33 (user *User, data []byte) {
	m = Base{
			code: data[0]}

			//Send user all requested userid on the connected server
			messages,err := Request_Userid_Messages(user.serverid)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}
}

func Message34 (user *User, data []byte) {
	m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Converts message to code type 49
			num := Byte_Converter(data[1:4])
			requested_display_name, err := Request_Display_Name(user.serverid, num)
			Check_Error(err)

			new_m := Display_Name{
				code: recieve_display_name,
				userid: data[1:4],
				display_name: requested_display_name}

				Parse_Data(user, new_m.Build_Message())
}

func Message35 (user *User, data []byte) {
	m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Send user all serverid from requested userid
			num := Byte_Converter(data[1:4])
			messages, err := Request_Serverid_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}
}

func Message36 (user *User, data []byte) {
	m = Serverid{
			code: data[0],
			serverid: data[1:4]}

			//Converts message to code type 51
			num := Byte_Converter(data[1:4])
			requested_server_display_name, err := Request_Server_Display_Name(num)
			Check_Error(err)

			new_m := Server_Display_Name{
				code: recieve_server_display_name,
				serverid: data[1:4],
				server_display_name: requested_server_display_name}

				Parse_Data(user, new_m.Build_Message())
}

func Message37 (user *User, data []byte) {
	m = Serverid{
			code: data[0],
			serverid: data[1:4]}

			//Send user all roomid from requested serverid
			num := Byte_Converter(data[1:4])
			messages, err := Request_Roomid_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}
}

func Message38 (user *User, data []byte) {
	m = Serverid_Roomid{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8]}

			//Converts message to code type 53
			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			requested_room_display_name, err := Request_Room_Display_Name(num1, num2)
			Check_Error(err)

			new_m := Room_Display_Name{
				code: recieve_room_display_name,
				serverid: data[1:4],
				roomid: data[5:8],
				room_display_name: requested_room_display_name}
			
				Parse_Data(user, new_m.Build_Message())
}

func Message48 (user *User, data []byte) {
	m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Sends user the requested userid
			Send_Message(user, m)
}

func Message49 (user *User, data []byte) {
	m = Display_Name{
			code: data[0],
			userid: data[5:8],
			display_name: data[13:]}

			//Sends user the requested display name
			Send_Message(user, m)
}

func Message50 (user *User, data []byte) {
	m = Serverid_Userid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}

			//Send user the requested serverid from specified userid
			Send_Message(user, m)
}

func Message51 (user *User, data []byte) {
	m = Server_Display_Name{
			code: data[0],
			serverid: data[1:4],
			server_display_name: data[5:]}

			//Send user the requested server display name
			Send_Message(user, m)
}

func Message52 (user *User, data []byte) {
	m = Serverid_Roomid{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8]}

			//Send user requested roomid from specified serverid
			Send_Message(user, m)
}

func Message53 (user *User, data []byte) {
	m = Room_Display_Name{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			room_display_name: data[9:]}

			//Send user room display name from specificed severid and roomid
			Send_Message(user, m)
}

func Message64 (user *User, data []byte) {
	m = New_Display_Name{
			code: data[0],
			new_display_name: data[1:]}

			//Sets a new display name specific to the connected server
			err := Set_New_Display_Name(user.serverid, user.id, data[1:])
			Check_Error(err)

			arr := Int_Converter(user.id)
			new_m := Display_Name{
				code: recieve_display_name,
				userid: arr,
				display_name: data[1:]}

				Send_Broadcast_Server(user.serverid, new_m)
}

func Message65 (user *User, data []byte) {
	m = New_Server_Display_Name{
			code: data[0],
			serverid: data[1:4],
			new_server_display_name: data[5:]}

			//
			/*num := Byte_Converter(data[1:4])
			messages, err := Set_New_Server_Display_Name(num, data[5:])
			Check_Error(err)

			for i := 0, i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}*/
}

func Message66 (user *User, data []byte) {
	m = New_Room_Display_Name{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			new_room_display_name: data[9:]}

			//
			/*num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			message, err := Set_New_Room_Display_Name(num1, num2, data[9:])
			Check_Error(err)

			Send_Broadcast_Server(num1, message)*/
}

func Message80 (user *User, data []byte) {
	m = Status{
			code: data[0],
			status: data[1]}

			/*num := uint(data[1])
			message, err := Set_Status(user.id, num)
			Check_Error(err)

			Parse_Data(user, message.Build_Message())*/
}

func Message81 (user *User, data []byte) {
	m = Userid_Status{
			code: data[0],
			userid: data[1:4],
			status: data[5]}

			//
			//Send_Broadcast(m)
}

func Message82 (user *User, data []byte) {
	m = Serverid{
			code: data[0],
			serverid: data[1:4]}

			//
			/*num := Byte_Converter(data[1:4])
			message, err := Set_Server(num, user.id)
			Check_Error(err)

			Parse_Data(user, message.Build_Message())*/
}

func Message83 (user *User, data []byte) {
	m = Serverid_Userid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}

			//
			/*num := Byte_Converter(data[1:4])
			Send_Broadcast_Server(num, m)*/
}

func Message84 (user *User, data []byte) {
	//
			/*num := Byte_Converter(data[1:4])
			message, err := Set_Room(user.serverid,num,user.id)
			Check_Error(err)

			Parse_Data(user, message)*/
}

func Message85 (user *User, data []byte) {
	//
			//Send_Broadcast(m)
}

func Message255 (user *User, data []byte) {

}

//Parse function: takes in type User from User.go and byte array recieved from client
//Identifies what type of message is being recieved and decides what type of message to send
func Parse_Data (user *User, data []byte) {

	//Compares first byte of data byte array to all code cases
	if (data[0] == create_user) {
		m = Username_Password{
			code: data[0],
			username: data[1:20],
			password: data[21:]}

			//Calls User_Create to check if success or fail
			err := User_Create(data[1:20], data[21:])

			if (err == nil) {
				data[0] = login_success
			} else {
				data[0] = create_user_fail
			}

			Parse_Data(user, data)
	
	} else if (data[0] == create_user_fail) {
		m = Base{
			code: data[0]}

			//Calls Write to send message to a user that does not have a userid
			user.Write(m.Build_Message())
	
	} else if (data[0] == login) {
		m = Username_Password{
			code: data[0],
			username: data[1:20],
			password: data[21:]}

			//Calls User_Login to check if success or fail
			id, err := User_Login(data[1:20], data[21:])

			if (err == nil) {
				data[0] = login_success
				copy(data[1:4], Int_Converter(id))
			} else {
				data[0] = login_fail
			}

			user.id = id
			Parse_Data(user, data)
		
	} else if (data[0] == login_fail) {
		m = Base{
			code: data[0]}

			//Calls Write to send message to a user that does not have a userid
			user.Write(m.Build_Message())
	
	} else if (data[0] == login_success) {
		m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Sends login success to user
			Send_Message(user, m)
	
	} else if (data[0] == message_client_send) {
		m = Message_Send{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			text: data[13:]}

			//Sends message to database
			err := Message_Create(user, data[13:])
			Check_Error(err)

			//Converts message to code type 17
			data[0] = message_client_recieve
			data = Time_Stamp(data)
			Parse_Data(user, data)
	
	} else if (data[0] == message_client_recieve) {
		m = Message_Recieve{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			time: data[13:16],
			text: data[17:]}

			//Send all users code type 17 message
			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			Send_Broadcast_Server_Room(num1, num2, m)
	
	} else if (data[0] == request_n_messages) {
		m = Multi_Message{
			code: data[0],
			offset: data[1]}

			//Send user multiple code type 17 messages depending on the offset value
			num := uint(data[1])
			messages, err := Request_Offset_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}
	
	} else if (data[0] == request_all_userid) {
		m = Base{
			code: data[0]}

			//Send user all requested userid on the connected server
			messages,err := Request_Userid_Messages(user.serverid)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}
	
	} else if (data[0] == request_display_name) {
		m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Converts message to code type 49
			num := Byte_Converter(data[1:4])
			requested_display_name, err := Request_Display_Name(user.serverid, num)
			Check_Error(err)

			new_m := Display_Name{
				code: recieve_display_name,
				userid: data[1:4],
				display_name: requested_display_name}

				Parse_Data(user, new_m.Build_Message())

	} else if (data[0] == request_all_serverid) {
		m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Send user all serverid from requested userid
			num := Byte_Converter(data[1:4])
			messages, err := Request_Serverid_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}
	
	} else if (data[0] == request_server_display_name) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4]}

			//Converts message to code type 51
			num := Byte_Converter(data[1:4])
			requested_server_display_name, err := Request_Server_Display_Name(num)
			Check_Error(err)

			new_m := Server_Display_Name{
				code: recieve_server_display_name,
				serverid: data[1:4],
				server_display_name: requested_server_display_name}

				Parse_Data(user, new_m.Build_Message())
	
	} else if (data[0] == request_all_roomid) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4]}

			//Send user all roomid from requested serverid
			num := Byte_Converter(data[1:4])
			messages, err := Request_Roomid_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}
	
	} else if (data[0] == request_room_display_name) {
		m = Serverid_Roomid{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8]}

			//Converts message to code type 53
			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			requested_room_display_name, err := Request_Room_Display_Name(num1, num2)
			Check_Error(err)

			new_m := Room_Display_Name{
				code: recieve_room_display_name,
				serverid: data[1:4],
				roomid: data[5:8],
				room_display_name: requested_room_display_name}
			
				Parse_Data(user, new_m.Build_Message())
	
	} else if (data[0] == recieve_userid) {
		m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Sends user the requested userid
			Send_Message(user, m)
	
	} else if (data[0] == recieve_display_name) {
		m = Display_Name{
			code: data[0],
			userid: data[5:8],
			display_name: data[13:]}

			//Sends user the requested display name
			Send_Message(user, m)
	
	} else if (data[0] == recieve_serverid) {
		m = Serverid_Userid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}

			//Send user the requested serverid from specified userid
			Send_Message(user, m)
	
	} else if (data[0] == recieve_server_display_name) {
		m = Server_Display_Name{
			code: data[0],
			serverid: data[1:4],
			server_display_name: data[5:]}

			//Send user the requested server display name
			Send_Message(user, m)
	
	} else if (data[0] == recieve_roomid) {
		m = Serverid_Roomid{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8]}

			//Send user requested roomid from specified serverid
			Send_Message(user, m)
	
	} else if (data[0] == recieve_room_display_name) {
		m = Room_Display_Name{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			room_display_name: data[9:]}

			//Send user room display name from specificed severid and roomid
			Send_Message(user, m)
	
	} else if (data[0] == send_new_display_name) {
		m = New_Display_Name{
			code: data[0],
			new_display_name: data[1:]}

			//Sets a new display name specific to the connected server
			err := Set_New_Display_Name(user.serverid, user.id, data[1:])
			Check_Error(err)

			arr := Int_Converter(user.id)
			new_m := Display_Name{
				code: recieve_display_name,
				userid: arr,
				display_name: data[1:]}

				Send_Broadcast_Server(user.serverid, new_m)

	} else if (data[0] == send_new_server_display_name) {
		m = New_Server_Display_Name{
			code: data[0],
			serverid: data[1:4],
			new_server_display_name: data[5:]}

			//
			/*num := Byte_Converter(data[1:4])
			messages, err := Set_New_Server_Display_Name(num, data[5:])
			Check_Error(err)

			for i := 0, i < len(messages); i++ {
				Parse_Data(user, messages[i].Build_Message())
			}*/

	} else if (data[0] == send_new_room_display_name) {
		m = New_Room_Display_Name{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			new_room_display_name: data[9:]}

			//
			/*num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			message, err := Set_New_Room_Display_Name(num1, num2, data[9:])
			Check_Error(err)

			Send_Broadcast_Server(num1, message)*/
	
	} else if (data[0] == status_change) {
		m = Status{
			code: data[0],
			status: data[1]}

			/*num := uint(data[1])
			message, err := Set_Status(user.id, num)
			Check_Error(err)

			Parse_Data(user, message.Build_Message())*/

	} else if (data[0] == status_broadcast) {
		m = Userid_Status{
			code: data[0],
			userid: data[1:4],
			status: data[5]}

			//
			//Send_Broadcast(m)
		
	} else if (data[0] == server_change) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4]}

			//
			/*num := Byte_Converter(data[1:4])
			message, err := Set_Server(num, user.id)
			Check_Error(err)

			Parse_Data(user, message.Build_Message())*/
		
	} else if (data[0] == server_broadcast) {
		m = Serverid_Userid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}

			//
			/*num := Byte_Converter(data[1:4])
			Send_Broadcast_Server(num, m)*/
		
	} else if (data[0] == room_change) {
		m = Roomid{
			code: data[0],
			roomid: data[1:4]}

			//
			/*num := Byte_Converter(data[1:4])
			message, err := Set_Room(user.serverid,num,user.id)
			Check_Error(err)

			Parse_Data(user, message)*/
		
	} else if (data[0] == room_broadcast) {
		m = Roomid_Userid{
			code: data[0],
			roomid: data[1:4],
			userid: data[5:8]}

			//
			//Send_Broadcast(m)
		
	} else if (data[0] == error_check) {
		m = Text{
			code: data[0],
			text: data[1:]}
		
	}
}