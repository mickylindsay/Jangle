package com.jangle.ConroyTest;

import java.io.IOException;

import com.jangle.*;
import com.jangle.client.*;
import com.jangle.communicate.Client_ParseData;
import com.jangle.communicate.CommUtil;


public class Test {

	public static void main(String[] args) {
		// TODO Auto-generated method stub

		
//		Client Cl = new Client();
//		Client_ParseData Parse = null;
//		TestServer server = new TestServer(9090);
//		
//		try {
//			Parse = new Client_ParseData(Cl, "localhost", 9090);
//		} catch (IOException e1) {
//			// TODO Auto-generated catch block
//			e1.printStackTrace();
//		}
		
		
		byte[] test = new byte[2004];
		
		test[0] = (byte) 20;
		test[1] = (byte) 0;
		test[2] = (byte) 0;
		test[3] = (byte) 0;
		
		System.out.println(CommUtil.byteToInt(test));
		
		
	}

}
