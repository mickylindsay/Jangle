package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"path/filepath"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//Returns a connection to the mysql database at the location either prompted or found in the file .databasedsn in the directory of the executable
func Connect_Database() (*sql.DB, error){
	if(!jangle.no_database){
		fmt.Println("Connecting to MySQL Database.")
		//Finds executables current directory and read the data in .databasedsn
		var dsn string;
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
		dat, err := ioutil.ReadFile(dir + "/../.databasedsn");
		//If such file does not exist prompt the user to enter a DSN
		if err != nil{
			fmt.Println("Please enter mysql database DSN:");
			reader := bufio.NewReader(os.Stdin);
			text, _ := reader.ReadString('\n');
			dsn = text[:len(text)-1];
		}else{
			dsn = string(dat);
		}
		//Attempts to open a conncetion to the mysql database
		db , err := sql.Open("mysql", dsn);
		if(err == nil){
			Color_Println("green", "Database Connection Successful.");
		}else{
			Color_Println("orange", "Failed to connect to Database.");
			fmt.Println("Running in 'no_database' mode.");
			jangle.no_database = true;
		}
		return db, err;
	}
	fmt.Println("Running in 'no_database' mode.")
	return nil, nil;
	
}

//Returns userid from database when correct username and password(hash) is given
func User_Login(u []byte, p []byte) (uint, error) {
	if(!jangle.no_database){
		var userid uint;
		err := jangle.db.QueryRow("SELECT userid FROM users WHERE username =? AND passwordhash=?",string(u), string(p)).Scan(&userid);
		return userid, err;
	}
	return 1, nil;
}

//Returns the next valid userid
func Next_Userid() uint{
	if(!jangle.no_database){
		var temp uint;
		_ = jangle.db.QueryRow("SELECT MAX(userid) AS userid FROM users").Scan(&temp);
		return temp + 1;
	}
	return 1;
}

//Returns the next valid messageid
func Next_Messageid() uint{
	if(!jangle.no_database){
		var temp uint;
		_ = jangle.db.QueryRow("SELECT MAX(messageid) AS messageid FROM messages").Scan(&temp);
		return temp + 1;
	}
	return 1;
}

//Inserts a new user into the database
//TODO implement Image Path and Password hashing
func User_Create(u []byte, p []byte) error{
	if(!jangle.no_database){
		_, err := jangle.db.Exec("INSERT INTO users (userid, username, displayname, imagepath, passwordhash, salt) VALUES (?,?,?,?,?,?);",Next_Userid(), string(u), string(u), "TEMPPATH", string(p), "0000");
		return err;
	}
	return nil;
}

//Inserts a new Message into the database
//TODO implement roomid and serverid
func Message_Create(user *User, messagetext []byte) error{
	if(!jangle.no_database){
		var err error
		if(user.id == 0){
			_, err = jangle.db.Exec("INSERT INTO messages (userid, time, messageid, messagetext, serverid, roomid) VALUES (?,?,?,?,?,?);", 1, Milli_Time(), Next_Messageid(), string(messagetext), 1, 1);
		}else{
			_, err = jangle.db.Exec("INSERT INTO messages (userid, time, messageid, messagetext, serverid, roomid) VALUES (?,?,?,?,?,?);", user.id, Milli_Time(), Next_Messageid(), string(messagetext), user.serverid, user.roomid);
		}
		return err;
	}
	return nil;
}


//Request chunks of 50 messages offset by (offset*50) and returns them as array of message objects
func Request_Offset_Messages(offset uint) ([]Message, error){
	if(!jangle.no_database){
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
	return nil, nil;
}

//Requests all userids with same serverid as user
func Request_Userid_Messages(serverid uint) ([]Message, error){
	if(!jangle.no_database){
		var userid uint;
		i := 0;
		messages := make([]Message, 50);
		//Query 50 rows of messages
		rows, err := jangle.db.Query("SELECT userid FROM users AS u WHERE ? = u.serverid",serverid);
		Check_Error(err);
		defer rows.Close();
		//Iterate through the rows
		for rows.Next() {
			//Scan the columns into variables
			err := rows.Scan(&userid);
			Check_Error(err);
			//Create a "48" message to send back to user
			m := Userid{
				code: 48,
				userid: Int_Converter(userid)};
			//Add that message to the array which will be returned
			messages = append(messages, m);
			i++;
		}
		//Return array of messages
		return messages, err;
	}
	return nil, nil;
}

//Request all serverids which a userid is in
func Request_Serverid_Messages(userid uint) ([]Message, error){
	if(!jangle.no_database){
		var serverid uint;
		i := 0;
		messages := make([]Message, 50);
		//Query 50 rows of messages
		rows, err := jangle.db.Query("SELECT serverid FROM members AS m WHERE ? = m.userid", userid);
		Check_Error(err);
		defer rows.Close();
		//Iterate through the rows
		for rows.Next() {
			//Scan the columns into variables
			err := rows.Scan(&serverid);
			Check_Error(err);
			//Create a "50" message to send back to user
			m := Serverid_Userid{
				code: 50,
				serverid: Int_Converter(serverid),
				userid: Int_Converter(userid)};
			//Add that message to the array which will be returned
			messages = append(messages, m);
			i++;
		}
		//Return array of messages
		return messages, err;
	}
	return nil, nil;
}

//Request the Name of a server by serverid
func Request_Server_Display_Name(serverid uint) ([]byte,error) {
	if(!jangle.no_database){
		var temp string;
		err := jangle.db.QueryRow("SELECT servername FROM servers AS s WHERE s.serverid = ?", serverid).Scan(&temp);
		return []byte(temp), err;
	}
	return []byte("TEMP"), nil;
}

//Request all Roomids a server contains
func Request_Roomid_Messages(serverid uint) ([]Message, error){
	var roomid uint;
	i := 0;
	messages := make([]Message, 50);
	if(!jangle.no_database){
		rows, err := jangle.db.Query("SELECT roomid FROM rooms AS r WHERE ? = r.serverid", serverid);
		Check_Error(err);
		defer rows.Close();
		//Iterate through the rows
		for rows.Next() {
			//Scan the columns into variables
			err := rows.Scan(&roomid);
			Check_Error(err);
			//Create a "52" message to send back to user
			m := Serverid_Userid{
				code: 52,
				serverid: Int_Converter(serverid),
				userid: Int_Converter(roomid)};
			//Add that message to the array which will be returned
			messages = append(messages, m);
			i++;
		}
		//Return array of messages
		return messages, err;
	}
	return nil, nil;
}

//Request the name of the Room by serverid and roomid
func Request_Room_Display_Name(serverid uint, roomid uint) ([]byte,error) {
	if(!jangle.no_database){
		var temp string;
		err := jangle.db.QueryRow("SELECT roomname FROM roomss AS r WHERE r.serverid = ? AND r.roomid = ?", serverid, roomid).Scan(&temp);
		return []byte(temp), err;
	}
	return []byte("TEMP"), nil;
}

//Request the Name of a server by serverid
func Request_Display_Name(serverid uint, userid uint) ([]byte,error) {
	if(!jangle.no_database){
		var temp string;
		err := jangle.db.QueryRow("SELECT displayname FROM display WHERE serverid = ? and userid = ?", serverid, userid).Scan(&temp);
		if(err==nil){
			err = jangle.db.QueryRow("SELECT displayname FROM users WHERE userid = ?", userid).Scan(&temp);
		}
		return []byte(temp), err;
	}
	return []byte("TEMP_DISPLAY_NAME"), nil;
}

//Inserts or update a new server specific displayname
func Set_New_Display_Name(serverid uint, userid uint, name []byte) error{
	if(!jangle.no_database){
		err := jangle.db.QueryRow("SELECT displayname FROM display WHERE serverid = ? and userid = ?", serverid, userid);
		if(err != nil){
			_, e := jangle.db.Exec("UPDATE display  SET displayname = ? WHERE userid = ? AND serverid = ?;", string(name), userid, serverid);
			return e;
		}else{	
			_, e := jangle.db.Exec("INSERT INTO display (userid, serverid, displayname) VALUES (?,?,?);", userid, serverid, string(name));
			return e;
		}
	}
	return nil;
}
