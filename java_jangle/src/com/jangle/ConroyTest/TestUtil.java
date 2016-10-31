package com.jangle.ConroyTest;

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
		User ret = new User("Nathan", 2);
		
		return ret;
		
	}
}
