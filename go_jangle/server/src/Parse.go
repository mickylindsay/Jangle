package main

//Initializes function array that contains all the functions necessary to handle every
//message code
func Init_Parse () {
	Messages := make([]func(user *User, data []byte) Message, 256)

	Messages[0] = Message0
	Messages[1] = Message1
	Messages[2] = Message2
	Messages[3] = Message3
	Messages[4] = Message4

	Messages[16] = Message16
	Messages[17] = Message17

	Messages[32] = Message32
	Messages[33] = Message33
	Messages[34] = Message34
	Messages[35] = Message35
	Messages[36] = Message36
	Messages[37] = Message37
	Messages[38] = Message38
	Messages[39] = Message39
	Messages[40] = Message40
	Messages[41] = Message41

	Messages[48] = Message48
	Messages[49] = Message49
	Messages[50] = Message50
	Messages[51] = Message51
	Messages[52] = Message52
	Messages[53] = Message53
	Messages[54] = Message54
	Messages[55] = Message55
	Messages[56] = Message56
	
	Messages[64] = Message64
	Messages[65] = Message65
	Messages[66] = Message66
	Messages[67] = Message67

	Messages[80] = Message80
	Messages[81] = Message81
	Messages[82] = Message82

	Messages[96] = Message96
	Messages[97] = Message97
	Messages[98] = Message98
	Messages[99] = Message99
	Messages[100] = Message100
	Messages[101] = Message101
	Messages[102] = Message102

	Messages[255] = Message255

	jangle.Messages = Messages
}

//Master function: takes paramaters type User struct and byte array
//byte array is the message that is recieved from the client
//the type User struct is a reference to the connection that represents
//the client side user that is associated with the byte array message
//this function determines what type of message is being recieved
//and calls the appropriate function based off the code type
func Parse_Data (user *User, data []byte) {

	jangle.Messages[data[0]](user, data)
}

//If create user successful, convert message to code type 4
//If create user fail, convert message to code type 1
func Message0 (user *User, data []byte) Message {

	m := Username_Password{
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

			return m
}

//Writes to user message code type 1, create user fail
func Message1 (user *User, data []byte) Message {

	m := Base{
		code: data[0]}

			user.Write(m.Build_Message())
			return m
}

//If login successful, convert message to code type 4
//If login fail, convert message to code type 3
func Message2 (user *User, data []byte) Message {

	m := Username_Password{
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

			return m
}

//Writes to user message code type 3, login fail
func Message3 (user *User, data []byte) Message {

	m := Base{
		code: data[0]}

			user.Write(m.Build_Message())
			return m
}

//Send message code type 4 to client, login success
func Message4 (user *User, data []byte) Message {

	m := Userid{
		code: data[0],
		userid: data[1:4]}

			Send_Message(user, m)
			return m
}

//Converts message code type 16, message client send
//to message code type 17, message client recieve
func Message16 (user *User, data []byte) Message {

	m := Message_Send{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		userid: data[9:12],
		text: data[13:]}
		
			if (user.muted == user_unmuted) {
				Check_Command(user, m.text);
				err := Message_Create(user, data[13:])
				Check_Error(err)

				check := Check_Command(user, m.text)

				if (check == false) {
					data[0] = message_client_recieve
					data = Time_Stamp(data)
					Message17(user, data)
				}
			}

			return m		
}

//Sends message code type 17, message client recieve,
//to a chat room on a specific server
func Message17 (user *User, data []byte) Message {

	m := Message_Recieve{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		userid: data[9:12],
		time: data[13:16],
		text: data[17:]}

			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			Send_Broadcast_Server_Room(num1, num2, m)
			return m
}

//Requests n message code type 17's, message client recieve, from database 
//dependent on offset value from message code type 32, request n messages
func Message32 (user *User, data []byte) Message {

	m := Multi_Message{
		code: data[0],
		offset: data[1]}

			num := uint(data[1])
			messages, err := Request_Offset_Messages(user, num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Message17(user, messages[i].Build_Message())
			}

			return m
}

//Requests message code type 48's, recieve userid, from database
//which consists of all userids conected to a specific server
func Message33 (user *User, data []byte) Message {

	m := Base{
		code: data[0]}

			messages, err := Request_Userid_Messages(user.serverid)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Message48(user, messages[i].Build_Message())
			}

			return m
}

//Requests a display name from the database from a user connected to a specific server
//from requested userid in message code type 34, requested display name, then
//builds a new byte array in the format of message code type 49, recieve display name
func Message34 (user *User, data []byte) Message {

	m := Userid{
		code: data[0],
		userid: data[1:4]}

			num := Byte_Converter(data[1:4])
			requested_display_name, err := Request_Display_Name(user.serverid, num)
			Check_Error(err)

			data = make([]byte, len(requested_display_name) + 5)
			data[0] = recieve_display_name
			copy(data[1:4], Int_Converter(num))
			copy(data[5:], requested_display_name)

			Message49(user, data)
			return m				
}

//Requests message code type 50's, recieve serverid, which consists of all the serverids
//that a specific user is connected to from message code type 35, request all serverid
func Message35 (user *User, data []byte) Message {

	m := Userid{
		code: data[0],
		userid: data[1:4]}

			num := Byte_Converter(data[1:4])
			messages, err := Request_Serverid_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Message50(user, messages[i].Build_Message())
			}

			return m
}

//Requests server display name from a specific server from the requested servid in
//message code tye 36, request server display name, then builds
//new byte array in the format of message code type 51, recieve server display name
func Message36 (user *User, data []byte) Message {

	m := Serverid{
		code: data[0],
		serverid: data[1:4]}

			num := Byte_Converter(data[1:4])
			requested_server_display_name, err := Request_Server_Display_Name(num)
			Check_Error(err)

			data = make([]byte, len(requested_server_display_name) + 5)
			data[0] = recieve_server_display_name
			copy(data[1:4], Int_Converter(num))
			copy(data[5:], requested_server_display_name)

			Message51(user, data)
			return m
}

//Request message code type 52's, recieve roomid, which consist of all the roomids
//on a specific server from the serverid in message code type 37, request all roomid
func Message37 (user *User, data []byte) Message {

	m := Serverid{
		code: data[0],
		serverid: data[1:4]}

			num := Byte_Converter(data[1:4])
			messages, err := Request_Roomid_Messages(num)
			Check_Error(err)

			for i := 0; i < len(messages); i++ {
				Message52(user, messages[i].Build_Message())
			}

			return m
}

//Requests room display name from a specific room on specific server from the requested
//serverid and roomid in message code type 38, request room display name, then builds
//a new byte array in the format of message code type 53, recieve room display name
func Message38 (user *User, data []byte) Message {

	m := Serverid_Roomid{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8]}

			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			requested_room_display_name, err := Request_Room_Display_Name(num1, num2)
			Check_Error(err)

			data = make([]byte, len(requested_room_display_name) + 9)
			data[0] = recieve_room_display_name
			copy(data[1:4], Int_Converter(num1))
			copy(data[5:8], Int_Converter(num2))
			copy(data[9:], requested_room_display_name)

			Message53(user, data)
			return m
}

//Requests master display name from a specific user from the requested userid in 
//message code type 39, request master diplay name, then builds a new byte array 
//in the format of message code type 54, recieve master display name
func Message39 (user *User, data []byte) Message {

	m := Userid{
		code: data[0],
		userid: data[1:4]}

			num := Byte_Converter(data[1:4])
			requested_master_display_name, err := Request_Master_Display_Name(num)
			Check_Error(err)

			data = make([]byte, len(requested_master_display_name) + 5)
			data[0] = recieve_master_display_name
			copy(data[1:4], Int_Converter(num))
			copy(data[5:], requested_master_display_name)

			Message54(user, data)
			return m
}

//Requests the status from a specific user from the requested userid in message code
//type 40, request status, then builds a new byte array in the format of message code
//type 55, recieve status
func Message40 (user *User, data []byte) Message {

	m := Userid{
		code: data[0],
		userid: data[1:4]}

			arr := data[1:4]
			data = make([]byte, 6)
			data[0] = recieve_status
			copy(data[1:4], arr)
			data[5] = user.status
			data[6] = user.muted
		
			Message55(user, data)
			return m
}

//Requests the local ip address of any user via userid, used for client side voice communication
func Message41 (user *User, data []byte) Message {

	m := Userid{
		code: data[0],
		userid: data[1:4]}

			num := Byte_Converter(data[1:4])
			addr := Get_User_From_Userid(num).Get_Local_Address();
			arr := data[1:4]
			data = make([]byte, 5 + len(addr))
			data[0] = recieve_status
			copy(data[1:4], arr)
			copy(data[5:], addr)
		
			Message56(user, data)
			return m
}

//Sends message code type 48, recieve userid, to client
func Message48 (user *User, data []byte) Message {

	m := Userid{
		code: data[0],
		userid: data[1:4]}

			Send_Message(user, m)
			return m
}

//Sends message code type 49, recieve display name, to client
func Message49 (user *User, data []byte) Message {

	m := Display_Name{
		code: data[0],
		userid: data[1:4],
		display_name: data[5:]}

			Send_Message(user, m)
			return m
}

//Sends message code type 50, recieve serverid, to client
func Message50 (user *User, data []byte) Message {

	m := Serverid_Userid{
		code: data[0],
		serverid: data[1:4],
		userid: data[5:8]}

			Send_Message(user, m)
			return m
}

//Sends message code type 51, recieve server display name, to client
func Message51 (user *User, data []byte) Message {

	m := Server_Display_Name{
		code: data[0],
		serverid: data[1:4],
		server_display_name: data[5:]}

			Send_Message(user, m)
			return m
}

//Sends message code type 52, recieve roomid, to client
func Message52 (user *User, data []byte) Message {

	m := Serverid_Roomid{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8]}

			Send_Message(user, m)
			return m
}

//Sends message code type 53, recieve room display name, to client
func Message53 (user *User, data []byte) Message {

	m := Room_Display_Name{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		room_display_name: data[9:]}

			Send_Message(user, m)
			return m
}

//Sends message code type 54, recieve master display name, to client
func Message54 (user *User, data []byte) Message {

	m := Display_Name{
		code: data[0],
		userid: data[1:4],
		display_name: data[5:]}

			Send_Message(user, m)
			return m
}

//Sends message code type 55, recieve status, to client
func Message55 (user *User, data []byte) Message {

	m := Userid_Status{
		code: data[0],
		userid: data[1:4],
		status: data[5],
		muted: data[6]}

			Send_Message(user, m)
			return m
}

//Broadcasts to a single user the ip of a requested user
func Message56 (user *User, data []byte) Message {

	m := Display_Name{
		code: data[0],
		userid: data[1:4],
		display_name: data[5:]}

			Send_Message(user, m)
			return m
}

//Replaces the user's display name with the new display name in message code type 64,
//send new display name, then sends message code type 99, broadcast display name, to
//all users on the user's connected server
func Message64 (user *User, data []byte) Message {

	m := New_Display_Name{
		code: data[0],
		new_display_name: data[1:]}

			err := Set_New_Display_Name(user.serverid, user.id, data[1:])
			Check_Error(err)

			data = make([]byte, len(m.new_display_name) + 5)
			data[0] = broadcast_display_name
			copy(data[1:4], Int_Converter(user.id))
			copy(data[5:], m.new_display_name)

			Message99(user, data)
			return m
}

//Replaces the server's display name with the new server display name in message
//code type 65, send new  server display name, then sends message code type 100,
//broadcast server display name, to all users that are members of the server
func Message65 (user *User, data []byte) Message {

	m := New_Server_Display_Name{
		code: data[0],
		serverid: data[1:4],
		new_server_display_name: data[5:]}

			num := Byte_Converter(data[1:4])
			err := Set_New_Server_Display_Name(num, data[5:])
			Check_Error(err)

			data = make([]byte, len(m.new_server_display_name) + 5)
			data[0] = broadcast_server_display_name
			copy(data[1:4], m.serverid)
			copy(data[5:], m.new_server_display_name)

			Message100(user, data)
			return m
}

//Replaces the room's display name with the new room display name in message
//code type 66, send new room display name, then sends message code type 101,
//broadcast room display name, to all users on that are connected to the server
func Message66 (user *User, data []byte) Message {

	m := New_Room_Display_Name{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		new_room_display_name: data[9:]}

			num1 := Byte_Converter(data[1:4])
			num2 := Byte_Converter(data[5:8])
			err := Set_New_Room_Display_Name(num1, num2, data[9:])
			Check_Error(err)

			data = make([]byte, len(m.new_room_display_name) + 9)
			data[0] = broadcast_room_display_name
			copy(data[1:4], m.serverid)
			copy(data[5:8], m.roomid)
			copy(data[9:], m.new_room_display_name)

			Message101(user, data)
			return m	
}

//Replaces the user's master display name with the new master display name in message
//code type 67, send new master display name, then sends message code type 102,
//broadcast master display name, to all of the user's friends
func Message67 (user *User, data []byte) Message {
	m := New_Display_Name{
		code: data[0],
		new_display_name: data[1:]}

			err := Set_New_Master_Display_Name(user.id, data[1:])
			Check_Error(err)

			data = make([]byte, len(m.new_display_name) + 5)
			data[0] = broadcast_master_display_name
			copy(data[1:4], Int_Converter(user.id))
			copy(data[5:], m.new_display_name)

			Message102(user, data)
			return m
}

//Changes the user's status to the new status in message code type 80, change status,
//then builds a new byte array in the format of message code type 96, broadcast status
func Message80 (user *User, data []byte) Message {

	m := Status{
		code: data[0],
		status: data[1],
		muted: data[2]}

			user.status = data[1]
			user.muted = data[2]

			data = make([]byte, 6)
			data[0] = broadcast_status
			copy(data[1:4], Int_Converter(user.id))
			data[5] = user.status
			data[6] = user.muted

			Message96(user, data)
			return m
}

//Changes the user's server to the new server in message code type 81, change server,
//then builds a new byte array in the format of message code type 97, broadcast server
func Message81 (user *User, data []byte) Message {

	m := Serverid{
		code: data[0],
		serverid: data[1:4]}

			user.serverid = Byte_Converter(data[1:4])

			data = make([]byte, 9)
			data[0] = broadcast_server
			copy(data[1:4], Int_Converter(user.serverid))
			copy(data[5:8], Int_Converter(user.id))

			Message97(user, data)
			return m
}

//Changes the user's room to the new room in message code type 82, change room, then
//builds a new byte array in the format of message code type 98, broadcast room
func Message82 (user *User, data []byte)  Message {

	m := Roomid{
		code: data[0],
		roomid: data[1:4]}

			user.roomid = Byte_Converter(data[1:4])
			
			data = make([]byte, 9)
			data[0] = broadcast_room
			copy(data[1:4], Int_Converter(user.roomid))
			copy(data[5:8], Int_Converter(user.id))		

			Message98(user, data)
			return m
}

//Sends message code type 96, broadcast status, to all users on the user's connected
//server
func Message96 (user *User, data []byte) Message {

	m := Userid_Status{
		code: data[0],
		userid: data[1:4],
		status: data[5],
		muted: data[6]}

			Send_Broadcast_Server(user.serverid, m)
			return m
}

//Sends message code type 97, broadcast server, to all users on the specified server
func Message97 (user *User, data []byte) Message {

	m := Serverid_Userid{
		code: data[0],
		serverid: data[1:4],
		userid: data[5:8]}

			Send_Broadcast_Server(Byte_Converter(m.serverid), m)
			return m
}

//Sends message code type 98, broadcast room, to all users on the user's connected 
//server
func Message98 (user *User, data []byte) Message {

	m := Roomid_Userid{
		code: data[0],
		roomid: data[1:4],
		userid: data[5:8]}

			Send_Broadcast_Server(user.serverid, m)
			return m
}

//Sends message code type 99, broadcast display name, to all users on the user's
//connected server
func Message99 (user *User, data []byte) Message {

	m := Display_Name{
		code: data[0],
		userid: data[1:4],
		display_name: data[5:]}

			Send_Broadcast_Server(user.serverid, m)
			return m
}

//Sends message code type 100, broadcast server display name, to all users that are
//members of the specified server
func Message100 (user *User, data []byte) Message {

	m := Server_Display_Name{
		code: recieve_server_display_name,
		serverid: data[1:4],
		server_display_name: data[5:]}

			Send_Broadcast_Members(Byte_Converter(m.serverid), m)
			return m
}

//Sends message code type 101, broadcast room display name, to all users on the
//specified server
func Message101 (user *User, data []byte) Message {

	m := Room_Display_Name{
		code: data[0],
		serverid: data[1:4],
		roomid: data[5:8],
		room_display_name: data[9:]}

			Send_Broadcast_Server(Byte_Converter(m.serverid), m)
			return m
}

//Sends message code type 102, broadcast master display name, to all users that are
//the user's friends
func Message102 (user *User, data []byte) Message {

	m := Display_Name{
		code: data[0],
		userid: data[1:4],
		display_name: data[5:]}

			Send_Broadcast_Friends(Byte_Converter(m.userid), m)
			return m
}

//Sends message code type 255, error check, to client
func Message255 (user *User, data []byte) Message {
	
	m := Text{
		code: data[0],
		text: data[1:]}

			Send_Message(user, m)
			return m
}