package main

//Master function: takes paramaters type User struct and byte array
//byte array is the message that is recieved from the client
//the type User struct is a reference to the connection that represents
//the client side user that is associated with the byte array message
//this function determines what type of message is being recieved
//and calls the appropriate function based off the code type
func Parse_Data(user *User, data []byte) Message {
	m := jangle.Parsers[data[0]](user, data)
	return m
}

//Initializes function array that contains all the functions necessary to handle every
//message code
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

	jangle.Parsers = Parsers
}

//TODO
func Create_User_Message(user *User, data []byte) Message {
	m := Create_Message(create_user, data[1:20], data[21:])
	id, err := User_Create(m.username, m.password)
	if err == nil {
		user.id = id
		m = Create_Message(login_success, Int_Converter(id))
		Send_Message(user, m)
	} else {
		m = Create_Message(create_user_fail)
		user.Write(m.Build_Message())
	}
	return m
}

//TODO
func Login_Message(user *User, data []byte) Message {
	m := Create_Message(login, data[1:20], data[21:])
	id, err := User_Login(m.username, m.password)
	if err == nil {
		user.id = id
		Create_Message(login_success, Int_Converter(id))
		Send_Message(user, m)
	} else {
		Create_Message(login_fail)
		user.Write(m.Build_Message())
	}
	return m
}

//TODO
func Standard_Message(user *User, data []byte) Message {
	m := Create_Message(message_client_send, data[1:5], data[6:10], data[11:15], data[16:])
	if user.muted != 1 {
		messageid, err := Message_Create(user, m.text)
		if err != nil {
			m = Create_Message(error_check, []byte("Failed to insert message into database"))
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

//TODO
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

//TODO
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

//I HAVE NO FUCKING IDEA
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
	m := Create_Message(request_all_serverid, data[1:4])
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
	m := Create_Message(request_server_display_name, data[1:4])
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
	m := Create_Message(request_all_roomid, data[1:4])
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
	m := Create_Message(request_room_display_name, data[1:4], data[5:8])
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
	m := Create_Message(request_master_display_name, data[1:4])
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
	m := Create_Message(request_status, data[1:4])
	m = Create_Message(recieve_status, m.userid, byte(user.status), byte(user.muted), byte(user.voice))
	Send_Message(user, m)
	return m
}

//TODO
func Location_Message(user *User, data []byte) Message {
	m := Create_Message(request_location, data[1:4])
	m = Create_Message(recieve_location, Int_Converter(user.serverid), Int_Converter(user.roomid), Int_Converter(user.id))
	Send_Message(user, m)
	return m
}

//TODO
func User_Ip_Message(user *User, data []byte) Message {
	m := Create_Message(request_user_ip, data[1:4])
	address := Get_User_From_Userid(Byte_Converter(m.userid)).Get_Local_Address()
	m = Create_Message(recieve_user_ip, m.userid, String_Converter(address))
	Send_Message(user, m)
	return m
}

//TODO
func User_Icon_Message(user *User, data []byte) Message {
	m := Create_Message(request_user_icon, data[1:4])
	url, err := Get_User_Icon(Byte_Converter(m.userid))
	if err != nil {
		m = Create_Message(error_check, []byte("Failed to retrieve user icon from database"))
		Send_Message(user, m)
	} else {
		m = Create_Message(recieve_user_icon, m.userid, String_Converter(url))
		Send_Message(user, m)
	}
	return m
}

//TODO
func Server_Icon_Message(user *User, data []byte) Message {
	m := Create_Message(request_server_icon, data[1:4])
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
	m := Create_Message(send_new_server_display_name, data[1:4], data[5:])
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
	m := Create_Message(send_new_room_display_name, data[1:4], data[5:8], data[9:])
	err := Set_New_Room_Display_Name(Byte_Converter(m.serverid), Byte_Converter(m.roomid), data[9:])
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
	m := Create_Message(send_new_server_icon, data[1:4], data[5:])
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
	m := Create_Message(change_location, data[1:4], data[5:8])
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
