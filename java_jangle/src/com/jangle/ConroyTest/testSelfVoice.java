package com.jangle.ConroyTest;

import java.io.IOException;
import java.net.SocketException;
import java.net.UnknownHostException;

import javax.sound.sampled.LineUnavailableException;

import com.jangle.client.User;
import com.jangle.voice.VoiceChat;

public class testSelfVoice {

	public static void main(String[] args) throws SocketException, LineUnavailableException {

		VoiceChat test = new VoiceChat(7800, TestUtil.genUserList());
		  
		  User Nate = TestUtil.newNathan();
		  
		  Nate.setIP("localhost");
		  
		  try {
			test.addUserToChat(Nate);
		} catch (UnknownHostException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		  test.startBrodcast();
		  test.startSpeakers();
		  
		  
		  test.recieveData();

	}

}
