package main

import (
	"fmt"
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
func Kick_User (agrs []string) {

}

//TODO
func Mute_User (args []string) {

}

//TODO
func Unmute_User (args []string) {

}

//TODO
func Move_User (args []string) {

}

//TODO
type Command interface {
	Execute([]string)
	Send()
}

//TODO
type Kick struct {

}

//TODO
func (c Kick) Execute(args []string) {
	user := Get_User_From_Userid(Byte_Converter([]byte(args[1])))
	user.serverid = default_value
	user.roomid = default_value
}

//TODO
func (c Kick) Send() {

}

//TODO
type Mute struct {

}

//TODO
func (c Mute) Execute(args []string) {
	user := Get_User_From_Userid(Byte_Converter([]byte(args[1])))
	user.muted = 1
}

//TODO
func (c Mute) Send() {

}

//TODO
type Unmute struct {

}

//TODO
func (c Unmute) Execute(args []string) {
	user := Get_User_From_Userid(Byte_Converter([]byte(args[1])))
	user.muted = 2
}

//TODO
func (c Unmute) Send() {

}

//TODO
type Move struct {

}

//TODO
func (c Move) Execute(args []string) {
	user := Get_User_From_Userid(Byte_Converter([]byte(args[1])))
	user.roomid = Byte_Converter([]byte(args[1]))
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