package com.jangle.client;

/**
 * Created by Jess on 9/28/2016.
 */
public class User {

    private int id;
    private int status;
    private String displayName;
    private String userName;

    public User(String displayName, int id, int status) {
        this.displayName = displayName;
        this.id = id;
        this.status = status;
        this.userName = "";
    }
    
    public User(String displayName, String userName, int id, int status){
    	this.displayName = displayName;
        this.id = id;
        this.status = status;
        this.userName = userName;
    }

    public User(String displayName, int id) {
        this.displayName = displayName;
        this.id = id;
        this.status = 0;
        this.userName = "";
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

}

