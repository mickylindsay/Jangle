package main

import (
	"strings"
)

//TODO
func Init_Command () {
	Commands := make([]func(args []string), 4)

	Commands[0] = Kick_User
	Commands[1] = Mute_User
	Commands[2] = Unmute_User
	Commands[3] = Move_User

	jangle.Commands = Commands
}

//TODO
func Check_Command (user *User, data []byte) bool {
	var check bool
	if (string(data[0]) == "/") {
		check = true
		args := strings.Split(string(data), " ")
		trigger := Switcher(args[0])
		jangle.Commands[trigger](args[1:])
	} else {
		check = false
	}
	return check
}

//TODO
func Switcher (s string) byte {
	var switcher byte
	switch {
		case s == "kick":
			switcher = 0
		case s == "mute":
			switcher = 1
		case s == "unmute":
			switcher = 2
		case s == "move":
			switcher = 3
	}
	return switcher
}

//TODO
func Kick_User (args []string) {
	c := Kick{
		user: Get_User_From_Userid(Byte_Converter([]byte(args[1])))}
	c.Execute()
	c.Send()
}

//TODO
func Mute_User (args []string) {
	c := Mute{
		user: Get_User_From_Userid(Byte_Converter([]byte(args[1])))}
	c.Execute()
	c.Send()
}

//TODO
func Unmute_User (args []string) {
	c := Unmute{
		user: Get_User_From_Userid(Byte_Converter([]byte(args[1])))}
	c.Execute()
	c.Send()
}

//TODO
func Move_User (args []string) {
	c := Move{
		user: Get_User_From_Userid(Byte_Converter([]byte(args[1]))),
		roomid: Byte_Converter([]byte(args[2]))}
	c.Execute()
	c.Send()
}

//TODO
type Command interface {
	Execute()
	Send()
}

//TODO
type Kick struct {
	user *User
}

//TODO
func (c Kick) Execute() {
	c.user.serverid = uint(default_value)
	c.user.roomid = uint(default_value)
	Remove_User_From_Userlist(c.user.id);
}

//TODO
func (c Kick) Send() {
	m := Serverid_Userid{
		code: broadcast_server,
		serverid: Int_Converter(c.user.serverid),
		userid: Int_Converter(c.user.id)}
	Message97(c.user, m.Build_Message())
}

//TODO
type Mute struct {
	user *User
}

//TODO
func (c Mute) Execute() {
	c.user.muted = 1
}

//TODO
func (c Mute) Send() {
	m := Userid_Status{
		code: broadcast_status,
		userid: Int_Converter(c.user.id),
		status: c.user.status,
		muted: c.user.muted}
	Message96(c.user, m.Build_Message())
}

//TODO
type Unmute struct {
	user *User
}

//TODO
func (c Unmute) Execute() {
	c.user.muted = 2
}

//TODO
func (c Unmute) Send() {
	m := Userid_Status{
		code: broadcast_status,
		userid: Int_Converter(c.user.id),
		status: c.user.status,
		muted: c.user.muted}
	Message96(c.user, m.Build_Message())
}

//TODO
type Move struct {
	user *User
	roomid uint
}

//TODO
func (c Move) Execute() {
	c.user.roomid = c.roomid
}

//TODO
func (c Move) Send() {
	m := Roomid_Userid{
		code: broadcast_room,
		roomid: Int_Converter(c.user.roomid),
		userid: Int_Converter(c.user.id)}
	Message98(c.user, m.Build_Message())
}