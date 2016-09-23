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
	dat, err := ioutil.ReadFile(dir + "/.databasedsn")
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
