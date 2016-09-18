package main

import(
	"fmt"
	"container/list"
)

type User struct {
	name string
}

func main(){
	
	micky := &User{
		name:"Micky",
	}

	lindsay := &User{
		name: "Lindsay",
	}
	l := list.New()
	e := l.PushFront(micky)
	l.InsertAfter(lindsay, e)
	for element := l.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	
}
