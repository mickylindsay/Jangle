package com.jangle.Main;

import java.io.IOException;
import java.util.ArrayList;

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

		
		Message test = new Message(0, "test", "1234", 0, 0);
		
		byte[] s = test.getByteArray();
		
		for (int i = 0; i < s.length; i++){
			System.out.print(s[i] + " ");
		}
		
		
		
	}
}
