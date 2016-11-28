package main

import (
	"net"
)

const MOVE_USER = 0x0001
const KICK_USER = 0x0002
const DELETE_MESSAGE = 0x0004

type User struct {
	c            *net.Conn
	display_name string
	id           uint
	roomid       uint
	serverid     uint
	logged_in    bool
	voice        uint
	muted        uint
	status       uint
	permissions  uint
	location     *Address
}

func (u *User) equals(u2 *User) bool {
	return u.id == u2.id
}

func (u *User) User_Initialize() {
	u.roomid = 1
	u.serverid = 1
	Join_Server(u);
	u.status = uint(online)
	u.location = &Address{
		userid:   u.id,
		roomid:   u.roomid,
		serverid: u.serverid,
	}
	m := Create_Message(recieve_status, Int_Converter(u.id), byte(u.status), byte(u.muted))
	Send_Broadcast_Server(u.serverid, m)
	Send_Broadcast_Friends(u.id, m)
}

func (u *User) Get_Userid() uint {
	return u.id
}

func (u *User) Get_Serverid() uint {
	return u.serverid
}

func (u *User) Get_Roomid() uint {
	return u.roomid
}

func (u *User) Set_Userid(id uint) {
	u.id = id
}

func (u *User) Set_Roomid(id uint) {
	u.roomid = id
}

func (u *User) Set_Serverid(id uint) {
	u.serverid = id
	Join_Server(u);
}

func (u *User) Read(read_data []byte) (int, error) {
	return (*(*u).c).Read(read_data)
}

func (u *User) Write(write_data []byte) (int, error) {
	data := make([]byte, len(write_data)+4)
	copy(data[:3], Int_Converter(uint(len(write_data))))
	copy(data[4:], write_data[:])
	return (*(*u).c).Write(data)
}

//Returns true if user has permission passed into function
func (u *User) Has_Permission(perm uint) bool {
	return (u.permissions & perm) != 0
}

//Returns string representing the ip address of the local side of connection
func (u *User) Get_Local_Address() string {
	return (*(*u).c).LocalAddr().String()
}

//Returns string representing the ip address of the remote side of connection
func (u *User) Get_Remote_Address() string {
	return (*(*u).c).RemoteAddr().String()
}
>>>>>>> 81a2ccfc63dde464c167b934ce6c030180bcb2f6
