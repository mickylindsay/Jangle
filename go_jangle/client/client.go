package main

import (
       "fmt"
       "log"
       "net"
)

func main(){
     conn, err := net.Dial("tcp", "localhost:9090")
     if err != nil {
	log.Fatal(err)
	}    
     fmt.Fprintf(conn, "Hello from Client")
}