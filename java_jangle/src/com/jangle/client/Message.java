package com.jangle.client;

import java.util.ArrayList;
import static com.jangle.communicate.Comm_CONSTANTS.*;

/**
 * Created by Jess on 9/28/2016.
 */
public class Message {

	private int userID;
	private String messageContent;
	private String timeStamp;
	private int serverID;
	private int channelID;

	public Message(int userID, String messageContent, String timeStamp, int serverID, int channelID) {
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
		this.timeStamp = null;
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

	public String getTimeStamp() {
		return timeStamp;
	}

	public void setTimeStamp(String timeStamp) {
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

	/**
	 * Creates a byte array of this message, in the format required to send the
	 * message to the server
	 * 
	 * @return the byte array to send to the server.
	 */
	public byte[] getByteArray() {

		byte[] ret = new byte[messageContent.length() + 13];
		byte[] serverid = String.valueOf(serverID).getBytes();
		byte[] channelid = String.valueOf(channelID).getBytes();
		byte[] userid = String.valueOf(userID).getBytes();
		int j = 0;
		
		ret[0] = MESSAGE_TO_SERVER;

		j = 3;
		for (int i = 0; i < 4; i++) {
			if (serverid.length > i) {
				ret[j + 1] = serverid[i];
			}
			else {
				ret[j + 1] = (byte) 0;
			}

			if (channelid.length > i) {
				ret[j + 5] = channelid[i];
			}
			else {
				ret[j + 5] = (byte) 0;
			}
			if (userid.length > i) {
				ret[j + 9] = userid[i];
			}
			else {
				ret[j + 9] = (byte) 0;
			}
			j--;
			
		}

		for (int i = 0; i < messageContent.length(); i++) {
			ret[i + 13] = (byte) messageContent.charAt(i);
		}

		return ret;

	}

}
