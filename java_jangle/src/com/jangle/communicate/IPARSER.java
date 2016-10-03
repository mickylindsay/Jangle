package com.jangle.communicate;
import java.io.IOException;
import java.net.UnknownHostException;

import com.jangle.client.Message;

public interface IPARSER {
	
	public void sendMessage(Message mess) throws IOException;
	
	

}
