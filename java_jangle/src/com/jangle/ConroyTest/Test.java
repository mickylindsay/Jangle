package com.jangle.ConroyTest;

import java.io.IOException;

import com.jangle.*;
import com.jangle.client.*;
import com.jangle.communicate.Client_ParseData;
import com.jangle.communicate.CommUtil;

public class Test {

	public static void main(String[] args) throws IOException, InterruptedException {

		Client Cl = new Client();
		Client_ParseData Parse = null;
		TestServer server = new TestServer(9090);

		try {
			Parse = new Client_ParseData(Cl, "localhost", 9090);
		} catch (IOException e1) {
			// TODO Auto-generated catch block
			e1.printStackTrace();
		}
		
		Thread.sleep(1000);
		
		// EDIT BELOW HERE

		Parse.setNewDisplayNameOnServer("test");
	}

}
