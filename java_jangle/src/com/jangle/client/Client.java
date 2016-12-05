package com.jangle.client;

import com.jangle.communicate.CommUtil;
import com.jangle.communicate.CommUtil.*;

import java.util.*;

/**
 * Created by Jess on 9/28/2016.
 */

public class Client {

    private ArrayList<User> mUsers;
    private ArrayList<Message> mMessages;
    private HashMap<Integer, Server> mServers;

    private boolean loggedIn;
    private LoginResult mLoginResult;
    private long mLoginTime;
    private int currentServerID;
    private int currentChannelID;
    private int userID;
    private String displayName;
    private String IP;
    private boolean isBroadcasting;
    private CommUtil.UserStatus status;
    private boolean isMuted;
    private boolean locationChanged;
    private boolean connectedToVoice;
    private boolean statusChanged;



    public Client(ArrayList<User> users, ArrayList<Message> messages, int currentServerID, int currentChannelID) {
        this.currentChannelID = currentChannelID;
        this.currentServerID = currentServerID;
        this.mMessages = messages;
        this.mUsers = users;
        this.userID = 0;
        this.loggedIn = false;
        this.mLoginTime = 0;
        this.IP = "";
        this.mServers = new HashMap<>();
        this.status = CommUtil.UserStatus.ONLINE;
        this.isBroadcasting = false;
        this.connectedToVoice = false;
    }

    public Client(ArrayList<User> users, ArrayList<Message> messages) {
        this.mUsers = users;
        this.mMessages = messages;
        this.userID = 0;
        this.loggedIn = false;
        this.mLoginTime = 0;
        this.IP = "";
        this.mServers = new HashMap<>();
        this.status = CommUtil.UserStatus.ONLINE;
        this.isBroadcasting = false;
        this.connectedToVoice = false;
    }

    public Client(int currentServerID, int currentChannelID) {
        this.currentServerID = currentServerID;
        this.currentChannelID = currentChannelID;
        this.mUsers = new ArrayList<>();
        this.mMessages = new ArrayList<>();
        this.loggedIn = false;
        this.mLoginTime = 0;
        this.IP = "";
        this.mServers = new HashMap<>();
        this.status = CommUtil.UserStatus.ONLINE;
        this.isBroadcasting = false;
        this.connectedToVoice = false;
    }

    public Client() {
        this.mUsers = new ArrayList<>();
        this.mMessages = new ArrayList<>();
        currentServerID = 0;
        currentChannelID = 0;
        this.userID = 0;
        this.loggedIn = false;
        this.mLoginTime = 0;
        this.IP = "";
        this.mServers = new HashMap<>();
        this.status = CommUtil.UserStatus.ONLINE;
        this.isBroadcasting = false;
        this.locationChanged = false;
        this.connectedToVoice = false;
        this.statusChanged = false;
    }

    public void addMessage(Message message, int sId, int chId) {
        if (mServers.get(sId) != null) {
            if(mServers.get(sId).getChannel(chId) != null) {
                mServers.get(sId).getChannel(chId).addMessage(message);
            }
            else {
                System.out.println("Trying to add message to server: " + sId + " in channel: " + chId
                    + " failed. Channel does not exist");
            }
        }
        else {
            System.out.println("Trying to add message to server: " + sId + " in channel: " + chId
                    + " failed. Server does not exist");
        }
        mMessages.add(message);
    }

    public void addUser(User user) {
        mUsers.add(user);
        sortUsers();
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

    //At the moment this is redundant
    public List<Message> getMessageListCustom(int sId, int chID) {
        ArrayList<Message> list = new ArrayList<>();
        for(Message m : mMessages) {
            if (m.getServerID() == sId) {
                if (m.getChannelID() == chID) {
                    list.add(m);
                }
            }
        }
        return list;
    }

    public List<Message> getMessages(int sId, int chId){
        if (mServers.get(sId) != null) {
            if(mServers.get(sId).getChannel(chId) != null) {
                return mServers.get(sId).getChannel(chId).getmMessages();
            }
            else {
                System.out.println("Trying to get messages from server: " + sId + " in channel: " + chId
                        + " failed. Channel does not exist");
            }
        }
        else {
            System.out.println("Trying to get messages from server: " + sId + " in channel: " + chId
                    + " failed. Server does not exist");
        }
        return null;
    }

    public void setMessages(ArrayList<Message> messages) {
        mMessages = messages;
    }

    public int getCurrentServerID() {
        return currentServerID;
    }

    public void setCurrentServerID(int currentServerID) {
        this.currentServerID = currentServerID;
    }

    public int getCurrentChannelID() {
        return currentChannelID;
    }

    public void setCurrentChannelID(int currentChannelID) {
        this.currentChannelID = currentChannelID;
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

    public boolean isLocationChanged() {
        return locationChanged;
    }

    public void setLocationChanged(boolean locationChanged) {
        this.locationChanged = locationChanged;
    }

    public boolean isStatusChanged() {
        return statusChanged;
    }

    public void setStatusChanged(boolean statusChanged) {
        this.statusChanged = statusChanged;
    }

    public User findUser(int id) {

        for (User mUser : mUsers) {
            if (mUser.getId() == id) {
                return mUser;
            }
        }
        //if not found return null
        return null;
    }

    public boolean isDuplicateMessage(Message newMess) {
        for (Message mMessage : mMessages) {
            if (newMess.toString().equals(mMessage.toString()))
                return true;
        }
        return false;
    }

    public HashMap<Integer, Server> getServers() {
        return mServers;
    }

    public void addServer(Server server) {
        if (mServers.get(server.getId()) != null) {
            return;
        }
        mServers.put(server.getId(), server);
    }

    public Server getServer(int id) {
        return mServers.get(id);
    }
    
    public void setIsMuted(boolean status){
    	this.isMuted = status;
    }
    
    public boolean getIsMuted(){
    	return this.isMuted;
    }
    
    public void setBroadcastStatus(boolean status){
    	this.isBroadcasting = status;
    }
    
    public boolean getBroadcastStatus(){
    	return this.isBroadcasting;
    }
    
    public CommUtil.UserStatus getStatus() {
        return status;
    }

    public void setStatus(CommUtil.UserStatus status) {
        this.status = status;
    }

    public void changeChannel(int id) {
        if (currentChannelID == id) {
            return;
        }
        else if (getServer(currentServerID).getChannel(id) == null){
            System.out.println("Trying to switch to channel: " + id +" failed. Channel does not exist");
        }
        else{
            currentChannelID = id;
        }
    }

    public void updateUserPosition(int userID, int newServerID, int newChannelID) {
        findUser(userID).setChannelID(newChannelID);
        sortUsers();
    }

    public void sortUsers(){
        Collections.sort(mUsers, new Comparator<User>() {
            @Override
            public int compare(User u1, User u2) {
                //System.out.println("here boi");
                if (u1.isChannel() && u1.getChannelID() < u2.getChannelID())
                    return -1;

                else if (u1.isChannel() && u1.getChannelID() > u2.getChannelID())
                    return 1;

                else if (u2.isChannel() && u2.getChannelID() < u1.getChannelID())
                    return 1;

                else if (u2.isChannel() && u2.getChannelID() > u1.getChannelID())
                    return -1;

                else if (u1.getChannelID() < u2.getChannelID())
                    return -1;

                else if (u1.getChannelID() == u2.getChannelID())
                    return u1.getDisplayName().compareTo(u2.getDisplayName());
                else
                    return 1;
            }
        });
    }
    
    public boolean isConnectedToVoice(){
    	return this.connectedToVoice;
    }
    
    public void setConnectedToVocie(boolean value){
    	this.connectedToVoice = value;
    }
    
    
}
