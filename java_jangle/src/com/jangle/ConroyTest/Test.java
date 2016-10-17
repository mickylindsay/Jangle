package com.jangle.ConroyTest;

import java.io.IOException;

import com.jangle.*;
import com.jangle.client.*;
import com.jangle.communicate.Client_ParseData;


public class Test {

	public static void main(String[] args) {
		// TODO Auto-generated method stub

		
		Client Cl = new Client();
		Client_ParseData Parse = new Client_ParseData(Cl);
		
		
		try {
			Parse.submitLogIn("Test", "123456789");
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		
		System.out.println("done");
		
	}

}
