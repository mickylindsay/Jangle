package com.jangle.Main;

import java.io.IOException;

import com.jangle.*;
import com.jangle.client.Client;
import com.jangle.client.Message;
import com.jangle.communicate.*;

public class jangle_Main {

	public static void main(String[] args) {

		/*
		 * Sends "test" to the server. Reads data back from the server, as the
		 * server currently echos what was written to it
		 * 
		 * You can feel free to delete this file, and overwrite this file if
		 * need be.
		 */

		
		byte T = (byte)10;
		System.out.println(T);
		
		Client test = new Client();
		Client_ParseData Parse = null;
		try {
			Parse = new Client_ParseData(test, "localhost", 9090);
		} catch (IOException e1) {
			// TODO Auto-generated catch block
			System.out.println("No server for you");
			e1.printStackTrace();
		}

		Message mess = new Message();
		mess.setServerID(10);
		mess.setMessageContent("test");

		try {
			Parse.sendMessage(mess);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}

		try {
			Thread.sleep(1000);
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		System.out.println(test.getMessages().size());
		
		Message mess2 = new Message();
		mess2.setMessageContent("test2");

		try {
			Parse.sendMessage(mess2);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}

		try {
			Thread.sleep(1000);
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
		System.out.println(test.getMessages().size());
		
		System.out.println("END");
	}
}
