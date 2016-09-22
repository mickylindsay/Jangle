package main

import "fmt"
import "bufio"
import "os"

type Message interface {
	Build_message() []byte
}

type Base struct {
	code byte
}

func(m Base) Build_message() []byte {
	var message [1]byte
	message[0] = m.code
	return message[:]
}

type Userpass struct {
	code byte
	username [20]byte
	password []byte
}

func(m Userpass) Build_message() []byte {
	message := make([]byte, 21 + len(m.password))
	message[0] = m.code
	copy(message[1:20], m.username[:])
	copy(message[21:], m.password[:])
	return message
}
/*type Message struct {
	code byte
	username []byte
	password []byte
	serverid []byte
	userid []byte
	display_name []byte
	text []byte
	num []byte
}*/

func parse_text (data []byte) {
	var create_user byte = 0
	create_user_fail := "1"
	login := "2"
	login_fail := "3"
	login_success := "4"

	message_client_send := "16"
	message_client_recieve := "17"

	request_n_messages := "32"
	request_all_userid := "33"
	request_display_name := "34"
	request_all_serverid := "35"
	request_server_display_name := "36"

	recieve_userid := "48"
	recieve_display_name := "49"
	recieve_serverid := "50"

	send_new_display_name := "64"

	if (data[0] == create_user[0]) {
		
	}

	if (data[0] == create_user_fail[0]) {
	
	}

	if(data[0] == login[0]) {
		
	}

	if(data[0] == login_fail[0]) {

	}

	if(data[0] == login_success[0]) {

	}

	if(data[0] == message_client_send[0]) {

	}

	if(data[0] == message_client_recieve[0]) {

	}

	if(data[0] == request_n_messages[0]) {

	}

	if(data[0] == request_all_userid[0]) {

	}

	if(data[0] == request_display_name[0]) {

	}

	if(data[0] == request_all_serverid[0]) {

	}

	if(data[0] == request_server_display_name[0]) {

	}

	if(data[0] == recieve_userid[0]) {

	}

	if(data[0] == recieve_display_name[0]) {
	
	}

	if(data[0] == recieve_serverid[0]) {

	}

	if(data[0] == send_new_display_name[0]) {

	}
}

func main() {
	fmt.Print("Enter text: ")
    reader := bufio.NewReader(os.Stdin)
    text,_:= reader.ReadString('\n')
	arr := []byte(text)
	parse_text(arr[:])
	
 }