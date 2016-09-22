package main
import (
	"net"
	"fmt"

)

type User struct {
	c *net.Conn
	name string
	id int
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

func (u *User) Message(message string) {
	u.Printf("%s: %s", u.name, message)
}
