package main

import (
	"fmt"
	"log"
	"net"
//	"bufio"
//	"os"
	"flag"
	"github.com/jroimartin/gocui"
)

type Client struct {
	conn *net.Conn
	debug bool
	g *gocui.Gui	
}
var client Client;

func main(){
	Init_Flags();
	read_data := make([]byte, 1024)
	conn, err := net.Dial("tcp", "localhost:9090")
	client.conn = &conn;
	if err != nil {
		log.Fatal(err)
	}
	GUI_Init();

	go func(){
		for {
			read_len, _ := (*(client).conn).Read(read_data)
			if(read_len >= 18){
				Append_Message(read_data[17:read_len]);
				if(client.debug){
					fmt.Println("IN: ", read_data[:read_len])
				}
			}
		}
	}();
	GUI_Run();
}

func Write_To_Server(data []byte){
	(*(client).conn).Write(data);
}
//Converts unsigned int to byte array
func Int_Converter (num uint) []byte {
	data := make([]byte, 4)
	for i := 0; i < 4; i++ {
		mod := num % 256
		data[i] = byte(mod)
		num /= 256
	}
	return data
}

func Init_Flags(){

	debug_flag := flag.Bool("debug", false, "Turns on client debugging");

	flag.Parse();

	client.debug = *debug_flag;
}
