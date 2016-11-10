package main

import (
	"net"
)

const MOVE_USER 	= 0x0001
const KICK_USER 	= 0x0002
const DELETE_MESSAGE 	= 0x0004
const EDIT_MESSAGE      = 0x0008
const PERM_5            = 0x0010

type User struct {
	c *net.Conn
	display_name []byte
	id uint
	roomid uint
	serverid uint
	logged_in bool
	muted byte
	status byte
	permissions uint
	location *Address
}

func (u *User) equals(u2 *User) bool{
	return u.id == u2.id;
}

func (u *User) User_Initialize() err{
	u.roomid = 1;
	u.serverid = 1;
	u.status = online;
	u.location = &Address{
		userid : u.id,
		roomid : u.roomid,
		serverid : u.serverid,
	}
	display, err := Request_Display_Name(u.serverid, u.id);
	if(err != nil) {
		return err;
	}
	u.display_name = display;
	data := make([]byte, 6);
	data[0] = broadcast_status;
	copy(data[1:4], Int_Converter(u.id));
	data[5] = u.status;
	Message96(u, data);
	return nil;
}

func (u *User) Get_Userid() uint{
	return u.id;
}

func (u *User) Get_Serverid() uint{
	return u.serverid;
}

func (u *User) Get_Roomid() uint{
	return u.roomid;
}

func (u *User) Set_Userid(id uint) {
	u.id = id;
}

func (u *User) Set_Roomid(id uint) {
	u.roomid = id;
}

func (u *User) Set_Serverid(id uint) {
	u.serverid = id;
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
