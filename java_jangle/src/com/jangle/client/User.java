package com.jangle.client;

import java.io.InputStream;

/**
 * Created by Jess on 9/28/2016.
 */
public class User {

    private final String DEFAULT_AVATAR = "http://res.freestockphotos.biz/pictures/17/17384-illustration-of-a-red-santa-hat-pv.png";

    private int id;
    private int status;
    private String displayName;
    private String userName;
    private String avatarURL;
    private String IP;
    

    public User(String displayName, int id, int status) {
        this.displayName = displayName;
        this.id = id;
        this.status = status;
        this.userName = "";
        this.avatarURL = DEFAULT_AVATAR;
        this.IP = "";
    }
    
    public User(String displayName, String userName, int id, int status){
    	this.displayName = displayName;
        this.id = id;
        this.status = status;
        this.userName = userName;
        this.avatarURL = DEFAULT_AVATAR;
        this.IP = "";
    }

    public User(String displayName, int id) {
        this.displayName = displayName;
        this.id = id;
        this.status = 0;
        this.userName = "";
        this.avatarURL = DEFAULT_AVATAR;
        this.IP = "";
    }

    public String getDisplayName() {
        return displayName;
    }

    public void setDisplayName(String displayName) {
        this.displayName = displayName;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getStatus() {
        return status;
    }

    public void setStatus(int status) {
        this.status = status;
    }
    
    public String getUserName(){
    	return userName;
    }
    
    public void setUserName(String userName){
    	this.userName = userName;
    }

    public String toString(){
        return displayName;
    }
    
    /**
     * Used to compare if this user object is the same as the passed in user object
     * @param user User object to compare equality with
     * @return true if the same, false if not the same
     */
    public boolean equals(User user){
    	if (this.id == user.getId()){
    		return true;
    	}
    	else{
    		return false;
    	}
    }

    public String getAvatarURL() {
        return avatarURL;
    }

    public void setAvatar(String newURL) {
        this.avatarURL = newURL;
    }
    
    public void setIP(String gIP){
    	IP = gIP;
    }
    
    public String getIP(){
    	return IP;
    }
}

