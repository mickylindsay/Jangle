package main

//Constructs generic Message type
type Message interface {
	Build_Message() []byte
}

//[code:1]
type Base struct {
	code byte
}

//Builds Message type 
func(m Base) Build_Message() []byte {
	message := make([]byte, 1)
	message[0] = m.code
	return message[:]
}

//[code:1,username:20,password:]
type Username_Password struct {
	code byte
	username []byte
	password []byte
}

func(m Username_Password) Build_Message() []byte {
	message := make([]byte, 21 + len(m.password))
	message[0] = m.code
	copy(message[1:20], m.username[:])
	copy(message[21:], m.password[:])
	return message
}

//[code:1,userid:4]
type Userid struct {
	code byte
	userid []byte
}

func(m Userid) Build_Message() []byte {
	message := make([]byte, 5)
	message[0] = m.code
	copy(message[1:4], m.userid[:])
	return message
}

//[code:1,requested_userid:4]
type Requested_Userid struct {
	code byte
	requested_userid []byte
}

func(m Requested_Userid) Build_Message() []byte {
	message := make([]byte, 5)
	message[0] = m.code
	copy(message[1:4], m.requested_userid[:])
	return message
}

//[code:1,serverid:4]
type Serverid struct {
	code byte
	serverid []byte
}

func(m Serverid) Build_Message() []byte {
	message := make([]byte, 5)
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	return message
}

//[code:1,requested_serverid:4]
type Requested_Serverid struct {
	code byte
	requested_serverid []byte
}

func(m Requested_Serverid) Build_Message() []byte {
	message := make([]byte, 5)
	message[0] = m.code
	copy(message[1:4], m.requested_serverid[:])
	return message
}

//[code:1,serverid:4,userid:4]
type Serverid_Userid struct {
	code byte
	serverid []byte
	userid []byte
}

func(m Serverid_Userid) Build_Message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.userid[:])
	return message
}

//[code:1,requested_serverid:4,requested_userid:4]
type Requested_Serverid_Userid struct {
	code byte
	requested_serverid []byte
	requested_userid []byte
}

func(m Requested_Serverid_Userid) Build_Message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.requested_serverid[:])
	copy(message[5:8], m.requested_userid[:])
	return message
}

//[code:1,roomid:4]
type Roomid struct {
	code byte
	roomid []byte
}

func(m Roomid) Build_Message() []byte {
	message := make([]byte, 5)
	message[0] = m.code
	copy(message[1:4], m.roomid[:])
	return message
}

//[code:1,roomid:4,userid:4]
type Roomid_Userid struct {
	code byte
	roomid []byte
	userid []byte
}

func(m Roomid_Userid) Build_Message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.roomid[:])
	copy(message[5:8], m.userid[:])
	return message
}

//[code:1,requested_serverid:4,requested_roomid:4]
type Requested_Serverid_Roomid struct {
	code byte
	requested_serverid []byte
	requested_roomid []byte
}

func(m Requested_Serverid_Roomid) Build_Message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.requested_serverid[:])
	copy(message[5:8], m.requested_roomid[:])
	return message
}

//[code:1,serverid:4,roomid:4,userid:4,text:]
type Message_Send struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
	text []byte
}

func(m Message_Send) Build_Message() []byte {
	message := make([]byte, 13 + len(m.text))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:12], m.userid[:])
	copy(message[13:], m.text[:])
	return message
}

//[code:1,serverid:4.roomid:4,userid:4,time:4,text:]
type Message_Recieve struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
	time []byte
	text []byte
}

func(m Message_Recieve) Build_Message() []byte {
	message := make([]byte, 17 + len(m.text))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:12], m.userid[:])
	copy(message[13:16], m.time[:])
	copy(message[17:], m.text[:])
	return message
}

//[code:1,offset:1]
type Multi_Message struct {
	code byte
	offset byte
}

func(m Multi_Message) Build_Message() []byte {
	message := make([]byte, 2)
	message[0] = m.code
	message[1] = m.offset
	return message
}

//[code:1,requested_userid:4,display_name:]
type Display_Name struct {
	code byte
	requested_userid []byte
	display_name []byte
}

func(m Display_Name) Build_Message() []byte {
	message := make([]byte, 5 + len(m.display_name))
	message[0] = m.code
	copy(message[1:4], m.requested_userid[:])
	copy(message[5:], m.display_name[:])
	return message
}

//[code:1,requested_serverid:4,server_display_name:]
type Server_Display_Name struct {
	code byte
	requested_serverid []byte
	server_display_name []byte
}

func(m Server_Display_Name) Build_Message() []byte {
	message := make([]byte, 5 + len(m.server_display_name))
	message[0] = m.code
	copy(message[1:4], m.requested_serverid[:])
	copy(message[5:], m.server_display_name[:])
	return message
}

//[code:1,requested_serverid:4,requested_roomid:4,room_display_name:]
type Room_Display_Name struct {
	code byte
	requested_serverid []byte
	requested_roomid []byte
	room_display_name []byte
}

func(m Room_Display_Name) Build_Message() []byte {
	message := make([]byte, 9 + len(m.room_display_name))
	message[0] = m.code
	copy(message[1:4], m.requested_serverid[:])
	copy(message[5:8], m.requested_roomid[:])
	copy(message[9:], m.room_display_name[:])
	return message
}

//[code:1,new_display_name:]
type New_Display_Name struct {
	code byte
	new_display_name []byte
}

func(m New_Display_Name) Build_Message() []byte {
	message := make([]byte, 1 + len(m.new_display_name))
	message[0] = m.code
	copy(message[1:], m.new_display_name[:])
	return message
}

//[code:1,serverid:4,new_server_display_name:]
type New_Server_Display_Name struct {
	code byte
	serverid []byte
	new_server_display_name []byte
}

func(m New_Server_Display_Name) Build_Message() []byte {
	message := make([]byte, 5 + len(m.new_server_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:], m.new_server_display_name[:])
	return message
}

//[code:1,serverid:4,roomid:4,new_room_display_name:]
type New_Room_Display_Name struct {
	code byte
	serverid []byte
	roomid []byte
	new_room_display_name []byte
}

func(m New_Room_Display_Name) Build_Message() []byte {
	message := make([]byte, 9 + len(m.new_room_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:], m.new_room_display_name[:])
	return message
}

//[code:1,status:1]
type Status struct {
	code byte
	status byte
}

func(m Status) Build_Message() []byte {
	message := make([]byte, 2)
	message[0] = m.code
	message[1] = m.status
	return message
}

//[code:1,userid:4,status:1]
type Userid_Status struct {
	code byte
	userid []byte
	status byte
}

func(m Userid_Status) Build_Message() []byte {
	message := make([]byte, 6)
	message[0] = m.code
	copy(message[1:4], m.userid[:])
	message[5] = m.status
	return message
}

//[code:1,text:]
type Text struct {
	code byte
	text []byte
}

func(m Text) Build_Message() []byte {
	message := make([]byte, 1 + len(m.text))
	message[0] = m.code
	copy(message[1:], m.text[:])
	return message
}

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

			//Calls Time_Stamp to convert message to code type 17 or message_client_recieve
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

			//Calls Send_Broadcast because code type 17 messages will be sent to all users
			Send_Broadcast(m)
	
	} else if(data[0] == request_n_messages) {
		m = Multi_Message{
			code: data[0],
			offset: data[1]}

			//Send user multiple code type 17 messages depending on the offset value
			num := uint(data[13])
			messages,err := Request_Offset_Messages(num)
			Check_Error(err)
			for i := 0; i < len(messages); i++ {
				Send_Message(user, messages[i])
			}
	
	} else if(data[0] == request_all_userid) {
		m = Base{
			code: data[0]}
	
	} else if(data[0] == request_display_name) {
		m = Requested_Userid{
			code: data[0],
			requested_userid: data[1:4]}

			//Calls Request_Display_Name given a userid to reference database and builds code type 49 message
			id := Byte_Converter(data[1:4])
			display_name := Request_Display_Name(id)
			new_data := make([]byte, len(display_name) + 13)
			new_data[0] = recieve_display_name
			copy(new_data[1:4], data[1:4])
			copy(new_data[5:], display_name[:])
			Parse_Data(user, new_data)

	} else if(data[0] == request_all_serverid) {
		m = Requested_Userid{
			code: data[0],
			requested_userid: data[1:4]}
	
	} else if(data[0] == request_server_display_name) {
		m = Requested_Serverid{
			code: data[0],
			requested_serverid: data[1:4]}
	
	} else if(data[0] == request_all_roomid) {
		m = Requested_Serverid{
			code: data[0],
			requested_serverid: data[1:4]}
	
	} else if(data[0] == request_room_display_name) {
		m = Requested_Serverid_Roomid{
			code: data[0],
			requested_serverid: data[1:4],
			requested_roomid: data[5:8]}
	
	} else if(data[0] == recieve_userid) {
		m = Requested_Userid{
			code: data[0],
			requested_userid: data[1:4]}
	
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
	
	} else if(data[0] == recieve_server_display_name) {
		m = Server_Display_Name{
			code: data[0],
			requested_serverid: data[1:4],
			server_display_name: data[5:]}
	
	} else if(data[0] == recieve_roomid) {
		m = Requested_Serverid_Roomid{
			code: data[0],
			requested_serverid: data[1:4],
			requested_roomid: data[5:8]}
	
	} else if(data[0] == recieve_room_display_name) {
		m = Room_Display_Name{
			code: data[0],
			requested_serverid: data[1:4],
			requested_roomid: data[5:8],
			room_display_name: data[9:]}
	
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
