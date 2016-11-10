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
 func (m Base) Build_Message() []byte {
	message := make([]byte, 1)
	message[0] = m.code
	return message
}

//[code:1,username:20,password:]
type Username_Password struct {
	code byte
	username []byte
	password []byte
}

 func (m Username_Password) Build_Message() []byte {
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

 func (m Userid) Build_Message() []byte {
	message := make([]byte, 5)
	message[0] = m.code
	copy(message[1:4], m.userid[:])
	return message
}

//[code:1,serverid:4]
type Serverid struct {
	code byte
	serverid []byte
}

 func (m Serverid) Build_Message() []byte {
	message := make([]byte, 5)
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	return message
}

//[code:1,serverid:4,userid:4]
type Serverid_Userid struct {
	code byte
	serverid []byte
	userid []byte
}

 func (m Serverid_Userid) Build_Message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.userid[:])
	return message
}

//[code:1,roomid:4]
type Roomid struct {
	code byte
	roomid []byte
}

 func (m Roomid) Build_Message() []byte {
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

 func (m Roomid_Userid) Build_Message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.roomid[:])
	copy(message[5:8], m.userid[:])
	return message
}

//[code:1,serverid:4,roomid:4]
type Serverid_Roomid struct {
	code byte
	serverid []byte
	roomid []byte
}

 func (m Serverid_Roomid) Build_Message() []byte {
	message := make([]byte, 9)
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
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

 func (m Message_Send) Build_Message() []byte {
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

 func (m Message_Recieve) Build_Message() []byte {
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

 func (m Multi_Message) Build_Message() []byte {
	message := make([]byte, 2)
	message[0] = m.code
	message[1] = m.offset
	return message
}

//[code:1,userid:4,display_name:]
type Display_Name struct {
	code byte
	userid []byte
	display_name []byte
}

 func (m Display_Name) Build_Message() []byte {
	message := make([]byte, 5 + len(m.display_name))
	message[0] = m.code
	copy(message[1:4], m.userid[:])
	copy(message[5:], m.display_name[:])
	return message
}

//[code:1,serverid:4,server_display_name:]
type Server_Display_Name struct {
	code byte
	serverid []byte
	server_display_name []byte
}

 func (m Server_Display_Name) Build_Message() []byte {
	message := make([]byte, 5 + len(m.server_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:], m.server_display_name[:])
	return message
}

//[code:1,serverid:4,roomid:4,room_display_name:]
type Room_Display_Name struct {
	code byte
	serverid []byte
	roomid []byte
	room_display_name []byte
}

 func (m Room_Display_Name) Build_Message() []byte {
	message := make([]byte, 9 + len(m.room_display_name))
	message[0] = m.code
	copy(message[1:4], m.serverid[:])
	copy(message[5:8], m.roomid[:])
	copy(message[9:], m.room_display_name[:])
	return message
}

//[code:1,new_display_name:]
type New_Display_Name struct {
	code byte
	new_display_name []byte
}

 func (m New_Display_Name) Build_Message() []byte {
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

 func (m New_Server_Display_Name) Build_Message() []byte {
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

 func (m New_Room_Display_Name) Build_Message() []byte {
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
	muted byte
}

 func (m Status) Build_Message() []byte {
	message := make([]byte, 3)
	message[0] = m.code
	message[1] = m.status
	message[2] = m.muted
	return message
}

//[code:1,userid:4,status:1]
type Userid_Status struct {
	code byte
	userid []byte
	status byte
	muted byte
}

 func (m Userid_Status) Build_Message() []byte {
	message := make([]byte, 6)
	message[0] = m.code
	copy(message[1:4], m.userid[:])
	message[5] = m.status
	message[6] = m.muted
	return message
}

//[code:1,text:]
type Text struct {
	code byte
	text []byte
}

 func (m Text) Build_Message() []byte {
	message := make([]byte, 1 + len(m.text))
	message[0] = m.code
	copy(message[1:], m.text[:])
	return message
}