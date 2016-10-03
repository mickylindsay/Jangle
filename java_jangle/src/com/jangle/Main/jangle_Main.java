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

		
		byte[] s2 = { (byte) 16, (byte) 0, (byte) 0, (byte) 0, (byte) 0, (byte) 0, (byte) 0, (byte) 0, (byte) 0,
				(byte) 0, (byte) 0, (byte) 0, (byte) 0, (byte) 49, (byte) 50, (byte) 51, (byte) 52, (byte) 116,
				(byte)101, (byte)115, (byte)116 };
		Message test2 = new Message(s2);
		
		System.out.println(test2.getMessageContent());
		
		
		for (int i = 0; i < s.length; i++) {
			System.out.print(s[i] + " ");
		}
		System.out.println();
		System.out.println();
		for (int i = 0; i < s2.length; i++) {
			System.out.print(s2[i] + " ");
		}

	}
}
