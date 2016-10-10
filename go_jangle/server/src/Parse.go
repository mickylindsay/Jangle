package main

//Parse function: takes in type User from User.go and byte array recieved from client
//Identifies what type of message is being recieved and decides what type of message to send
func Parse_Data (user *User, data []byte) {

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

	//Client recieve type codes
	var recieve_userid byte = 48
	var recieve_display_name byte = 49
	var recieve_serverid byte = 50
	var recieve_server_display_name byte = 51
	var recieve_roomid byte = 52
	var recieve_room_display_name byte = 53

	//Client send type codes
	var send_new_display_name byte = 64
	var send_new_server_display_name byte = 65
	var send_new_room_display_name byte = 66

	//Status of client type codes
	var status_change byte = 80
	var status_broadcast byte = 81
	var server_change byte = 82
	var server_broadcast byte = 83
	var room_change byte = 84
	var room_broadcast byte = 85

	//Error type codes
	var error_check byte = 255

	//Initializes Message type
	var m Message

	//Compares first byte of data byte array to all code cases
	if(data[0] == create_user) {
		m = Username_Password{
			code: data[0],
			username: data[1:20],
			password: data[21:]}

			//Calls User_Create to check if success or fail
			err := User_Create(data[1:20], data[21:])
			if(err == nil) {
				data[0] = login_success
			} else {
				data[0] = create_user_fail
			}
			Parse_Data(user, data)
	
	} else if(data[0] == create_user_fail) {
		m = Base{
			code: data[0]}

			//Calls Write to send message to a user that does not have a userid
			user.Write(m.Build_Message())
	
	} else if(data[0] == login) {
		m = Username_Password{
			code: data[0],
			username: data[1:20],
			password: data[21:]}

			//Calls User_Login to check if success or fail
			id, err := User_Login(data[1:20], data[21:])
			if(err == nil) {
				data[0] = login_success
				copy(data[1:4], Int_Converter(id))
			} else {
				data[0] = login_fail
			}
			user.id = id
			Parse_Data(user, data)
		
	} else if(data[0] == login_fail) {
		m = Base{
			code: data[0]}

			//Calls Write to send message to a user that does not have a userid
			user.Write(m.Build_Message())
	
	} else if(data[0] == login_success) {
		m = Userid{
			code: data[0],
			userid: data[1:4]}

			//Sends login success to user
			Send_Message(user, m)
	
	} else if(data[0] == message_client_send) {
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
	
	} else if(data[0] == message_client_recieve) {
		m = Message_Recieve{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			time: data[13:16],
			text: data[17:]}

			//Send all users code type 17 message
			Send_Broadcast(m)
	
	} else if(data[0] == request_n_messages) {
		m = Multi_Message{
			code: data[0],
			offset: data[1]}

			//Send user multiple code type 17 messages depending on the offset value
			num := uint(data[1])
			messages,err := Request_Offset_Messages(num)
			Check_Error(err)
			for i := 0; i < len(messages); i++ {
				Send_Message(user, messages[i])
			}
	
	} else if(data[0] == request_all_userid) {
		m = Base{
			code: data[0]}

			//Send user all requested userid on the connected server
			messages,err := Request_Userid_Messages(user.serverid)
			Check_Error(err)
			for i := 0; i < len(messages); i++ {
				Send_Message(user, messages[i])
			}
	
	} else if(data[0] == request_display_name) {
		m = Requested_Userid{
			code: data[0],
			requested_userid: data[1:4]}

			//Converts message to code type 49
			num := Byte_Converter(data[1:4])
			display_name := Request_Display_Name(num)
			new_data := make([]byte, len(display_name) + 5)
			new_data[0] = recieve_display_name
			copy(new_data[1:4], data[1:4])
			copy(new_data[5:], display_name[:])
			Parse_Data(user, new_data)

	} else if(data[0] == request_all_serverid) {
		m = Requested_Userid{
			code: data[0],
			requested_userid: data[1:4]}

			//Send user all serverid from requested userid
			num := Byte_Converter(data[1:4])
			messages,err := Request_Serverid_Messages(num)
			Check_Error(err)
			for i := 0; i < len(messages); i++ {
				Send_Message(user, messages[i])
			}
	
	} else if(data[0] == request_server_display_name) {
		m = Requested_Serverid{
			code: data[0],
			requested_serverid: data[1:4]}

			//Converts message to code type 51
			num := Byte_Converter(data[1:4])
			server_display_name := Request_Server_Display_Name(num)
			new_data := make([]byte, len(server_display_name) + 5)
			new_data[0] = recieve_server_display_name
			copy(new_data[1:4], data[1:4])
			copy(new_data[5:], server_display_name[:])
			Parse_Data(user, new_data)
	
	} else if(data[0] == request_all_roomid) {
		m = Requested_Serverid{
			code: data[0],
			requested_serverid: data[1:4]}

			//Send user all roomid from requested serverid
			num := Byte_Converter(data[1:4])
			messages,err := Request_Roomid_Messages(num)
			Check_Error(err)
			for i := 0; i < len(messages); i++ {
				Send_Message(user, messages[i])
			}
	
	} else if(data[0] == request_room_display_name) {
		m = Requested_Serverid_Roomid{
			code: data[0],
			requested_serverid: data[1:4],
			requested_roomid: data[5:8]}

			//Converts message to code type 53
			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			room_display_name := Request_Room_Display_Name(num1, num2)
			new_data := make([]byte, len(r0om_display_name) + 9)
			new_data[0] = recieve_room_display_name
			copy(new_data[1:4], data[1:4])
			copy(new_data[5:8], data[5:8])
			copy(new_data[9:], room_display_name[:])
			Parse_Data(user, new_data)
	
	} else if(data[0] == recieve_userid) {
		m = Requested_Userid{
			code: data[0],
			requested_userid: data[1:4]}

			//Sends user the requested userid
			Send_Message(user, m)
	
	} else if(data[0] == recieve_display_name) {
		m = Display_Name{
			code: data[0],
			requested_userid: data[5:8],
			display_name: data[13:]}

			//Sends user the requested display name
			Send_Message(user, m)
	
	} else if(data[0] == recieve_serverid) {
		m = Requested_Serverid_Userid{
			code: data[0],
			requested_serverid: data[1:4],
			requested_userid: data[5:8]}

			//Send user the requested serverid from specified userid
			Send_Message(user, m)
	
	} else if(data[0] == recieve_server_display_name) {
		m = Server_Display_Name{
			code: data[0],
			requested_serverid: data[1:4],
			server_display_name: data[5:]}

			//Send user the requested server display name
			Send_Message(user, m)
	
	} else if(data[0] == recieve_roomid) {
		m = Requested_Serverid_Roomid{
			code: data[0],
			requested_serverid: data[1:4],
			requested_roomid: data[5:8]}

			//Send user requested roomid from specified serverid
			Send_Message(user, m)
	
	} else if(data[0] == recieve_room_display_name) {
		m = Room_Display_Name{
			code: data[0],
			requested_serverid: data[1:4],
			requested_roomid: data[5:8],
			room_display_name: data[9:]}

			//Send user room display name from specificed severid and roomid
			Send_Message(user, m)
	
	} else if(data[0] == send_new_display_name) {
		m = New_Display_Name{
			code: data[0],
			new_display_name: data[1:]}
	
	} else if(data[0] == send_new_server_display_name) {
		m = New_Server_Display_Name{
			code: data[0],
			serverid: data[1:4],
			new_server_display_name: data[5:]}
	
	} else if(data[0] == send_new_room_display_name) {
		m = New_Room_Display_Name{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			new_room_display_name: data[9:]}
	
	} else if(data[0] == status_change) {
		m = Status{
			code: data[0],
			status: data[1]}

	} else if(data[0] == status_broadcast) {
		m = Userid_Status{
			code: data[0],
			userid: data[1:4],
			status: data[5]}
		
	} else if(data[0] == server_change) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4]}
		
	} else if(data[0] == server_broadcast) {
		m = Serverid_Userid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}
		
	} else if(data[0] == room_change) {
		m = Roomid{
			code: data[0],
			roomid: data[1:4]}
		
	} else if(data[0] == room_broadcast) {
		m = Roomid_Userid{
			code: data[0],
			roomid: data[1:4],
			userid: data[5:8]}
		
	} else if(data[0] == error_check) {
		m = Text{
			code: data[0],
			text: data[1:]}
		
	}
}