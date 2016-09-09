package main

import (
       "fmt"
       "net"
       "log"
)

func main(){
     address := "localhost:9090"
     read_data := make([]byte, 1024)
     
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
	     	read_len, err := c.Read(read_data)
		if err != nil {
		   log.Fatal(err)
		}
		read_string := string(read_data)
		fmt.Printf("Read %d bytes\n", read_len);
		fmt.Println(read_string)
	 	c.Close()
	 }(conn)

     }
}