package com.jangle.ConroyTest;

import java.io.IOException;
import java.net.SocketException;
import java.net.UnknownHostException;
import java.util.ArrayList;

import javax.sound.sampled.LineUnavailableException;

import com.jangle.client.Client;
import com.jangle.client.User;
import com.jangle.voice.VoiceChat;

public class testSelfVoice {

	public static void main(String[] args) throws SocketException, LineUnavailableException {

		Client Cl = new Client();
		
		Cl.getUsersArrayList();
		
		User Nate = TestUtil.newNathan();
		Nate.setIP("localhost");
		Cl.addUser(Nate);
		Nate.setChannelID(1);
		Cl.setChannelID(1);


		/*
		VoiceChat test = new VoiceChat(7800, false, Cl);
		
		test.connectToVoice();
		test.startBrodcast();
		*/


		while (true){
			
		}

		

		

	}

}
