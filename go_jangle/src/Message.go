package main

type Message interface {
	Build_message() []byte
}

//[code:1]
type Base struct {
	code byte
}

func(m Base) Build_message() []byte {
	message := make([]byte, 1)
	message[0] = m.code
	return message[:]
}

//[code:1,username:20,password:]
type Username_password struct {
	code byte
	username []byte
	password []byte
}

func(m Username_password) Build_message() []byte {
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

func(m Userid) Build_message() []byte {
	message := make([]byte, 5)
	message[0] = m.code
	copy(message[1:4], m.userid[:])
	return message
}

//[code:1,userid:4,requested_userid:4]
type Double_userid struct {
	code byte
	userid []byte
	requested_userid []byte
}

func(m Double_userid) Build_message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.userid[:])
	copy(message[5:8], m.requested_userid[:])
	return message
}

//[code:1,serverid:4,userid:4]
type Serverid struct {
	code byte
	serverid []byte
	userid []byte
}

func(m Serverid) Build_message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.userid[:])
	return message
}

//[code:1,serverid:4,roomid:4,userid:4]
type Roomid struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
}

func(m Roomid) Build_message() []byte {
	message := make([]byte, 13)
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:12], m.userid[:])
	return message
}

//[code:1,serverid:4,roomid:4,userid:4,text:]
type Message_send struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
	text []byte
}

func(m Message_send) Build_message() []byte {
	message := make([]byte, 13 + len(m.text))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:12], m.userid[:])
	copy(message[13:], m.text[:])
	return message
}

//[code:1,serverid:4.roomid:4,userid:4,time:4,text:]
type Message_recieve struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
	time []byte
	text []byte
}

func(m Message_recieve) Build_message() []byte {
	message := make([]byte, 17 + len(m.text))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:12], m.userid[:])
	copy(message[13:16], m.time[:])
	copy(message[17:], m.text[:])
	return message
}

//[code:1,serverid:4,roomid:4,userid:4,num_message:1]
type Multi_message struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
	num_message byte
}

func(m Multi_message) Build_message() []byte {
	message := make([]byte, 14)
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:12], m.userid[:])
	message[13] = m.num_message
	return message
}

//[code:1,serverid:4,userid:4,requested_userid:4,display_name:]
type Display_name struct {
	code byte
	serverid []byte
	userid []byte
	requested_userid []byte
	display_name []byte
}

func(m Display_name) Build_message() []byte {
	message := make([]byte, 13 + len(m.display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.userid[:])
	copy(message[9:12], m.requested_userid[:])
	copy(message[13:], m.display_name[:])
	return message
}

//[code:1,serverid:4,userid:4,server_display_name:]
type Server_display_name struct {
	code byte
	serverid []byte
	userid []byte
	server_display_name []byte
}

func(m Server_display_name) Build_message() []byte {
	message := make([]byte, 9 + len(m.server_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.userid[:])
	copy(message[9:], m.server_display_name[:])
	return message
}

//[code:1,roomid:4,userid:4,room_display_name:]
type Room_display_name struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
	room_display_name []byte
}

func(m Room_display_name) Build_message() []byte {
	message := make([]byte, 13 + len(m.room_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:12], m.userid[:])
	copy(message[13:], m.room_display_name[:])
	return message
}

//[code:1,serverid:4,userid:4,new_display_name:]
type New_display_name struct {
	code byte
	serverid []byte
	userid []byte
	new_display_name []byte
}

func(m New_display_name) Build_message() []byte {
	message := make([]byte, 9 + len(m.new_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.userid[:])
	copy(message[9:], m.new_display_name[:])
	return message
}

//[code:1,serverid:4,userid:4,new_server_display_name:]
type New_server_display_name struct {
	code byte
	serverid []byte
	userid []byte
	new_server_display_name []byte
}

func(m New_server_display_name) Build_message() []byte {
	message := make([]byte, 9 + len(m.new_server_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.userid[:])
	copy(message[9:], m.new_server_display_name[:])
	return message
}

//[code:1,serverid:4,roomid:4,userid:4,new_room_display_name:]
type New_room_display_name struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
	new_room_display_name []byte
}

func(m New_room_display_name) Build_message() []byte {
	message := make([]byte, 13 + len(m.new_room_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:12], m.userid[:])
	copy(message[13:], m.new_room_display_name[:])
	return message
}

func Parse_data (user User, data []byte) {
	var create_user byte = 0
	var create_user_fail byte = 1
	var login byte = 2
	var login_fail byte = 3
	var login_success byte = 4

	var message_client_send byte = 16
	var message_client_recieve byte = 17

	var request_n_messages byte = 32
	var request_all_userid byte = 33
	var request_display_name byte = 34
	var request_all_serverid byte = 35
	var request_server_display_name byte = 36
	var request_all_roomid byte = 37
	var request_room_display_name byte = 38

	var recieve_userid byte = 48
	var recieve_display_name byte = 49
	var recieve_serverid byte = 50
	var recieve_server_display_name byte = 51
	var recieve_roomid byte = 52
	var recieve_room_display_name byte = 53

	var send_new_display_name byte = 64
	var send_new_server_display_name byte = 65
	var send_new_room_display_name byte = 66

	var m Message

	if(data[0] == create_user) {
		m = Username_password{
			code: data[0],
			username: data[1:20],
			password: data[21:]}

		test := User_Create(m.username, m.password)
		if(test == 1) {
			data[0] = login_success
		} else if(test == -1) {
			data[0] = create_user_fail
		}
		Parse_data(user, data)
	
	} else if(data[0] == create_user_fail) {
		m = Base{
			code: data[0]}

		user.Write(m.Build_message())
	
	} else if(data[0] == login) {
		m = Username_password{
			code: data[0],
			username: data[1:20],
			password: data[21:]}

		test := User_Login(m.username, m.password)
		if(test == 1) {
			data[0] = login_success
		} else if(test == -1) {
			data[0] = login_fail
		}
		Parse_data(user, data)
	
	} else if(data[0] == login_fail) {
		m = Base{
			code: data[0]}

		user.Write(m.Build_message())
	
	} else if(data[0] == login_success) {
		m = Userid{
			code: data[0],
			userid: data[1:4]}

		u := Btye_Converter(m.userid)
		Send_Message(u, m.Build_message())
	
	} else if(data[0] == message_client_send) {
		m = Message_send{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			text: data[13:]}

		data[0] = message_client_recieve
		data = Time_Stamp(data)
		Parse_data(user, data)
	
	} else if(data[0] == message_client_recieve) {
		m = Message_recieve{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			time: data[13:16],
			text: data[17:]}

		Send_Broadcast(m.Build_message())
	
	} else if(data[0] == request_n_messages) {
		m = Multi_message{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			num_message: data[13]}

		num := Btye_Converter(m.num_message)
		for i := 1; i <= num; i++ {
			
		}
	
	} else if(data[0] == request_all_userid) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}
	
	} else if(data[0] == request_display_name) {
		m = Double_userid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8],
			requested_userid: data[9:12]}

	} else if(data[0] == request_all_serverid) {
		m = Double_userid{
			code: data[0],
			userid: data[1:4],
			requested_userid: data[5:8]}
	
	} else if(data[0] == request_server_display_name) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}
	
	} else if(data[0] == request_all_roomid) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}
	
	} else if(data[0] == request_room_display_name) {
		m = Roomid{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12]}
	
	} else if(data[0] == recieve_userid) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}
	
	} else if(data[0] == recieve_display_name) {
		m = Display_name{
			code: data[0],
			userid: data[1:4],
			requested_userid: data[5:8],
			display_name: data[9:]}
	
	} else if(data[0] == recieve_serverid) {
		m = Serverid{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8]}
	
	} else if(data[0] == recieve_server_display_name) {
		m = Server_display_name{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8],
			server_display_name: data[9:]}
	
	} else if(data[0] == recieve_roomid) {
		m = Roomid{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12]}
	
	} else if(data[0] == recieve_room_display_name) {
		m = Room_display_name{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			room_display_name: data[13:]}
	
	} else if(data[0] == send_new_display_name) {
		m = New_display_name{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8],
			new_display_name: data[9:]}
	
	} else if(data[0] == send_new_server_display_name) {
		m = New_server_display_name{
			code: data[0],
			serverid: data[1:4],
			userid: data[5:8],
			new_server_display_name: data[9:]}
	
	} else if(data[0] == send_new_room_display_name) {
		m = New_room_display_name{
			code: data[0],
			serverid: data[1:4],
			roomid: data[5:8],
			userid: data[9:12],
			new_room_display_name: data[13:]}
	
	} else {
		return nil
	}
}