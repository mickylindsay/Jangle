package main

import (
	"net"
)

const MOVE_USER 			= 0x0001
const KICK_USER 			= 0x0002
const DELETE_MESSAGE 	= 0x0004

type User struct {
	c *net.Conn
	display_name string
	id uint
	roomid uint
	serverid uint
	logged_in bool
	status byte
	permissions uint
	status uint
}

func (u *User) Read (read_data []byte) (int, error) {
	return (*(*u).c).Read(read_data);
}

func (u *User) Write (write_data []byte) (int, error) {
	data := make([]byte, len(write_data) + 4);
	copy(data[:3], Int_Converter(uint(len(write_data))));
	copy(data[4:], write_data[:]);
	return (*(*u).c).Write(data);
}

//Returns true if user has permission passed into function
func (u *User) Has_Permission (perm uint) bool {
	return (u.permissions & perm) != 0;
}

//Returns string representing the ip address of the local side of connection
func (u *User) Get_Local_Address () string{
	return (*(*u).c).LocalAddr().String();
}

//Returns string representing the ip address of the remote side of connection
func (u *User) Get_Remote_Address () string{
	return (*(*u).c).RemoteAddr().String();
}