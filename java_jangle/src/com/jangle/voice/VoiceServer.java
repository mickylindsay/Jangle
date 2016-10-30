package com.jangle.voice;

import java.net.Socket;
import java.util.ArrayList;


public class VoiceServer implements Runnable {
	
	private int numChatWith;
	private ArrayList<Socket> connections;
	
	public VoiceServer(int port){
		numChatWith = 0;
		connections = new ArrayList<Socket>();
		
		
		
		
	}

	@Override
	public void run() {
		// TODO Auto-generated method stub
		
	}

}
