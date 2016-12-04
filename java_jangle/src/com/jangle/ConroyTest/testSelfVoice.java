package com.jangle.ConroyTest;

import java.io.IOException;
import java.net.SocketException;
import java.net.UnknownHostException;
import java.util.ArrayList;

import javax.sound.sampled.LineUnavailableException;
import javax.swing.text.html.parser.Parser;

import com.jangle.client.Client;
import com.jangle.client.User;
import com.jangle.communicate.Client_ParseData;
import com.jangle.voice.VoiceChat;

public class testSelfVoice {

	public static void main(String[] args) throws SocketException, LineUnavailableException {

		Client Cl = new Client();
		Client_ParseData Parser = null;
		TestServer server = new TestServer(9090);
		try {
			Parser = new Client_ParseData(Cl, "localhost", 9090);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			//e.printStackTrace();
		}
		
		Cl.getUsersArrayList();
		
		User Nate = TestUtil.newTom();
		Nate.setIP("10.1.40.163");
		Cl.addUser(Nate);
		Nate.setChannelID(1);
		Cl.setCurrentChannelID(1);
		Nate.setIsMuted(false);
		Nate.setVoiceStatus(true);
		//Cl.addUser(Nate);
		
		
		VoiceChat test = new VoiceChat(7800, false, Cl, Parser);
		
		test.connectToVoice();
		test.startBrodcast();

		while (true){
		}

		

		

	}

}
