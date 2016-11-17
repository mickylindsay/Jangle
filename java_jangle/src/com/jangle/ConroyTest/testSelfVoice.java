package com.jangle.ConroyTest;

import java.net.SocketException;

import javax.sound.sampled.LineUnavailableException;

import com.jangle.voice.VoiceChat;

public class testSelfVoice {

	public static void main(String[] args) throws SocketException, LineUnavailableException {

		VoiceChat test = new VoiceChat(7800, TestUtil.genUserList());
		  
		  
		  test.startBrodcast();
		  test.startSpeakers();
		  
		  
		  test.recieveData();

	}

}
