package com.jangle.voice;

import java.util.Scanner;

import javax.sound.sampled.AudioFormat;

public final class VoiceUtil {

	//private constructor to prevent instantation
	private VoiceUtil(){
		
	}
	
	//Constants
	
	public static final int VOICE_DATA_BUFFER_SIZE = 256;
	public static final int VOICE_DATA_SIZE = 2048;
	public static final int SLEEP_MILLI_LENGTH = 50;
	
	//The audio format used in voice chat. It is passed from voice chat to the other classes.
	public static AudioFormat genFormat(){
		return  new AudioFormat(16000.0f, 16, 1, true, false);
	}
	
	public static byte[] byteIP(String IP){
		String tmp = IP.concat(".");
		
		Scanner scan = new Scanner(tmp);
		scan.useDelimiter("\\.");
		byte[] ret = new byte[4];
				
				
		ret[0] = (byte)scan.nextInt();
		ret[1] = (byte)scan.nextInt();
		ret[2] = (byte)scan.nextInt();
		ret[3] = (byte)scan.nextInt();
		
		scan.close();
		
		return ret;
	}
	
	public static byte[] intToByteArr(int data) {
		byte[] ret = new byte[4];

		for (int i = 0; i < ret.length; i++){
			ret[i] = (byte) (data % 256);
			data = data / 256;
		}
		return ret;
	}
	
	public static int byteToInt(byte[] data) {

		return (unsignByte(data[0]) * 1) + (unsignByte(data[1]) * 256) + (unsignByte(data[2]) * 256 * 256)
				+ (unsignByte(data[3]) * 256 * 256 * 256);
	}
	
	public static int unsignByte(byte data) {
		return data & 0xFF;
	}
}
