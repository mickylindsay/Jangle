package main

import(
	"fmt"
)
type User_List struct{
	head User_Node
}

type User_Node struct {
	user User
	next *User_Node
}

type User struct{
	name string
}

func main(){
	micky := &User{
		name:"Micky",
	}

	lindsay := &User{
		name: "Lindsay",
	}

	head := new(User_Node)
	head.next = &User_Node{
		user: *micky,
	}
	fmt.Println(head.next.user)
	fmt.Println(lindsay.name)
}
