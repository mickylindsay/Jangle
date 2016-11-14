package main

//Constructs generic Message type
type Message interface {

}

//[code:1]
type Base struct {
	code byte
}

//[code:1,username:20,password:]
type Username_Password struct {
	code byte
	username []byte
	password []byte
}

//[code:1,serverid:4]
type Serverid struct {
	code byte
	serverid []byte
}

//[code:1,roomid:4]
type Roomid struct {
	code byte
	roomid []byte
}

//[code:1,userid:4]
type Userid struct {
	code byte
	userid []byte
}

//
type Messageid struct {
	code byte
	messageid []byte
}

//
type Messageid_Text struct {
	code byte
	messageid []byte
	text []byte
}

//[code:1,serverid:4,roomid:4]
type Serverid_Roomid struct {
	code byte
	serverid []byte
	roomid []byte
}

//[code:1,serverid:4,userid:4]
type Serverid_Userid struct {
	code byte
	serverid []byte
	userid []byte
}

//[code:1,roomid:4,userid:4]
type Roomid_Userid struct {
	code byte
	roomid []byte
	userid []byte
}

//[code:1,serverid:4,server_display_name:]
type Server_Display_Name struct {
	code byte
	serverid []byte
	server_display_name []byte
}

//[code:1,serverid:4,roomid:4,room_display_name:]
type Room_Display_Name struct {
	code byte
	serverid []byte
	roomid []byte
	room_display_name []byte
}

//[code:1,new_display_name:]
type Display_Name struct {
	code byte
	display_name []byte
}

//[code:1,userid:4,display_name:]
type Userid_Display_Name struct {
	code byte
	userid []byte
	display_name []byte
}

//[code:1,serverid:4,roomid:4,userid:4,text:]
type Message_Send struct {
	code byte
	serverid []byte
	roomid []byte
	userid []byte
	text []byte
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

//[code:1,offset:1]
type Multi_Message struct {
	code byte
	offset byte
}

//[code:1,status:1]
type Status struct {
	code byte
	status byte
	muted byte
}

//[code:1,userid:4,status:1]
type Userid_Status struct {
	code byte
	userid []byte
	status byte
	muted byte
}

//
type Userid_Address struct {
	code byte
	userid []byte
	address []byte
}

//
type Url struct {
	code byte
	url []byte
}

//
type Serverid_Url struct {
	code byte
	serverid []byte
	url []byte
}

//
type Userid_Url struct {
	code byte
	userid byte
	url []byte
}

//[code:1,text:]
type Text struct {
	code byte
	text []byte
}

func Init_Message () {
	Messages := make([]func(op byte, args ...interface{}) Message, 256)

	Messages[0] = Username_Password
	Messages[1] = Base
	Messages[2] = Username_Password
	Messages[3] = Base
	Messages[4] = Userid

	Messages[16] = Message_Send
	Messages[17] = Message_Recieve
	Messages[18] = Messageid_Text
	Messages[19] = Messageid
	Messages[20] = Messageid

	Messages[32] = Multi_Message
	Messages[33] = Base
	Messages[34] = Userid
	Messages[35] = Userid
	Messages[36] = Serverid
	Messages[37] = Serverid
	Messages[38] = Serverid_Roomid
	Messages[39] = Userid
	Messages[40] = Userid
	Messages[41] = Userid
	Messages[43] = Userid
	Messages[44] = Userid

	Messages[48] = Userid
	Messages[49] = Userid_Display_Name
	Messages[50] = Serverid_Userid
	Messages[51] = Server_Display_Name
	Messages[52] = Serverid_Roomid
	Messages[53] = Room_Display_Name
	Messages[54] = Userid_Display_Name
	Messages[55] = Userid_Status
	Messages[56] = Userid_Address
	Messages[57] = Userid_Url
	Messages[58] = Serverid_Url
	
	Messages[64] = Display_Name
	Messages[65] = Server_Display_Name
	Messages[66] = Room_Display_Name
	Messages[67] = Userid_Display_Name
	Messages[68] = Url
	Messages[70] = Serverid_Url

	Messages[80] = Status
	Messages[81] = Serverid
	Messages[82] = Roomid

	Messages[255] = Text

	jangle.Messages = Messages
}

//
func Build_Message (op byte, args ...interface{}) Message {
	m := jangle.Messages[op](op, args)
	return m
}

//
func Base (op byte, args ...interface{}) Message {
	m := Base {
		code: op}
	return m
}

//
func Username_Password (op byte, args ...interface{}) Message {
	m := Username_Password {
		code: op
		username: args[0],
		password: args[1]}
	return m
}

//
func Serverid (op byte, args ...interface{}) Message {
	m := Serverid {
		code: op,
		serverid: args[0]}
	return m
}

//
func Roomid (op byte, args ...interface{}) Message {
	m := Roomid {
		code: op,
		roomid: args[0]}
	return m
}

//
func Userid (op byte, args ...interface{}) Message {
	m := Userid {
		code: op,
		userid: args[0]}
	return m
}

//
func Messageid (op byte, args ...interface{}) Message {
	m := Messageid {
		code: op,
		messageid: args[0]}
	return m
}

//
func Messageid_Text (op byte, args ...interface{}) Message {
	m := Messageid_Text {
		code: op,
		messageid: args[0],
		text: args[1]}
	return m
}

//
func Serverid_Roomid (op byte, args ...interface{}) Message {
	m := Serverid_Roomid {
		code: op,
		serverid: args[0],
		roomid: args[1]}
	return m
}

//
func Serverid_Userid (op byte, args ...interface{}) Message {
	m := Serverid_Userid {
		code: op,
		serverid: args[0],
		userid: args[1]}
	return m
}

//
func Roomid_Userid (op byte, args ...interface{}) Message {
	m := Roomid_Userid {
		code: op,
		roomid: args[0],
		userid: args[1]}
	return m
}

//
func Server_Display_Name (op byte, args ...interface{}) Message {
	m := Server_Display_Name {
		code: op,
		serverid: args[0],
		server_display_name: args[1]}
	return m
}

//
func Room_Display_Name (op byte, args ...interface{}) Message {
	m := Room_Display_Name {
		code: op,
		serverid: args[0],
		roomid: args[1],
		room_display_name: args[2]}
	return m
}

//
func Display_Name (op byte, args ...interface{}) Message {
	m := Display_Name {
		code: op,
		display_name: args[0]}
	return m
}

//
func Userid_Display_Name (op byte, args ...interface{}) Message {
	m := Userid_Display_Name {
		code: op,
		userid: args[0],
		display_name: args[1]}
	return m
}

//
func Message_Send (op byte, args ...interface{}) Message {
	m := Message_Send {
		code: op,
		serverid: args[0],
		roomid: args[1],
		userid: args[2],
		text: args[3]}
	return m
}

//
func Message_Recieve (op byte, args ...interface{}) Message {
	m := Message_Recieve {
		code: op,
		serverid: args[0],
		roomid: args[1],
		userid: args[2],
		messageid: args[3],
		time: args[4],
		text: args[5]}
	return m
}

//
func Multi_Message (op byte, args ...interface{}) Message {
	m := Multi_Message {
		code: op,
		offset: args[0]}
	return m
}

//
func Status (op byte, args ...interface{}) Message {
	m := Status {
		code: op,
		status: args[0],
		muted: args[1]}
	return m
}

//
func Userid_Status (op byte, args ...interface{}) Message {
	m := Userid_Status {
		code: op,
		userid: args[0],
		status: args[1],
		muted: args[2]}
	return m
}

//
func Url (op byte, args ...interface{}) Message {
	m := Url {
		code: op,
		url: args[0]}
	return m
}

//
func Serverid_Url (op byte, args ...interface{}) Message {
	m := Serverid_Url {
		code: op,
		serverid: args[0],
		url: args[1]}
	return m
}

//
func Urserid_Url (op byte, args ...interface{}) Message {
	m := Userid_Url {
		code: op,
		userid: args[0],
		url: args[1]}
	return m
}

//
func Text (op byte, args ...interface{}) Message {
	m := Text {
		code: op,
		text: args[0]}
	return m
}

func main() {
	return 0;
}