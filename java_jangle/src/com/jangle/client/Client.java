package com.jangle.client;

import java.util.ArrayList;
import java.util.List;

/**
 * Created by Jess on 9/28/2016.
 */

public class Client {

    private ArrayList<User> mUsers;
    private ArrayList<Message> mMessages;
    private int serverID;
    private int channelID;
    private int userID;
    private String displayName;


    public Client(ArrayList<User> users, ArrayList<Message> messages, int serverID, int channelID) {
        this.channelID = channelID;
        this.serverID = serverID;
        this.mMessages = messages;
        this.mUsers = users;
    }

    public Client(ArrayList<User> users, ArrayList<Message> messages) {
        this.mUsers = users;
        this.mMessages = messages;
    }

    public Client(int serverID, int channelID) {
        this.serverID = serverID;
        this.channelID = channelID;
        this.mUsers = new ArrayList<>();
        this.mMessages = new ArrayList<>();
    }

    public Client() {
        this.mUsers = new ArrayList<>();
        this.mMessages = new ArrayList<>();
        serverID = 0;
        channelID = 0;
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
}
