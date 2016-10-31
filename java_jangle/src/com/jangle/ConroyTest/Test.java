package com.jangle.ConroyTest;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;

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

public class Test {

	public static void main(String[] args) throws IOException, InterruptedException, LineUnavailableException {

		// Client Cl = new Client();
		// Client_ParseData Parse = null;
		// // TestServer server = new TestServer(9090);
		//
		// try {
		// Parse = new Client_ParseData(Cl, "localhost", 9090);
		// System.out.println("generated client");
		// } catch (IOException e1) {
		// // TODO Auto-generated catch block
		// e1.printStackTrace();
		// }
		//
		// Thread.sleep(1000);

		// EDIT BELOW HERE

		// set up the TargetDataLine
		
		UDP_PeerToPeer test2 = new UDP_PeerToPeer();
		
		byte[] data = CommUtil.intToByteArr(1);
		
		test2.sendData(data);
		
		
		DatagramSocket socket = new DatagramSocket();
		
		byte[] buf = new byte[4];
		DatagramPacket packet = new DatagramPacket(buf, buf.length);
		System.out.println("test");
		socket.receive(packet);
		
		System.out.println(CommUtil.byteToInt(data));
		
		
		
		

	}

}
