package main

import (
	"strings"
)

//Initializes function array that contains all the functions necessary to handle all
//the bot commands
func Init_Command () {
	Commands := make([]func(args []string), 5)

	Commands[0] = Kick_User
	Commands[1] = Mute_User
	Commands[2] = Unmute_User
	Commands[3] = Move_User
	Commands[4] = Prune_N_Messages

	jangle.Commands = Commands
}

//Checks if the text part of the byte array from a message code type 16, message 
//client send, is a bot command
//If the text is bot command, the appropiate function is called from the function array
//that will execute the bot command properly
//If the text is not a bot command, the message code type 16 is handled normally
func Check_Command (user *User, data []byte) bool {
	var check bool
	if (string(data[0]) == "/") {
		args := strings.Split(string(data), " ")
		trigger := Switcher(args[0])
		if (trigger != 255) {
			check = true
			jangle.Commands[trigger](args[1:])
		} else {
			check = false
		}	
	} else {
		check = false
	}
	return check
}

//Changes the trigger word from the bot command to its corresponding byte value to
//reference the funciton array
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
		case s == "prune":
			switcher = 4
		default:
			switcher = 255
	}
	return switcher
}

//Kicks the user from the server
func Kick_User (args []string) {
	c := Kick{
		user: Get_User_From_Userid(Byte_Converter([]byte(args[1])))}
	c.Execute()
	c.Send()
}

//Mutes the user
func Mute_User (args []string) {
	c := Mute{
		user: Get_User_From_Userid(Byte_Converter([]byte(args[1])))}
	c.Execute()
	c.Send()
}

//Unmutes the user
func Unmute_User (args []string) {
	c := Unmute{
		user: Get_User_From_Userid(Byte_Converter([]byte(args[1])))}
	c.Execute()
	c.Send()
}

//Moves the user to a different room
func Move_User (args []string) {
	c := Move{
		user: Get_User_From_Userid(Byte_Converter([]byte(args[1]))),
		roomid: Byte_Converter([]byte(args[2]))}
	c.Execute()
	c.Send()
}

//TODO
func Prune_N_Messages (args []string) {
	var u *User
	if (len(args) > 2) {
		u = Get_User_From_Userid(Byte_Converter([]byte(args[2])))
	} else {
		u = nil
	}
	c := Prune{
		user: u,
		num_messages: Byte_Converter([]byte(args[1]))}
	c.Execute()
	c.Send()
}

//Constructs generic Command type
type Command interface {
	Execute()
	Send()
}

//Creates Kick struct with param User type
type Kick struct {
	user *User
}

//Sets the user's severid and roomid to the default value and removes the user from
//the user list
func (c Kick) Execute() {
	c.user.serverid = Byte_Converter(default_value)
	c.user.roomid = Byte_Converter(default_value)
	Remove_User_From_Userlist(c.user.id);
}

//Builds a message code type 97, broadcast server
func (c Kick) Send() {
	m := Serverid_Userid{
		code: broadcast_server,
		serverid: Int_Converter(c.user.serverid),
		userid: Int_Converter(c.user.id)}
	Message97(c.user, m.Build_Message())
}

//Creates Mute struct with param User type
type Mute struct {
	user *User
}

//Changes the user's muted to status to muted
func (c Mute) Execute() {
	c.user.muted = user_muted
}

//Builds a message code type 96, broadcast status
func (c Mute) Send() {
	m := Userid_Status{
		code: broadcast_status,
		userid: Int_Converter(c.user.id),
		status: c.user.status,
		muted: c.user.muted}
	Message96(c.user, m.Build_Message())
}

//Creates Unmute struct with param User type
type Unmute struct {
	user *User
}

//Changes the user's muted status to unmuted
func (c Unmute) Execute() {
	c.user.muted = user_unmuted
}

//Builds a message code type 96, broadcast status
func (c Unmute) Send() {
	m := Userid_Status{
		code: broadcast_status,
		userid: Int_Converter(c.user.id),
		status: c.user.status,
		muted: c.user.muted}
	Message96(c.user, m.Build_Message())
}

//Creates Move struct with param User type and roomid as an uint
type Move struct {
	user *User
	roomid uint
}

//Sets the user's roomid to the new roomid
func (c Move) Execute() {
	c.user.roomid = c.roomid
}

//Builds a message code type 98, broadcast room
func (c Move) Send() {
	m := Roomid_Userid{
		code: broadcast_room,
		roomid: Int_Converter(c.user.roomid),
		userid: Int_Converter(c.user.id)}
	Message98(c.user, m.Build_Message())
}

//TODO
type Prune struct {
	user *User
	num_messages uint
}

//TODO
func (c Prune) Execute() {

}

//TODO
func (c Prune) Send() {

}