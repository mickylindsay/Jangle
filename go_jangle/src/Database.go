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

/*func main() {
	//Attempt to connect to MySQL Database
	db, err := Connect_Database()
	if err != nil {
		log.Fatal(err)
	}
	//Test MySQL Execution
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users(userid INT,	username VARCHAR(20),	displayname VARCHAR(20),	imagepath VARCHAR(50),passwordhash VARCHAR(50),	salt VARCHAR(8),	status CHAR(1),	PRIMARY KEY (userid));")
	if err != nil {
		log.Fatal(err)
	}
}*/
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
	return sql.Open("mysql", dsn)
}

func User_Login(u []byte, p []byte) (int, error) {
	var userid int;
	err := jangle.db.QueryRow("SELECT userid FROM users WHERE username =? AND passwordhash=?",string(u), string(p)).Scan(&userid);
	return userid, err;
}

func User_Create(u []byte, p []byte) error{
	_, err := jangle.db.Exec("INSERT INTO users (userid, username, displayname, imagepath, passwordhash, salt) VALUES (?,?,?,?,?,?);",Next_Userid()+1, string(u), string(u), "TEMPPATH", string(p), "0000");
	return err;
}

func Next_Userid() int{
	var temp int;
	_ = jangle.db.QueryRow("SELECT MAX(userid) AS userid FROM users").Scan(&temp);
	return temp;
}
