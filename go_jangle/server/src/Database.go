package main

import (
	"fmt"
	//"log"	
	"os"
	"bufio"
	"io/ioutil"
	"path/filepath"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//Returns a connection to the mysql database at the location either prompted or found in the file .databasedsn in the directory of the executable
func Connect_Database() (*sql.DB, error){
	//Finds executables current directory and read the data in .databasedsn
	var dsn string
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) 
	dat, err := ioutil.ReadFile(dir + "/../.databasedsn")
	//If such file does not exist prompt the user to enter a DSN
	if err != nil{
		fmt.Println("Please enter mysql database DSN:")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		dsn = text[:len(text)-1]
	}else{
		dsn = string(dat)
	}
	//Attempts to open a conncetion to the mysql database
	db , err := sql.Open("mysql", dsn);
	return db, err
}

//Returns userid from database when correct username and password(hash) is given
func User_Login(u []byte, p []byte) (uint, error) {
	var userid uint;
	err := jangle.db.QueryRow("SELECT userid FROM users WHERE username =? AND passwordhash=?",string(u), string(p)).Scan(&userid);
	return userid, err;
}

//Returnst the next valid userid
func Request_Display_Name(id uint) []byte{
	var temp []byte;
	_ = jangle.db.QueryRow("SELECT displayname FROM users WHERE userid = ?", id).Scan(&temp);
	return temp;
}

//Returnst the next valid userid
func Next_Userid() uint{
	var temp uint;
	_ = jangle.db.QueryRow("SELECT MAX(userid) AS userid FROM users").Scan(&temp);
	return temp + 1;
}

//Returns the next valid messageid
func Next_Messageid() uint{
	var temp uint;
	_ = jangle.db.QueryRow("SELECT MAX(messageid) AS messageid FROM messages").Scan(&temp);
	return temp + 1;
}

//Inserts a new user into the database
//TODO implement Image Path and Password hashing
func User_Create(u []byte, p []byte) error{
	_, err := jangle.db.Exec("INSERT INTO users (userid, username, displayname, imagepath, passwordhash, salt) VALUES (?,?,?,?,?,?);",Next_Userid(), string(u), string(u), "TEMPPATH", string(p), "0000");
	return err;
}

//Inserts a new Message into the database
//TODO implement roomid and serverid
func Message_Create(user *User, messagetext []byte) error{
	var err error
	if(user.id == 0){
		_, err = jangle.db.Exec("INSERT INTO messages (userid, time, messageid, messagetext, serverid, roomid) VALUES (?,?,?,?,?,?);", 1, Milli_Time(), Next_Messageid(), string(messagetext), 1, 1);
	}else{
	_, err = jangle.db.Exec("INSERT INTO messages (userid, time, messageid, messagetext, serverid, roomid) VALUES (?,?,?,?,?,?);", user.id, Milli_Time(), Next_Messageid(), string(messagetext), 1, 1);
	}
	return err;
}

//Request chunks of 50 messages offset by (offset*50) and returns them as array of message objects
func Request_Offset_Messages(offset uint) ([]Message, error){
	i := 0;
	messages := make([]Message,50);
	var(
		time_read uint
		text_read string
		userid_read uint
	)
	//Query 50 rows of messages
	rows, err := jangle.db.Query("SELECT userid, time, messagetext FROM messages ORDER BY messageid DESC LIMIT 50 OFFSET  ?", offset*50)
	Check_Error(err);
	defer rows.Close();
	//Iterate through the rows
	for rows.Next() {
		//Scan the columns into variables
		err := rows.Scan(&userid_read, &time_read, &text_read);
		fmt.Println(text_read);
		Check_Error(err);
		//Create a "17" message to send back to user
		m := Message_Recieve{
			code: 17,
			serverid: Int_Converter(0),
			roomid: Int_Converter(0),
			userid: Int_Converter(userid_read),
			time: Int_Converter(time_read),
			text: []byte(text_read)};
		//Add that message to the array which will be returned
		messages[i] = m;
		i++;
	}
	//Return array of messages
	return messages, err;
}
