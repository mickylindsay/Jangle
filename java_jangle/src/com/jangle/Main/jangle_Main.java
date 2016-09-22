package com.jangle.Main;
import com.jangle.*;
import com.jangle.communicate.Client_Communicator;

public class jangle_Main {

	public static void main(String[] args){
		
	/*
	 * Sends "test" to the server. Reads data back from the server, as the server currently echos
	 * what was written to it
	 * 
	 * You can feel free to delete this file, and overwrite this file if need be.
	 */
	Client_Communicator Commun = new Client_Communicator("localhost", 9090);
	
	Commun.sendToServer("test");
	
	System.out.println(Commun.readFromServer());

	}
}
