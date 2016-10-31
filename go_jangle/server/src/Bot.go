package main

import (
	"fmt"
)

type Bot struct {
	serverid uint8
}

func (b *Bot) Bot_Move_User(userid uint, roomid uint) error {
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		if (e.Value.(*User).id == userid) {
			e.Value.(*User).roomid = roomid;
			return nil;
		}
	}
	return fmt.Errorf("Bot %d: Unable to Move User %d. No such User", b.serverid, userid);
}

//Possibly create message 18 Bot message
/*
func (b *Bot) Bot_Broadcast(text []byte, roomid uint){
	m := Message_Recieve{
		code: 17,
		serverid: Int_Converter(b.serverid),
		roomid: Int_Converter(roomid),
		userid: Int_Converter()
	}
}
*/
