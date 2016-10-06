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
	
	public static final byte MESSAGE_TO_SERVER = (byte) 16;
	public static final byte MESSAGE_FROM_SERVER = (byte) 17;
	

	/**
	 * Convert given byte array into an int. The highest (left mos)t byte is the LSB, while the lowest (right most) byte is the MSB
	 * @param data byte array to convert to int
	 * @return returns the byte array as an int
	 */
	public static int byteToInt(byte[] data) {

		return (unsignByte(data[3]) * 1) + (unsignByte(data[2]) * 256) + (unsignByte(data[1]) * 256 * 256)
				+ (unsignByte(data[0]) * 256 * 256 * 256);
	}

	/**
	 * takes in a byte that is unsigned, and returns it as an int.
	 * @param data byte to return as an int
	 * @return data as an int.
	 */
	public static int unsignByte(byte data) {
		return data & 0xFF;
	}
	
	public static enum LoginResult{
		SUCESS, FAIL, TIMEOUT
	}

}
