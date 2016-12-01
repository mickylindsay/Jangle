package com.jangle.client;

import java.io.InputStream;

import com.jangle.communicate.CommUtil;

/**
 * Created by Jess on 9/28/2016.
 */
public class User {

    private final String DEFAULT_AVATAR = "http://res.freestockphotos.biz/pictures/17/17384-illustration-of-a-red-santa-hat-pv.png";
    public final String OFFLINE_AVATAR = "https://d30y9cdsu7xlg0.cloudfront.net/png/215345-200.png";

    private int id;
    private CommUtil.UserStatus status;
    private boolean isMuted;
    private String displayName;
    private String userName;
    private String avatarURL;
    private String IP;
    private int channelID;
    private boolean voice;
    private boolean isChannel;
    

    public User(String displayName, int id, int status) {
        this.displayName = displayName;
        this.id = id;
        this.status = CommUtil.UserStatus.ONLINE;
        this.userName = "";
        this.avatarURL = DEFAULT_AVATAR;
        this.IP = "";
        this.channelID = 0;
        this.isMuted = false;
        this.voice = false;
        this.isChannel = false;
    }
    
    public User(String displayName, String userName, int id, int status){
    	this.displayName = displayName;
        this.id = id;
        this.status = CommUtil.UserStatus.OFFLINE;
        this.userName = userName;
        this.avatarURL = DEFAULT_AVATAR;
        this.IP = "";
        this.channelID = 0;
        this.isMuted = false;
        this.voice = false;
        this.isChannel = false;
    }

    public User(String displayName, int id) {
        this.displayName = displayName;
        this.id = id;
        this.status = CommUtil.UserStatus.OFFLINE;
        this.userName = "";
        this.avatarURL = DEFAULT_AVATAR;
        this.IP = "";
        this.channelID = 0;
        this.isMuted = false;
        this.voice = false;
        this.isChannel = false;
    }

    public User(Channel channel){
        this.displayName = channel.toString();
        this.id = 1000 + channel.getId();
        this.userName = channel.getName();
        this.status = CommUtil.UserStatus.ONLINE;
        this.IP = "";
        this.channelID = channel.getId();
        this.isChannel = true;
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

    public CommUtil.UserStatus getStatus() {
        return status;
    }

    public void setStatus(CommUtil.UserStatus status) {
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
    
    public void setChannelID(int id){
    	this.channelID = id;
    }
    
    public int getChannelID(){
    	return this.channelID;
    }

    public String getAvatarURL() {
        return avatarURL;
    }

    public void setAvatar(String newURL) {
        this.avatarURL = newURL;
    }
    
    public void setIP(String gIP){
    	this.IP = gIP;
    }
    
    public String getIP(){
    	return this.IP;
    }
    
    public void setIsMuted(boolean status){
    	this.isMuted = status;
    }
    
    public boolean getIsMuted(){
    	return this.isMuted;
    }
    
    public void setVoiceStatus(boolean status){
    	this.voice = status;
    }
    
    public boolean getVoiceStatus(){
    	return this.voice;
    }

    public boolean isChannel() {
        return isChannel;
    }
}

