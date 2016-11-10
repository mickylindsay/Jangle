package main

type Message interface {
	Build_Message() []byte
}

type Base struct {
	code byte
	username []byte
	password []byte
	serverid []byte
	roomid []byte
	userid []byte
	messageid []byte
	time []byte
	offset byte
	status byte
	muted byte
	text []byte
	server_display_name []byte
	room_display_name []byte
	master_display_name []byte
	display_name []byte
	address []byte
}

func (m Base) Build_Message() []byte {
	var message []byte
	message[0] = m.code
	if (m.username != 0) {
		copy(message[1:20], m.username)
	} else if (m.password != 0) {
		copy(message[21:], m.password)
	}
}