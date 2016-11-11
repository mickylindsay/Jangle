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

func Init_Message () {
	Messages := make([]func(args ...interface{}) Message, 256)

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

func Build_Message(code byte, args ...interface{}) Message {
	jangle.Messages[code](args)
}

