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

	public static final byte CREATE_USER = (byte) 0;
	public static final byte CREATE_USER_FAIL = (byte) 1;
	public static final byte LOGIN_FAIL = (byte) 3;
	public static final byte LOGIN_SUCCESS = (byte) 4;
	public static final byte LOGIN = (byte) 2;
	public static final byte MESSAGE_TO_SERVER = (byte) 16;
	public static final byte MESSAGE_FROM_SERVER = (byte) 17;

	public static final byte REQUEST_N_MESSAGES = (byte) 32;
	public static final byte REQUEST_ALL_USERID = (byte) 33;
	public static final byte REQUEST_DISPLAY_NAME = (byte) 34;
	public static final byte REQUEST_ALL_SERVERID = (byte) 35;
	public static final byte REQUEST_SERVER_DISPLAY_NAME = (byte) 36;
	public static final byte REQUEST_ALL_ROOM_ID = (byte) 37;
	public static final byte REQUEST_ROOM_DISPALY_NAME = (byte) 38;
	public static final byte REQUEST_USER_STATUS = (byte) 40;
    public static final byte REQUEST_USER_LOCATION = (byte) 41;
	public static final byte REQUEST_USER_IP = (byte) 43;
    public static final byte REQUEST_USER_ICON = (byte) 44;
    public static final byte REQUEST_SERVER_ICON = (byte) 45;
	public static final byte RECIEVE_USERID = (byte) 48;
	public static final byte RECIEVE_DISPLAY_NAME = (byte) 49;

	public static final byte RECIEVE_SERVER_ID = (byte) 50;
	public static final byte RECIEVE_SERVER_DISPLAY_NAME = (byte) 51;
	public static final byte RECIEVE_ROOM_ID = (byte) 52;
	public static final byte RECIEVE_ROOM_DISPLAY_NAME = (byte) 53;
	public static final byte RECIEVE_USER_STATUS = (byte) 55;
    public static final byte RECIEVE_USER_LOCATION = (byte) 56;
	public static final byte RECIEVE_USER_IP = (byte) 57;
    public static final byte RECIEVE_SERVER_ICON = (byte) 59;
    public static final byte RECIEVE_USER_ICON = (byte) 58;

	public static final byte SEND_NEW_DISPLAY_NAME = (byte) 64;
    public static final byte SEND_NEW_USER_ICON = (byte) 68;

	public static final byte SEND_STAUTS_CHANGE = (byte) 80;
	public static final byte SEND_ROOM_LOCATION_CHANGE = (byte) 81;

	public static final byte RECIEVE_ERROR = (byte) 255;
	
	
	public static final long TIME_OUT_MILLI = 3000;

	/**
	 * Convert given byte array into an int. The highest (left most, byte[0])
	 * byte is the LSB, while the lowest (3 most (byte.length -1)) byte is
	 * the MSB
	 * 
	 * @param data
	 *            byte array to convert to int
	 * @return returns the byte array as an int
	 */
	public static int byteToInt(byte[] data) {

		return (unsignByte(data[0]) * 1) + (unsignByte(data[1]) * 256) + (unsignByte(data[2]) * 256 * 256)
				+ (unsignByte(data[3]) * 256 * 256 * 256);
	}

	/**
	 * takes in a byte that is unsigned, and returns it as an int.
	 * 
	 * @param data
	 *            byte to return as an int
	 * @return data as an int.
	 */
	public static int unsignByte(byte data) {
		return data & 0xFF;
	}

	/**
	 * Converts the input to a 4 byte array. Data is negative, or bigger than
	 * what 4 bytes can represent, it does not check that
	 * 
	 * @param data int to be converted
	 * @return byte array, with LSB at byte[0]
	 */
	public static byte[] intToByteArr(int data) {
		byte[] ret = new byte[4];

		for (int i = 0; i < ret.length; i++){
			ret[i] = (byte) (data % 256);
			data = data / 256;
		}
		return ret;
	}

	/**
	 * Enum for login the Results of various things.
	 */
	public enum LoginResult {
		SUCESS, FAIL, TIMEOUT, NAME_TAKEN
	}

	public enum UserStatus {
		ONLINE, OFFLINE, AWAY
	}
	
	public enum UserRequestResult{
		SUCESS, TIMEOUT
	}
	

}
