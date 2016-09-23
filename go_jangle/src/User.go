package main

import (
	"net"
	"fmt"
	"container/list"
)

type User struct {
	c *net.Conn
	display_name string
	id int
	logged_in bool
}

/*func NewUser(conn *new.Conn, userid int) *User{
	user := &User{
		c: conn,
		id: userid,
		logged_in: false,
	}
	return user
}*/

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

func Send_Message(users *list.List, id int, message Message){
	for e := users.Front(); e != nil; e = e.Next() {
		if e.Value.(*User).id == id {
			e.Value.(*User).Write(message.Build_message())
		}
	}			
}

func Sent_Broadcast(users *list.List, id int, message Message){
	for e := users.Front(); e != nil; e = e.Next() {
		e.Value.(*User).Write(message.Build_message())
	}			
}
