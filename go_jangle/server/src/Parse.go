package main

//Master function: takes parameters type User and byte array. The byte array is the data in from the client and
//the type User represents the client connection that sent the byte array. Passes these parameters to a specific
//function in the below array of funtions corresponding to the first byte (code value) of the byte arrray.
func Parse_Data(user *User, data []byte) Message {
	m := jangle.Parsers[data[0]](user, data)
	return m
}

//Initializes function array that contains all the functions necessary to handle every message code.
func Init_Parse() {
	Parsers := make([]func(user *User, data []byte) Message, 256)

	Parsers[create_user] = Create_User_Message
	Parsers[login] = Login_Message
	Parsers[message_client_send] = Standard_Message
	Parsers[request_n_messages] = Offset_Message
	Parsers[request_all_userid] = Multi_Userid_Message
	Parsers[request_display_name] = Display_Name_Message
	Parsers[request_all_serverid] = Multi_Serverid_Message
	Parsers[request_server_display_name] = Server_Display_Name_Message
	Parsers[request_all_roomid] = Multi_Roomid_Message
	Parsers[request_room_display_name] = Room_Display_Name_Message
	Parsers[request_master_display_name] = Master_Display_Name_Message
	Parsers[request_status] = Status_Message
	Parsers[request_location] = Location_Message
	Parsers[request_user_ip] = User_Ip_Message
	Parsers[request_user_icon] = User_Icon_Message
	Parsers[request_server_icon] = Server_Icon_Message
	Parsers[send_new_display_name] = New_Display_Name_Message
	Parsers[send_new_server_display_name] = New_Server_Display_Name_Message
	Parsers[send_new_room_display_name] = New_Room_Display_Name_Message
	Parsers[send_new_master_display_name] = New_Master_Display_Name_Message
	Parsers[send_new_user_icon] = New_User_Icon_Message
	Parsers[send_new_server_icon] = New_Server_Icon_Message
	Parsers[change_status] = Change_Status_Message
	Parsers[change_location] = Change_Location_Message
	Parsers[create_server] = Create_Server_Message
	Parsers[create_room] = Create_Room_Message

	jangle.Parsers = Parsers
}

//Reads message code type 0, create user; creates a message with code type 4, login success, if the new user succeeds
//to be created the sends the message to the client; creates a message with code type 1, create user fail, if the new
//user fails to be created then writes the data to the user.
func Create_User_Message(user *User, data []byte) Message {
	m := Create_Message(create_user, data[1:21], data[21:])
	id, err := User_Create(m.username, m.password)
	if err == nil {
		user.id = id
		user.serverid = 1
		user.roomid = 1
		Join_Server(user)
		m = Create_Message(login_success, Int_Converter(id))
		Send_Message(user, m)
		user.status = uint(online);
		m = Create_Message(recieve_status, Int_Converter(user.id), byte(user.status), byte(user.muted), byte(user.voice))
		Send_Broadcast_Server(user.serverid, m)
	} else {
		m = Create_Message(create_user_fail)
		user.Write(m.Build_Message())
	}
	return m
}

//Reads message code type 2, login; creates a message with code type 4, login success, if the user succeeds to login
//then sends the message to the client; creates a message with code type 3, login fail, if the user fails to login then
//writes the data to the user.
func Login_Message(user *User, data []byte) Message {
	m := Create_Message(login, data[1:21], data[21:])
	id, err := User_Login(m.username, m.password)
	if err == nil {
		user.id = id
		m = Create_Message(login_success, Int_Converter(id))
		Send_Message(user, m)
		user.status = uint(online);
		m = Create_Message(recieve_status, Int_Converter(user.id), byte(user.status), byte(user.muted), byte(user.voice))
		Send_Broadcast_Server(user.serverid, m)
	} else {
		m = Create_Message(login_fail)
		user.Write(m.Build_Message())
	}
	return m
}

//Reads message code type 16, message client send; if user is not muted; creates a message with code type 17, message
//client recieve, if the new message succeeds to be created then broadcasts the message to the server that the user is
//connected to; checks if the text from the message type is a command; creates a message with code type 255,
//error check, if the new message fails to be created then sends the message to the client.
func Standard_Message(user *User, data []byte) Message {
	m := Create_Message(message_client_send, data[1:5], data[5:9], data[9:13], data[13:])
	if user.muted != uint(user_muted) {
		messageid, err := Message_Create(user, m.text)
		if err != nil {
			m = Create_Message(error_check, []byte("Failed to insert new message into database"))
			Send_Message(user, m)
		} else {
			check := Check_Command(user, m.text)
			if check == false {
				m = Create_Message(message_client_recieve, m.serverid, m.roomid, m.userid, Int_Converter(messageid), Int_Converter(Milli_Time()), m.text)
				Send_Broadcast_Server_Room(Byte_Converter(m.serverid), Byte_Converter(m.roomid), m)
			}
		}

	}
	return m
}

//Reads message code type 32, request n messages; creates multiple messages with code type 17, message client recieve,
//if the offset messages succeed to be retrieved then sends the messages to the client; creates a message with code type
//255, error check, if the offset messages fail to be retrieved then sends the message to the client.
func Offset_Message(user *User, data []byte) Message {
	m := Create_Message(request_n_messages, data[1])
	messages, err := Get_Offset_Messages(user, uint(m.offset))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve offset messages from database"))
		Send_Message(user, m)
	} else {
		for i := 0; i < len(messages); i++ {
			Send_Message(user, messages[i])
		}
	}
	return m
}

//Reads message code type 33, request all userid; creates multiple messages with code type 48, recieve userid, if the
//userid messages succeed to be retrieved then sends the messages to the client; creates a message with code type 255,
//error check, if the userid messages fail to be retrieved then sends the message to the client.
func Multi_Userid_Message(user *User, data []byte) Message {
	m := Create_Message(request_all_userid)
	messages, err := Get_Userid_Messages(user.serverid)
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve all userid from database"))
		Send_Message(user, m)
	} else {
		for i := 0; i < len(messages); i++ {
			Send_Message(user, messages[i])
		}
	}
	return m
}

//TODO
func Display_Name_Message(user *User, data []byte) Message {
	m := Create_Message(request_display_name, data[1:5])
	requested_display_name, err := Get_Display_Name(user.serverid, Byte_Converter(m.userid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve display name from database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_display_name, m.userid, requested_display_name)
		Send_Message(user, m)
	}
	return m
}

//TODO
func Multi_Serverid_Message(user *User, data []byte) Message {
	m := Create_Message(request_all_serverid, data[1:5])
	messages, err := Get_Serverid_Messages(Byte_Converter(m.userid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve all serverid from database"))
		Send_Message(user, m)
	} else {
		for i := 0; i < len(messages); i++ {
			Send_Message(user, messages[i])
		}
	}
	return m
}

//TODO
func Server_Display_Name_Message(user *User, data []byte) Message {
	m := Create_Message(request_server_display_name, data[1:5])
	requested_server_display_name, err := Get_Server_Display_Name(Byte_Converter(m.serverid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve server display name from database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_server_display_name, m.serverid, requested_server_display_name)
		Send_Message(user, m)
	}
	return m
}

//TODO
func Multi_Roomid_Message(user *User, data []byte) Message {
	m := Create_Message(request_all_roomid, data[1:5])
	messages, err := Get_Roomid_Messages(Byte_Converter(m.serverid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve all roomid from database"))
		Send_Message(user, m)
	} else {
		for i := 0; i < len(messages); i++ {
			Send_Message(user, messages[i])
		}
	}
	return m
}

//TODO
func Room_Display_Name_Message(user *User, data []byte) Message {
	m := Create_Message(request_room_display_name, data[1:5], data[5:9])
	requested_room_display_name, err := Get_Room_Display_Name(Byte_Converter(m.serverid), Byte_Converter(m.roomid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve room display name from database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_room_display_name, m.serverid, m.roomid, requested_room_display_name)
		Send_Message(user, m)
	}
	return m
}

//TODO
func Master_Display_Name_Message(user *User, data []byte) Message {
	m := Create_Message(request_master_display_name, data[1:5])
	requested_master_display_name, err := Get_Master_Display_Name(Byte_Converter(m.userid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve master display name from database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_master_display_name, m.userid, requested_master_display_name)
		Send_Message(user, m)
	}
	return m
}

//TODO
func Status_Message(user *User, data []byte) Message {
	m := Create_Message(request_status, data[1:5])
	m = Create_Message(recieve_status, m.userid, byte(user.status), byte(user.muted), byte(user.voice))
	Send_Message(user, m)
	return m
}

//TODO
func Location_Message(user *User, data []byte) Message {
	m := Create_Message(request_location, data[1:5])
	m = Create_Message(recieve_location, Int_Converter(user.serverid), Int_Converter(user.roomid), Int_Converter(user.id))
	Send_Message(user, m)
	return m
}

//TODO
func User_Ip_Message(user *User, data []byte) Message {
	m := Create_Message(request_user_ip, data[1:5])
	u := Get_User_From_Userid(Byte_Converter(m.userid))
	if u == nil{
		m = Create_Message(error_check, []byte("Failed to retrieve ip from user that is not logged in"))
		Send_Message(user, m)
	} else {
		address := u.Get_Local_Address()
		m = Create_Message(recieve_user_ip, m.userid, []byte(address))
		Send_Message(user, m)
	}
	return m
}

//TODO
func User_Icon_Message(user *User, data []byte) Message {
	m := Create_Message(request_user_icon, data[1:5])
	url, err := Get_User_Icon(Byte_Converter(m.userid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve user icon from database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_user_icon, m.userid, []byte(url))
		Send_Message(user, m)
	}
	return m
}

//TODO
func Server_Icon_Message(user *User, data []byte) Message {
	m := Create_Message(request_server_icon, data[1:5])
	url, err := Get_Server_Icon(Byte_Converter(m.serverid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve server icon from database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_server_icon, m.serverid, url)
		Send_Message(user, m)
	}
	return m
}

//TODO
func New_Display_Name_Message(user *User, data []byte) Message {
	m := Create_Message(send_new_display_name, data[1:])
	err := Set_New_Display_Name(user.serverid, user.id, m.display_name)
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to insert new display name into database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_display_name, Int_Converter(user.id), m.display_name)
		Send_Broadcast_Server(user.serverid, m)
	}
	return m
}

//TODO
func New_Server_Display_Name_Message(user *User, data []byte) Message {
	m := Create_Message(send_new_server_display_name, data[1:5], data[5:])
	err := Set_New_Server_Display_Name(Byte_Converter(m.serverid), m.server_display_name)
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to insert new server display name into database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_server_display_name, m.serverid, m.server_display_name)
		Send_Broadcast_Members(Byte_Converter(m.serverid), m)
	}
	return m
}

//TODO
func New_Room_Display_Name_Message(user *User, data []byte) Message {
	m := Create_Message(send_new_room_display_name, data[1:5], data[5:9], data[9:])
	err := Set_New_Room_Display_Name(Byte_Converter(m.serverid), Byte_Converter(m.roomid), m.room_display_name)
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to insert new room display name into database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_room_display_name, m.serverid, m.roomid, m.room_display_name)
		Send_Broadcast_Server(Byte_Converter(m.serverid), m)
	}
	return m
}

//TODO
func New_Master_Display_Name_Message(user *User, data []byte) Message {
	m := Create_Message(send_new_master_display_name, data[1:])
	err := Set_New_Master_Display_Name(user.id, m.master_display_name)
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to insert new master display name into database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_master_display_name, Int_Converter(user.id), m.master_display_name)
		Send_Broadcast_Friends(user.id, m)
	}
	return m
}

//TODO
func New_User_Icon_Message(user *User, data []byte) Message {
	m := Create_Message(send_new_user_icon, data[1:])
	err := Set_New_User_Icon(user.id, string(m.url))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to insert new user icon into database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_user_icon, Int_Converter(user.id), m.url)
		Send_Broadcast_Server(user.serverid, m)
		Send_Broadcast_Friends(user.id, m)
	}
	return m
}

//TODO
func New_Server_Icon_Message(user *User, data []byte) Message {
	m := Create_Message(send_new_server_icon, data[1:5], data[5:])
	err := Set_New_Server_Icon(Byte_Converter(m.serverid), string(m.url))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to insert new server icon into database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_server_icon, m.serverid, m.url)
		Send_Broadcast_Members(Byte_Converter(m.serverid), m)
	}
	return m
}

//TODO
func Change_Status_Message(user *User, data []byte) Message {
	m := Create_Message(change_status, data[1], data[2], data[3])
	user.status = uint(m.status)
	user.muted = uint(m.muted)
	user.voice = uint(m.voice)
	m = Create_Message(recieve_status, Int_Converter(user.id), m.status, m.muted, m.voice)
	Send_Broadcast_Server(user.serverid, m)
	Send_Broadcast_Friends(user.id, m)
	return m
}

//TODO
func Change_Location_Message(user *User, data []byte) Message {
	m := Create_Message(change_location, data[1:5], data[5:9])
	if user.serverid != Byte_Converter(m.serverid) {
		old_server := user.serverid
		user.serverid = Byte_Converter(m.serverid)
		user.roomid = Byte_Converter(m.roomid)
		m = Create_Message(recieve_location, m.serverid, m.roomid, Int_Converter(user.id))
		Send_Broadcast_Server(old_server, m)
		Send_Broadcast_Server(user.serverid, m)
	} else {
		user.roomid = Byte_Converter(m.roomid)
		m = Create_Message(recieve_location, m.serverid, m.roomid, Int_Converter(user.id))
		Send_Broadcast_Server(user.serverid, m)
	}
	return m
}

//TODO
func Create_Server_Message(user *User, data []byte) Message {
	m := Create_Message(create_server, data[1:])
	return m
}

//TOOD
func Create_Room_Message(user *User, data []byte) Message {
	m := Create_Message(create_room, data[1:])
	roomid, err := Room_Create(user.serverid, user.id, m.room_display_name)
	if err != nil {
		m = Create_Message(error_check, "Failed to create new room: user may not have permission")
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_roomid, Int_Converter(user.serverid), Int_Converter(roomid))
		Send_Broadcast_Server(user.serverid, m)
	}
	return m
}
