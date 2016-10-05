package com.jangle.client;

/**
 * Created by Jess on 9/28/2016.
 */
public class Message {

    private int userID;
    private String messageContent;
    private long timeStamp;
    private int serverID;
    private int channelID;

    public Message(int userID, String messageContent, long timeStamp, int serverID, int channelID) {
        this.channelID = channelID;
        this.userID = userID;
        this.messageContent = messageContent;
        this.timeStamp = timeStamp;
        this.serverID = serverID;
    }

    public Message() {
        this.channelID = 0;
        this.userID = 0;
        this.messageContent = null;
        this.timeStamp = 0;
        this.serverID = 0;
    }

    public int getUserID() {
        return userID;
    }

    public void setUserID(int userID) {
        this.userID = userID;
    }

    public String getMessageContent() {
        return messageContent;
    }

    public void setMessageContent(String messageContent) {
        this.messageContent = messageContent;
    }

    public long getTimeStamp() {
        return timeStamp;
    }

    public void setTimeStamp(long timeStamp) {
        this.timeStamp = timeStamp;
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



}
