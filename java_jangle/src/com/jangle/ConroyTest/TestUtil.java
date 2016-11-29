package com.jangle.ConroyTest;

import java.util.ArrayList;

import com.jangle.client.Message;
import com.jangle.client.User;

public final class TestUtil {

	/**
	 * private constructor to prevent Instantiation
	 */
	private TestUtil() {
	};

	public static byte[] TEST_MESS_FROM_SERVER = { (byte) 17, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, (byte) 'T',
			(byte) 'E', (byte) 'S', (byte) 'T' };
	
	public  static Message genTest(){
		Message ret = new Message();
		ret.setChannelID(1);
		ret.setServerID(1);
		ret.setTimeStamp(1);
		ret.setUserID(1);
		ret.setMessageContent("Test123456789Test");
		
		return ret;
		
	}
	
	public static User newNathan(){
		User ret = new User("nathan", 5);
		return ret;
	}
	
	public static User newMicky(){
		User ret = new User("micky", 2);
		return ret;
	}
	
	public static User newJess(){
		User ret = new User("jess", 4);
		return ret;
	}
	
	public static User newTom(){
		User ret = new User("tom", 3);
		return ret;
		
	}

	public static ArrayList<User> genUserList(){
		ArrayList<User> ret = new ArrayList<User>();
		
		ret.add(newNathan());
		ret.add(newMicky());
		ret.add(newTom());
		ret.add(newJess());
		
		return ret;
		
	}
}
