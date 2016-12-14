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
	voice               byte
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
		message = append(message, []byte{m.offset}...)
	}
	if m.code == change_status || m.code == recieve_status {
		message = append(message, []byte{m.status}...)
		message = append(message, []byte{m.muted}...)
		message = append(message, []byte{m.voice}...)
	}
	return message
}

//TODO
func Init_Message() {
	Messages := make([]func(args ...interface{}) Message, 256)

	Messages[create_user] = Username_Password
	Messages[create_user_fail] = Base
	Messages[login] = Username_Password
	Messages[login_fail] = Base
	Messages[login_success] = Userid
	Messages[message_client_send] = Message_Send
	Messages[message_client_recieve] = Message_Recieve
	Messages[message_edit] = Messageid_Text
	Messages[message_delete] = Messageid
	Messages[message_remove] = Messageid
	Messages[request_n_messages] = Multi_Message
	Messages[request_all_userid] = Base
	Messages[request_display_name] = Userid
	Messages[request_all_serverid] = Userid
	Messages[request_server_display_name] = Serverid
	Messages[request_all_roomid] = Serverid
	Messages[request_room_display_name] = Serverid_Roomid
	Messages[request_master_display_name] = Userid
	Messages[request_status] = Userid
	Messages[request_location] = Userid
	Messages[request_user_ip] = Userid
	Messages[request_user_icon] = Userid
	Messages[request_server_icon] = Userid
	Messages[recieve_userid] = Userid
	Messages[recieve_display_name] = Userid_Display_Name
	Messages[recieve_serverid] = Serverid_Userid
	Messages[recieve_server_display_name] = Serverid_Server_Display_Name
	Messages[recieve_roomid] = Serverid_Roomid
	Messages[recieve_room_display_name] = Serverid_Room_Display_Name
	Messages[recieve_master_display_name] = Userid_Master_Display_Name
	Messages[recieve_status] = Userid_Status
	Messages[recieve_location] = Serverid_Roomid_Userid
	Messages[recieve_user_ip] = Userid_Address
	Messages[recieve_user_icon] = Userid_Url
	Messages[recieve_server_icon] = Serverid_Url
	Messages[send_new_display_name] = Display_Name
	Messages[send_new_server_display_name] = Serverid_Server_Display_Name
	Messages[send_new_room_display_name] = Serverid_Room_Display_Name
	Messages[send_new_master_display_name] = Master_Display_Name
	Messages[send_new_user_icon] = Url
	Messages[send_new_server_icon] = Serverid_Url
	Messages[change_status] = Status
	Messages[change_location] = Serverid_Roomid
	Messages[create_server] = Server_Display_Name
	Messages[create_room] = Room_Display_Name
	Messages[error_check] = Text

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
func Serverid_Roomid_Userid(args ...interface{}) Message {
	m := Message{
		code:     args[0].(byte),
		serverid: args[1].([]byte),
		roomid:   args[2].([]byte),
		userid:   args[3].([]byte)}
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
	m := Message {
		code: args[0].(byte),
		server_display_name: args[1].([]byte)}
	return m
}

//TODO
func Serverid_Server_Display_Name(args ...interface{}) Message {
	m := Message{
		code:                args[0].(byte),
		serverid:            args[1].([]byte),
		server_display_name: args[2].([]byte)}
	return m
}

//TODO
func Room_Display_Name(args ...interface{}) Message {
	m := Message {
		code: args[0].(byte),
		room_display_name: args[1].([]byte)}
	return m
}

//TODO
func Serverid_Room_Display_Name(args ...interface{}) Message {
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
		muted:  args[2].(byte),
		voice:  args[3].(byte)}
	return m
}

//TODO
func Userid_Status(args ...interface{}) Message {
	m := Message{
		code:   args[0].(byte),
		userid: args[1].([]byte),
		status: args[2].(byte),
		muted:  args[3].(byte),
		voice:  args[4].(byte)}
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
