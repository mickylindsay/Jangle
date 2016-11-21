package com.jangle.client;

import com.jangle.communicate.CommUtil;
import com.jangle.communicate.CommUtil.*;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Jess on 9/28/2016.
 */

public class Client {

    private ArrayList<User> mUsers;
    private ArrayList<Message> mMessages;

    private boolean loggedIn;
    private LoginResult mLoginResult;
    private long mLoginTime;
    private int serverID;
    private int channelID;
    private int userID;
    private String displayName;
    private String IP;


    public Client(ArrayList<User> users, ArrayList<Message> messages, int serverID, int channelID) {
        this.channelID = channelID;
        this.serverID = serverID;
        this.mMessages = messages;
        this.mUsers = users;
        this.userID = 0;
        this.loggedIn = false;
        this.mLoginTime = 0;
        this.IP = "";
    }

    public Client(ArrayList<User> users, ArrayList<Message> messages) {
        this.mUsers = users;
        this.mMessages = messages;
        this.userID = 0;
        this.loggedIn = false;
        this.mLoginTime = 0;
        this.IP = "";
    }

    public Client(int serverID, int channelID) {
        this.serverID = serverID;
        this.channelID = channelID;
        this.mUsers = new ArrayList<>();
        this.mMessages = new ArrayList<>();
        this.loggedIn = false;
        this.mLoginTime = 0;
        this.IP = "";
    }

    public Client() {
        this.mUsers = new ArrayList<>();
        this.mMessages = new ArrayList<>();
        serverID = 0;
        channelID = 0;
        this.userID = 0;
        this.loggedIn = false;
        this.mLoginTime = 0;
        this.IP = "";
    }

    public void addMessage(Message message) {
        mMessages.add(message);
    }

    public void addUser(User user) {
        mUsers.add(user);
    }

    public void removeUser(User user) {
        if (mUsers.contains(user))
            mUsers.remove(user);
    }

    public void removeMessage(Message message) {
        if (mMessages.contains(message))
            mMessages.remove(message);
    }

    public List<User> getUsers() {
        return mUsers;
    }
    
    public ArrayList<User> getUsersArrayList(){
    	return this.mUsers;
    }

    public void setUsers(ArrayList<User> users) {
        mUsers = users;
    }

    public List<Message> getMessages() {
        return mMessages;
    }

    public void setMessages(ArrayList<Message> messages) {
        mMessages = messages;
    }

    public int getServerID() {
        return serverID;
    }

    public void setServerID(int serverID) {
        this.serverID = serverID;
    }

    public int getChannelID() {
        return channelID;
    }

    public void setChannelID(int channelID) {
        this.channelID = channelID;
    }

    public int getUserID() {
        return userID;
    }

    public void setUserID(int userID) {
        this.userID = userID;
    }

    public String getDisplayName() {
        return displayName;
    }

    public void setDisplayName(String displayName) {
        this.displayName = displayName;
    }

    public boolean isLoggedIn() {
        return loggedIn;
    }

    public void setLoggedIn(boolean loggedIn) {
        this.loggedIn = loggedIn;
    }

    public LoginResult getLoginResult() {
        return mLoginResult;
    }

    public void setLoginResult(LoginResult loginResult) {
        mLoginResult = loginResult;
    }

    public long getLoginTime() {
        return mLoginTime;
    }

    public void setLoginTime(long loginTime) {
        mLoginTime = loginTime;
    }
    
    public String getIP(){
    	return this.IP;
    }
    
    public void setIP(String gIP){
    	this.IP = gIP;
    }

}
