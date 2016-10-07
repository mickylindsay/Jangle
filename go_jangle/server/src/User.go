package main

import (
	"net"
	"fmt"
)

type User struct {
	c *net.Conn
	display_name string
	id uint
	roomid uint
	serverid uint
	logged_in bool
}

func (u *User) Read(read_data []byte) (int, error){
	return (*(*u).c).Read(read_data)
}

func (u *User) Write(write_data []byte) (int, error){
	return (*(*u).c).Write(write_data)
}

func (u *User) Printf(format string, a ...interface{}) (int, error){
	return fmt.Fprintf((*(*u).c), format, a...)
}

func (u *User) Scanf(format string, a ...interface{}) (int, error){
	return fmt.Fscanf(*(*u).c, format, a...)
}

func Send_Message(user *User, message Message) uint{
	user.Write(message.Build_Message())
	
	return 0;
}

func Send_Broadcast(message Message){
	for e := jangle.userlist.Front(); e != nil; e = e.Next() {
		e.Value.(*User).Write(message.Build_Message())
	}			
}
