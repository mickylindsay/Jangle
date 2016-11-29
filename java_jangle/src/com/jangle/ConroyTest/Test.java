package com.jangle.ConroyTest;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;

import javax.annotation.processing.SupportedSourceVersion;
import javax.sound.sampled.AudioFormat;

import javax.sound.sampled.AudioSystem;
import javax.sound.sampled.DataLine;
import javax.sound.sampled.LineUnavailableException;
import javax.sound.sampled.SourceDataLine;
import javax.sound.sampled.TargetDataLine;

import javax.swing.text.html.HTMLEditorKit.Parser;

import com.jangle.*;
import com.jangle.client.*;
import com.jangle.communicate.Client_ParseData;
import com.jangle.communicate.CommUtil;
import com.jangle.voice.Microphone;
import com.jangle.voice.VoiceChat;

public class Test {

	public static void main(String[] args) throws IOException, InterruptedException, LineUnavailableException {

		Client Cl = new Client();
		Client_ParseData Parser = null;
		//TestServer server = new TestServer(9090);

		try {
			Parser = new Client_ParseData(Cl, "10.25.70.96", 9090);
			//System.out.println("generated client");
		} catch (IOException e1) {
			// TODO Auto-generated catch block
			e1.printStackTrace();
		}
		
		
		User Nate = TestUtil.newNathan();
		User Jess = TestUtil.newJess();
		
		Jess.setVoiceStatus(true);
		Jess.setIsMuted(false);
		
		Parser.submitLogIn("nathan", "password");
		System.out.println(Parser.getUserIP(Nate));
		//Parser.submitLogIn("jess", "password");
		
		VoiceChat voice = new VoiceChat(7800, false, Cl, Parser);
		
		
		
		
		voice.connectToVoice();
		voice.startBrodcast();
		
		System.out.println(Parser.getUserIP(Jess));
		
		Cl.addUser(Jess);
		Jess.setVoiceStatus(true);
		Jess.setIsMuted(false);
		
	  while(true){ 
		  Thread.sleep(500);
	  }
		  
		 

	}

}
