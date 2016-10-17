package main

import (
	"fmt"
	"log"
	"net"
	"bufio"
	"os"
	"flag"
)

type Client struct {

	debug bool
}
var client Client;

func main(){
	Init_Flags();
	
	read_data := make([]byte, 1024)
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	go func(){
		for {
			read_len, _ := conn.Read(read_data)
			if(read_len < 18){
				
			}else{
				fmt.Printf("%s\n", string(read_data[17:read_len]))
				if(client.debug){
					fmt.Println("IN: ", read_data[:read_len])
				}
			}
		}
	}()
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		write_data := make([]byte, len(text) + 12)
		write_data[0] = 16;
		copy(write_data[1:4], Int_Converter(1)); 
		copy(write_data[5:8], Int_Converter(1)); 
		copy(write_data[9:12], Int_Converter(1)); 
		copy(write_data[13:], []byte(text));
		/*write_data := make([]byte, len(text) + 12)
		write_data[0] = 32;
		copy(write_data[1:4], Int_Converter(1)); 
		copy(write_data[5:8], Int_Converter(1)); 
		copy(write_data[9:12], Int_Converter(1)); 
		write_data[13] = 0;*/
		if(client.debug){
			fmt.Println("OUT: ",write_data)
		}
		conn.Write(write_data)
		//send_message(conn,text)
	}
}

func send_message(conn net.Conn, text string){
	fmt.Fprintf(conn, "%s", text)
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
