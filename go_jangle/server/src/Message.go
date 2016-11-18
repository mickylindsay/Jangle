package main

//TODO
type Message struct {
	code                byte
	username            []byte
	password            []byte
	serverid            []byte
	roomid              []byte
	userid              []byte
	messageid           []byte
	time                []byte
	server_display_name []byte
	room_display_name   []byte
	master_display_name []byte
	display_name        []byte
	text                []byte
	address             []byte
	url                 []byte
	offset              byte
	status              byte
	muted               byte
}

//TODO
func Create_Message(args ...interface{}) Message {
	m := jangle.Messages[args[0].(byte)](args...)
	return m
}

//TODO
func (m Message) Build_Message() []byte {
	var message []byte = []byte{m.code}
	if m.username != nil {
		message = append(message, m.username...)
	}
	if m.password != nil {
		message = append(message, m.password...)
	}
	if m.serverid != nil {
		message = append(message, m.serverid...)
	}
	if m.roomid != nil {
		message = append(message, m.roomid...)
	}
	if m.userid != nil {
		message = append(message, m.userid...)
	}
	if m.messageid != nil {
		message = append(message, m.messageid...)
	}
	if m.time != nil {
		message = append(message, m.time...)
	}
	if m.server_display_name != nil {
		message = append(message, m.server_display_name...)
	}
	if m.room_display_name != nil {
		message = append(message, m.room_display_name...)
	}
	if m.master_display_name != nil {
		message = append(message, m.master_display_name...)
	}
	if m.display_name != nil {
		message = append(message, m.display_name...)
	}
	if m.text != nil {
		message = append(message, m.text...)
	}
	if m.address != nil {
		message = append(message, m.address...)
	}
	if m.url != nil {
		message = append(message, m.url...)
	}
	if m.offset != 0 {
		var temp1 []byte = []byte{m.offset}
		message = append(message, temp1...)
	}
	if m.code == (change_status | recieve_status) {
		var temp2 []byte = []byte{m.status}
		message = append(message, temp2...)
		var temp3 []byte = []byte{m.muted}
		message = append(message, temp3...)
	}
	return message
}

//TODO
func Init_Message() {
	Messages := make([]func(args ...interface{}) Message, 256)

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
	Messages[54] = Userid_Master_Display_Name
	Messages[55] = Userid_Status
	Messages[56] = Userid_Address
	Messages[57] = Userid_Url
	Messages[58] = Serverid_Url

	Messages[64] = Display_Name
	Messages[65] = Server_Display_Name
	Messages[66] = Room_Display_Name
	Messages[67] = Master_Display_Name
	Messages[68] = Url
	Messages[70] = Serverid_Url

	Messages[80] = Status
	Messages[81] = Serverid
	Messages[82] = Roomid

	Messages[96] = Userid_Status
	Messages[97] = Serverid_Userid
	Messages[98] = Roomid_Userid
	Messages[99] = Userid_Display_Name
	Messages[100] = Server_Display_Name
	Messages[101] = Room_Display_Name
	Messages[102] = Userid_Master_Display_Name

	Messages[255] = Text

	jangle.Messages = Messages
}

//TODO
func Base(args ...interface{}) Message {
	m := Message{
		code: args[0].(byte)}
	return m
}

//TODO
func Username_Password(args ...interface{}) Message {
	m := Message{
		code:     args[0].(byte),
		username: args[1].([]byte),
		password: args[2].([]byte)}
	return m
}

//TODO
func Serverid(args ...interface{}) Message {
	m := Message{
		code:     args[0].(byte),
		serverid: args[1].([]byte)}
	return m
}

//TODO
func Roomid(args ...interface{}) Message {
	m := Message{
		code:   args[0].(byte),
		roomid: args[1].([]byte)}
	return m
}

//TODO
func Userid(args ...interface{}) Message {
	m := Message{
		code:   args[0].(byte),
		userid: args[1].([]byte)}
	return m
}

//TODO
func Messageid(args ...interface{}) Message {
	m := Message{
		code:      args[0].(byte),
		messageid: args[1].([]byte)}
	return m
}

//TODO
func Messageid_Text(args ...interface{}) Message {
	m := Message{
		code:      args[0].(byte),
		messageid: args[1].([]byte),
		text:      args[2].([]byte)}
	return m
}

//TODO
func Serverid_Roomid(args ...interface{}) Message {
	m := Message{
		code:     args[0].(byte),
		serverid: args[1].([]byte),
		roomid:   args[2].([]byte)}
	return m
}

//TODO
func Serverid_Userid(args ...interface{}) Message {
	m := Message{
		code:     args[0].(byte),
		serverid: args[1].([]byte),
		userid:   args[2].([]byte)}
	return m
}

//TODO
func Roomid_Userid(args ...interface{}) Message {
	m := Message{
		code:   args[0].(byte),
		roomid: args[1].([]byte),
		userid: args[2].([]byte)}
	return m
}

//TODO
func Server_Display_Name(args ...interface{}) Message {
	m := Message{
		code:                args[0].(byte),
		serverid:            args[1].([]byte),
		server_display_name: args[2].([]byte)}
	return m
}

//TODO
func Room_Display_Name(args ...interface{}) Message {
	m := Message{
		code:              args[0].(byte),
		serverid:          args[1].([]byte),
		roomid:            args[2].([]byte),
		room_display_name: args[3].([]byte)}
	return m
}

//TODO
func Master_Display_Name(args ...interface{}) Message {
	m := Message{
		code:                args[0].(byte),
		master_display_name: args[1].([]byte)}
	return m
}

//TODO
func Userid_Master_Display_Name(args ...interface{}) Message {
	m := Message{
		code:                args[0].(byte),
		userid:              args[1].([]byte),
		master_display_name: args[2].([]byte)}
	return m
}

//TODO
func Display_Name(args ...interface{}) Message {
	m := Message{
		code:         args[0].(byte),
		display_name: args[1].([]byte)}
	return m
}

//TODO
func Userid_Display_Name(args ...interface{}) Message {
	m := Message{
		code:         args[0].(byte),
		userid:       args[1].([]byte),
		display_name: args[2].([]byte)}
	return m
}

//TODO
func Message_Send(args ...interface{}) Message {
	m := Message{
		code:     args[0].(byte),
		serverid: args[1].([]byte),
		roomid:   args[2].([]byte),
		userid:   args[3].([]byte),
		text:     args[4].([]byte)}
	return m
}

//TODO
func Message_Recieve(args ...interface{}) Message {
	m := Message{
		code:      args[0].(byte),
		serverid:  args[1].([]byte),
		roomid:    args[2].([]byte),
		userid:    args[3].([]byte),
		messageid: args[4].([]byte),
		time:      args[5].([]byte),
		text:      args[6].([]byte)}
	return m
}

//TODO
func Multi_Message(args ...interface{}) Message {
	m := Message{
		code:   args[0].(byte),
		offset: args[1].(byte)}
	return m
}

//TODO
func Status(args ...interface{}) Message {
	m := Message{
		code:   args[0].(byte),
		status: args[1].(byte),
		muted:  args[2].(byte)}
	return m
}

//TODO
func Userid_Status(args ...interface{}) Message {
	m := Message{
		code:   args[0].(byte),
		userid: args[1].([]byte),
		status: args[2].(byte),
		muted:  args[3].(byte)}
	return m
}

//TODO
func Userid_Address(args ...interface{}) Message {
	m := Message{
		code:    args[0].(byte),
		userid:  args[1].([]byte),
		address: args[2].([]byte)}
	return m
}

//TODO
func Url(args ...interface{}) Message {
	m := Message{
		code: args[0].(byte),
		url:  args[1].([]byte)}
	return m
}

//TODO
func Serverid_Url(args ...interface{}) Message {
	m := Message{
		code:     args[0].(byte),
		serverid: args[1].([]byte),
		url:      args[2].([]byte)}
	return m
}

//TODO
func Userid_Url(args ...interface{}) Message {
	m := Message{
		code:   args[0].(byte),
		userid: args[1].([]byte),
		url:    args[2].([]byte)}
	return m
}

//TODO
func Text(args ...interface{}) Message {
	m := Message{
		code: args[0].(byte),
		text: args[1].([]byte)}
	return m
}