package com.jangle.voice;

import javax.sound.sampled.AudioFormat;

public final class VoiceUtil {

	//private constructor to prevent instantation
	private VoiceUtil(){
		
	}
	
	//Constants
	
	public static final int VOICE_DATA_BUFFER_SIZE = 1024;
	
	//The audio format used in voice chat. It is passed from voice chat to the other classes.
	public static AudioFormat genFormat(){
		return  new AudioFormat(8000.0f, 16, 1, true, true);
	}
}
