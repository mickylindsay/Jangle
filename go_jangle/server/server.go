package main

import (
       "fmt"
       "net"
       "log"
)

func main(){
     address := "localhost:9090"
     num_connections := 0
     fmt.Println("JANGLE GO SERVER")
     fmt.Println("address - " + address)

     listener, err := net.Listen("tcp", address)
     if err != nil{
     	log.Fatal(err)
     }
     defer listener.Close()
     for {
     	 conn, err := listener.Accept()
	 if err != nil {
	    log.Fatal(err)	
	 }
	 go func(c net.Conn){
	    num_connections++
	    fmt.Printf("%d\n", num_connections)
	    c.Close()
	 }(conn)

     }
}