package com.jangle.voice;

import java.util.Scanner;

import javax.sound.sampled.AudioFormat;

public final class VoiceUtil {

	//private constructor to prevent instantation
	private VoiceUtil(){
		
	}
	
	//Constants
	
	public static final int VOICE_DATA_BUFFER_SIZE = 1024;
	public static final int SLEEP_MILLI_LENGTH = 50;
	
	//The audio format used in voice chat. It is passed from voice chat to the other classes.
	public static AudioFormat genFormat(){
		return  new AudioFormat(8000.0f, 16, 1, true, true);
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
}
