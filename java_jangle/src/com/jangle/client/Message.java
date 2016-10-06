package com.jangle.client;

import java.nio.ByteBuffer;
import java.util.Arrays;

import static com.jangle.communicate.Comm_CONSTANTS.*;

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

	public Message(int userID, String messageContent, int serverID, int channelID) {
		this.userID = userID;
		this.messageContent = messageContent;
		this.serverID = serverID;
		this.channelID = channelID;
		this.timeStamp = 0;
	}

	/**
	 * Generate a Message object from a byte array received from the server.
	 * This message is received in the form of a 17 opcode
	 * 
	 * @param data
	 *            the byte array received from the object. The byte array still
	 *            as the opcode in it
	 */
	public Message(byte[] data) {
		byte[] chan = new byte[4];
		byte[] user = new byte[4];
		byte[] server = new byte[4];
		byte[] time = new byte[4];
		byte[] content = new byte[data.length - 17];

		for (int i = 0; i < 4; i++) {
			server[3 - i] = data[i + 1];
			chan[3 - i] = data[i + 5];
			user[3 - i] = data[i + 9];
			time[3 - i] = data[i + 13];
		}
		
		content = Arrays.copyOfRange(data, 17, data.length);

		this.userID = byteToInt(user);
		this.channelID = byteToInt(chan);
		this.serverID = byteToInt(server);
		this.timeStamp = (long)byteToInt(time);
		this.messageContent = new String(content);
		
		
		
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

	public String toString() {
		return userID + "\n" + messageContent + "    " + timeStamp;
	}

	/**
	 * Creates a byte array of this message, in the format required to send the
	 * message to the server
	 * 
	 * @return the byte array to send to the server.
	 */
	public byte[] getByteArray() {

		byte[] ret = new byte[messageContent.length() + 13];
		byte[] serverid = ByteBuffer.allocate(4).putInt(serverID).array();
		byte[] channelid = ByteBuffer.allocate(4).putInt(channelID).array();
		byte[] userid = ByteBuffer.allocate(4).putInt(userID).array();
		int j = 0;

		ret[0] = MESSAGE_TO_SERVER;

		for (int i = 0; i < 4; i++) {
			if (serverid.length > i) {
				ret[i + 1] = serverid[3 - i];
			}
			else {
				ret[i + 1] = (byte) 0;
			}

			if (channelid.length > i) {
				ret[i + 5] = channelid[3 - i];
			}
			else {
				ret[i + 5] = (byte) 0;
			}
			if (userid.length > i) {
				ret[i + 9] = userid[3 - i];
			}
			else {
				ret[i + 9] = (byte) 0;
			}

		}

		for (int i = 0; i < messageContent.length(); i++) {
			ret[i + 13] = (byte) messageContent.charAt(i);
		}

		return ret;

	}
	
	private int byteToInt(byte[] data){
		
		return (unsignByte(data[3]) * 1) + (unsignByte(data[2]) * 256) + (unsignByte(data[1]) * 256 * 256)
				+ (unsignByte(data[0]) * 256 * 256 * 256);
	}
	
	
	
	private int unsignByte(byte data){
		return data & 0xFF;
	}

}