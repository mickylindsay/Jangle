package com.jangle.communicate;

public final class CommUtil {

	/**
	 * private constructor to prevent Instantiation
	 */
	private CommUtil() {
	};
	
	public static final int OPCODE_WIDTH = 1;
	public static final int USERID_WIDTH = 4;
	public static final int ROOMID_WIDTH = 4;
	public static final int SERVERID_WIDTH = 4;
	public static final int TIME_WIDTH = 4;
	
	
	public static final byte CREATE_USER 					= (byte) 0;
	public static final byte CREATE_USER_FAIL 				= (byte) 1;
	public static final byte LOGIN_FAIL 					= (byte) 3;
	public static final byte LOGIN_SUCCESS 					= (byte) 4;
	public static final byte LOGIN 							= (byte) 2;
	public static final byte MESSAGE_TO_SERVER 				= (byte) 16;
	public static final byte MESSAGE_FROM_SERVER 			= (byte) 17;
	
	public static final byte REQUEST_N_MESSAGES 			= (byte) 32;
	public static final byte REQUEST_ALL_USERID 			= (byte) 33;
	public static final byte REQUEST_DISPLAY_NAME 			= (byte) 34;
	public static final byte REQUEST_ALL_SERVERID 			= (byte) 35;
	public static final byte REQUEST_SERVER_DISPLAY_NAME 	= (byte) 36;
	public static final byte REQUEST_ALL_ROOMID 			= (byte) 37;
	public static final byte REQUEST_ROOM_DISPALY 			= (byte) 38;
	
	public static final byte RECIEVE_USERID					= (byte) 48;
	public static final byte RECIEVE_DISPLAY_NAME			= (byte) 49;
	
	public static final long TIME_OUT_MILLI					= 3000;
	

	/**
	 * Convert given byte array into an int. The highest (left most, byte[0]) byte is the LSB, while the lowest (right most (byte.length -1)) byte is the MSB
	 * @param data byte array to convert to int
	 * @return returns the byte array as an int
	 */
	public static int byteToInt(byte[] data) {

		return (unsignByte(data[0]) * 1) + (unsignByte(data[1]) * 256) + (unsignByte(data[2]) * 256 * 256)
				+ (unsignByte(data[3]) * 256 * 256 * 256);
	}

	/**
	 * takes in a byte that is unsigned, and returns it as an int.
	 * @param data byte to return as an int
	 * @return data as an int.
	 */
	public static int unsignByte(byte data) {
		return data & 0xFF;
	}
	
	/**
	 * Enum for login the Results of various things.
	 */
	public enum LoginResult{
		SUCESS, FAIL, TIMEOUT, NAME_TAKEN
	}
	
	public enum UserStatus{
		ONLINE, OFFLINE, AWAY
	}

}
