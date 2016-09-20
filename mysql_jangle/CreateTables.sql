CREATE TABLE users(
	userid INT,
	username VARCHAR(20),
	displayname VARCHAR(20),
	imagepath VARCHAR(50),
	passwordhash VARCHAR(50),
	salt VARCHAR(8),
	status CHAR(1),
	PRIMARY KEY (userid)
);

CREATE TABLE servers(
	serverid INT,
	servername VARCHAR(20),	
	ownerid INT,
	FOREIGN KEY (ownerid)
        REFERENCES users (code)
	PRIMARY KEY (serverid)
);

CREATE TABLE friends(
	userid INT,
	friendid INT,	
	serverid INT,
	FOREIGN KEY (userid)
        REFERENCES users (code)
	FOREIGN KEY (friendid)
        REFERENCES users (userid)
	FOREIGN KEY (serverid)
        REFERENCES servers (serverid)
);

