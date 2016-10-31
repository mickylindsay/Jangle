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
	
	read_data := make([]byte, 1024);
	conn, err := net.Dial("tcp", "localhost:9090");
	if err != nil {
		log.Fatal(err);
	}
	go func(){
		for {
			read_len, _ := conn.Read(read_data);
			if(read_len < 18){
				
			}else{
				fmt.Printf("%s\n", string(read_data[17:read_len]));
				if(client.debug){
					fmt.Println("IN: ", read_data[:read_len]);
				}
			}
		}
	}()
	//for {
		
		write_data := Size_Meta_Data(Message65());
		if(client.debug){
			fmt.Println("OUT: ",write_data);
		}
		conn.Write(write_data);
	//}
	reader := bufio.NewReader(os.Stdin);
	reader.ReadString('\n');
}

//Converts unsigned int to byte array
func Int_Converter (num uint) []byte {
	data := make([]byte, 4);
	for i := 0; i < 4; i++ {
		mod := num % 256;
		data[i] = byte(mod);
		num /= 256;
	}
	return data;
}

func Init_Flags(){

	debug_flag := flag.Bool("debug", false, "Turns on client debugging");

	flag.Parse();

	client.debug = *debug_flag;
}
func Size_Meta_Data(b []byte) []byte {
	data := make([]byte, 4 + len(b));
	copy(data[0:3], Int_Converter(uint(len(b))));
	copy(data[4:], b[:]);
	return data;
}
func Message() []byte {
	reader := bufio.NewReader(os.Stdin);
	text, _ := reader.ReadString('\n');
	write_data := make([]byte, len(text) + 12 + 4);
	copy(write_data[0:3], Int_Converter(uint(len(text) + 12)));
	write_data[4] = 16;
	copy(write_data[5:8], Int_Converter(1)); 
	copy(write_data[9:12], Int_Converter(1)); 
	copy(write_data[13:16], Int_Converter(1)); 
	copy(write_data[17:], []byte(text));
	return write_data;
}

func Message33() []byte {
	write_data := make([]byte, 1);
	write_data[0] = 33;
	return write_data;
}

func Message34() []byte {
	write_data := make([]byte, 5);
	write_data[0] = 34;
	copy(write_data[1:4], Int_Converter(1));
	return write_data;
}

func Message35() []byte {
	write_data := make([]byte, 5);
	write_data[0] = 35;
	copy(write_data[1:4], Int_Converter(1));
	return write_data;
}

func Message36() []byte {
	write_data := make([]byte, 5);
	write_data[0] = 36;
	copy(write_data[1:4], Int_Converter(1));
	return write_data;
}

func Message37() []byte {
	write_data := make([]byte, 5);
	write_data[0] = 37;
	copy(write_data[1:4], Int_Converter(1));
	return write_data;
}

func Message38() []byte {
	write_data := make([]byte, 9);
	write_data[0] = 38;
	copy(write_data[1:4], Int_Converter(1));
	copy(write_data[5:8], Int_Converter(1));
	return write_data;
}

func Message39() []byte {
	write_data := make([]byte, 5);
	write_data[0] = 39;
	copy(write_data[1:4], Int_Converter(1));
	return write_data;
}

func Message40() []byte {
	write_data := make([]byte, 5);
	write_data[0] = 40;
	copy(write_data[1:4], Int_Converter(1));
	return write_data;
}

func Message64() []byte {
	reader := bufio.NewReader(os.Stdin);
	text, _ := reader.ReadString('\n');
	write_data := make([]byte, len(text) + 12);
	write_data[0] = 64;
	copy(write_data[1:], []byte(text));
	return write_data;
}

func Message65() []byte {
	reader := bufio.NewReader(os.Stdin);
	text, _ := reader.ReadString('\n');
	write_data := make([]byte, len(text) + 12);
	write_data[0] = 65;
	copy(write_data[1:4], Int_Converter(1));
	copy(write_data[5:], []byte(text));
	return write_data;
}

func Message66() []byte {
	reader := bufio.NewReader(os.Stdin);
	text, _ := reader.ReadString('\n');
	write_data := make([]byte, len(text) + 12);
	write_data[0] = 66;
	copy(write_data[1:4], Int_Converter(1));
	copy(write_data[5:8], Int_Converter(1));
	copy(write_data[9:], []byte(text));
	return write_data;
}

func Message67() []byte {
	reader := bufio.NewReader(os.Stdin);
	text, _ := reader.ReadString('\n');
	write_data := make([]byte, len(text) + 12);
	write_data[0] = 67;
	copy(write_data[1:], []byte(text));
	return write_data;
}