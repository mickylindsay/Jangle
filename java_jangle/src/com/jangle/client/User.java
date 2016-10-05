package com.jangle.client;

/**
 * Created by Jess on 9/28/2016.
 */
public class User {

    private int id;
    private int status;
    private String displayName;

    public User(String displayName, int id, int status) {
        this.displayName = displayName;
        this.id = id;
        this.status = status;
    }

    public User(String displayName, int id) {
        this.displayName = displayName;
        this.id = id;
        this.status = 0;
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

    public String toString(){
        return displayName;
    }

}

