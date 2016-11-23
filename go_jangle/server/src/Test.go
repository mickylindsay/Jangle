package main

import (
	"bufio"
	"fmt"
	"os"
)

//TODO
func Size_Meta_Data(b []byte) []byte {
	data := make([]byte, 4+len(b))
	copy(data[0:3], Int_Converter(uint(len(b))))
	copy(data[4:], b[:])
	return data
}

//TODO
func Create_User_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	var username string
	var valid_username bool = false
	for valid_username != true {
		fmt.Print("Enter New Username: ")
		username, _ = reader.ReadString('\n')
		if len(username) > 20 {
			fmt.Print("Invalid Username: must be 20 characters or less\n")
		} else {
			valid_username = true
		}
	}
	fmt.Print("Enter New Password: ")
	password, _ := reader.ReadString('\n')
	temp := make([]byte, 20)
	copy(temp, []byte(username))
	m := Create_Message(create_user, temp, []byte(password))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Login_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Enter Password: ")
	password, _ := reader.ReadString('\n')
	temp := make([]byte, 20)
	copy(temp, []byte(username))
	m := Create_Message(login, temp, []byte(password))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Standard_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Message: ")
	text, _ := reader.ReadString('\n')
	m := Create_Message(message_client_send, default_id, default_id, default_id, []byte(text))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Offset_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Offset Value: ")
	offset, _ := reader.ReadString('\n')
	m := Create_Message(request_n_messages, []byte(offset))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Multi_Userid_Message_Test() []byte {
	m := Create_Message(request_all_userid)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Display_Name_Message_Test() []byte {
	m := Create_Message(request_display_name, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Multi_Serverid_Message_Test() []byte {
	m := Create_Message(request_all_serverid, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Server_Display_Name_Message_Test() []byte {
	m := Create_Message(request_server_display_name, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Multi_Roomid_Message_Test() []byte {
	m := Create_Message(request_all_roomid, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Room_Display_Name_Message_Test() []byte {
	m := Create_Message(request_room_display_name, default_id, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Master_Display_Name_Message_Test() []byte {
	m := Create_Message(request_master_display_name, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Status_Message_Test() []byte {
	m := Create_Message(request_status, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Location_Message_Test() []byte {
	m := Create_Message(request_location, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func User_Ip_Message_Test() []byte {
	m := Create_Message(request_user_ip, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func User_Icon_Message_Test() []byte {
	m := Create_Message(request_user_icon, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Server_Icon_Message_Test() []byte {
	m := Create_Message(request_server_icon, default_id)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func New_Display_Name_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter New Display Name: ")
	display_name, _ := reader.ReadString('\n')
	m := Create_Message(send_new_display_name, []byte(display_name))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func New_Server_Display_Name_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter New Server Display Name: ")
	server_display_name, _ := reader.ReadString('\n')
	m := Create_Message(send_new_server_display_name, default_id, []byte(server_display_name))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func New_Room_Display_Name_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter New Room Display Name: ")
	room_display_name, _ := reader.ReadString('\n')
	m := Create_Message(send_new_room_display_name, default_id, default_id, []byte(room_display_name))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func New_Master_Display_Name_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter New Master Display Name: ")
	master_display_name, _ := reader.ReadString('\n')
	m := Create_Message(send_new_master_display_name, []byte(master_display_name))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func New_User_Icon_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter New User Icon Url: ")
	url, _ := reader.ReadString('\n')
	m := Create_Message(send_new_user_icon, []byte(url))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func New_Server_Icon_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter New Server Icon Url: ")
	url, _ := reader.ReadString('\n')
	m := Create_Message(send_new_server_icon, default_id, []byte(url))
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Change_Status_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	var valid_status bool = false
	var valid_muted bool = false
	var valid_voice bool = false
	var switcher_status byte
	var switcher_muted byte
	var switcher_voice byte
	for valid_status != true {
		fmt.Print("Enter New Status (offline, online, away): ")
		status, _ := reader.ReadString('\n')
		switch {
		case status == "offline":
			switcher_status = offline
		case status == "online":
			switcher_status = online
		case status == "away":
			switcher_status = away
		default:
			switcher_status = 255
		}
		if switcher_status != 255 {
			valid_status = true
		} else {
			fmt.Print("Invalid Status: must be offline, online, away")
		}
	}
	for valid_muted != true {
		fmt.Print("Enter New Muted (unmuted, muted): ")
		muted, _ := reader.ReadString('\n')
		switch {
		case muted == "unmuted":
			switcher_muted = user_unmuted
		case muted == "muted":
			switcher_muted = user_muted
		default:
			switcher_muted = error_check
		}
		if switcher_muted != error_check {
			valid_muted = true
		} else {
			fmt.Print("Invalid Muted: must be unmuted or muted")
		}
	}
	for valid_voice != true {
		fmt.Print("Enter New Voice (novoice, voice): ")
		voice, _ := reader.ReadString('\n')
		switch {
		case voice == "novoice":
			switcher_voice = user_no_voice
		case voice == "voice":
			switcher_voice = user_voice
		default:
			switcher_voice = error_check
		}
		if switcher_voice != error_check {
			valid_voice = true
		} else {
			fmt.Print("Invalid Voice: must be novoice or voice")
		}
	}
	m := Create_Message(change_status, switcher_status, switcher_muted, switcher_voice)
	return Size_Meta_Data(m.Build_Message())
}

//TODO
func Change_Location_Message_Test() []byte {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter New Server Location: ")
	server_location, _ := reader.ReadString('\n')
	fmt.Print("Enter New Room Location: ")
	room_location, _ := reader.ReadString('\n')
	m := Create_Message(change_location, String_Converter(server_location), String_Converter(room_location))
	return Size_Meta_Data(m.Build_Message())
}
