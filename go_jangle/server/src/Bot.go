package main

import (
	"fmt"
	"strings"
)

//TODO
func Init_Command () {
	Commands := make([]func(arg1 string, arg2 string), 3)

	Commands[0] = Kick_User
	Commands[1] = Mute_User
	Commands[2] = Move_User

	jangle.Commands = Commands
}

//TODO
func Check_Command (user *User, data []byte) bool {
	var check bool
	if (string(data[0]) == "/") {
		check = true
		args := strings.Split(string(data), " ")
		trigger := Switcher(args[0])
		jangle.Commands[trigger](args[1], args[2])
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
		case s == "move":
			switcher = 2
	}
	return switcher
}

//TODO
func Kick_User (agr1 string, arg2 string) {

}

//TODO
func Mute_User (arg1 string, arg2 string) {

}

//TODO
func Move_User (arg1 string, arg2 string) {

}

//TODO
type Command interface {
	Execute()
	Send()
}

//TODO
type Kick struct {

}

//TODO
func (c Kick) Execute() {

}

//TODO
func (c Kick) Send() {

}

//TODO
type Mute struct {

}

//TODO
func (c Mute) Execute() {

}

//TODO
func (c Mute) Send() {

}

//TODO
type Move struct {

}

//TODO
func (c Move) Execute() {

}

//TODO
func (c Move) Send() {

}

//TODO
type Bot struct {

}

//Bot attempts to move a user from one room to another returns and error if no user with this userid is connected
func (b *Bot) Bot_Move_User(userid uint, serverid uint, roomid uint) error {
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		if (e.Value.(*User).id == userid) {
			e.Value.(*User).roomid = roomid;
			return nil;
		}
	}
	return fmt.Errorf("Bot %d: Unable to Move User %d. No such User", serverid, userid);
}

//Bot sends message to all users
func (b *Bot) Bot_Broadcast(text []byte, serverid uint, roomid uint){
	m := Message_Recieve{
		code: 17,
		serverid: Int_Converter(serverid),
		roomid: Int_Converter(roomid),
		userid: Int_Converter(1),
		time: Int_Converter(Milli_Time()),
		text: text[:],
	};
	Send_Broadcast_Server_Room(serverid, roomid, m);
}