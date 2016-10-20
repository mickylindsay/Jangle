package com.jangle.ConroyTest;

import java.io.IOException;

import com.jangle.*;
import com.jangle.client.*;
import com.jangle.communicate.Client_ParseData;
import com.jangle.communicate.CommUtil;

public class Test {

	public static void main(String[] args) {

	
		
		

		 Client Cl = new Client();
		 Client_ParseData Parse = null;
		// TestServer server = new TestServer(9090);
		
		 try {
		 Parse = new Client_ParseData(Cl, "10.25.72.96", 9090);
		 } catch (IOException e1) {
		 // TODO Auto-generated catch block
		 e1.printStackTrace();
		 }
		
		//EDIT BELOW HERE
		 
		 
		try {
			Parse.request50MessagesWithOffset(0);
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
		
		for (int i = 0; i < Cl.getMessages().size(); i++){
			System.out.println(Cl.getMessages().get(i).getMessageContent());
		}
		
		

	}

}
