package com.jangle.communicate;

import java.io.IOException;
import java.net.UnknownHostException;
import java.nio.ByteBuffer;

import com.jangle.client.*;
import static com.jangle.communicate.Comm_CONSTANTS.*;

public class Client_ParseData implements IPARSER {

	Client Cl;
	Client_Communicator Comm;

	/**
	 * Create a parser object with no Client_Commmunicator attached to it.
	 * 
	 * @param Clie
	 *            The client object this communicator references
	 */
	public Client_ParseData(Client Clie) {
		Cl = Clie;

	}

	/**
	 * Create a parser object with a Client_Communicator attached to it.
	 * 
	 * @param Clie
	 *            The client object this communicator references
	 * @param Host
	 *            The IP address of the server
	 * @param port
	 *            port to communicate though with the server
	 * @throws UnknownHostException
	 * @throws IOException
	 */
	public Client_ParseData(Client Clie, String Host, int port) throws UnknownHostException, IOException {
		Cl = Clie;
		Comm = new Client_Communicator(this, Host, port);
	}

	/**
	 * Send a Message Objects's info to the server, as per the message
	 * specification
	 */
	public void sendMessage(Message mess) throws IOException {

		//when ready to send opcode
		//Comm.sendToServer(generateMessage(mess)); when opcode is ready;
		Comm.sendToServer(mess.getMessageContent().getBytes());
	}

	/**
	 * Generate a message to be sent, based off of message specification
	 * 
	 * @param mess
	 *            the message object to create from
	 * @return a byte array formatted for sending to the server.
	 */
	private byte[] generateMessage(Message mess) {

		byte[] toSend = new byte[1024];
		byte[] channelID = ByteBuffer.allocate(4).putInt(mess.getChannelID()).array();
		byte[] serverID = ByteBuffer.allocate(4).putInt(mess.getServerID()).array();
		byte[] userID = ByteBuffer.allocate(4).putInt(mess.getUserID()).array();

		toSend[0] = MESSAGE_TO_SERVER;

		for (int i = 0; i < 4; i++) {
			toSend[1 + i] = serverID[i];
			toSend[5 + i] = channelID[i];
			toSend[9 + i] = userID[i];
		}

		byte[] tmp = mess.getMessageContent().getBytes();

		for (int i = 0; i < tmp.length; i++) {
			toSend[17 + i] = tmp[i];
		}

		return toSend;
	}

	/**
	 * Figure out what the data that was received is.
	 * 
	 * @param tmp the character array to parse, and figure out what it is
	 */
	public void parseData(String tmp) {

		// get the opcode
		// byte opcode = Data.substring(0, 1).getBytes()[0];

		/*
		 * opcode not implemented in version of go server I was running, force
		 * adding message if (opcode == MESSAGE_FROM_SERVER) { String ServerID =
		 * Data.substring(1, 5); int serverID = Integer.valueOf(ServerID);
		 * 
		 * String ChannelID = Data.substring(5, 9); int channelID =
		 * Integer.valueOf(ChannelID);
		 * 
		 * String UserID = Data.substring(9, 13); int userID =
		 * Integer.valueOf(UserID);
		 * 
		 * String timeStamp = Data.substring(13, 17); String messageContent =
		 * Data.substring(17); Message tmp = new Message(userID, messageContent,
		 * timeStamp, serverID, channelID);
		 * 
		 * Cl.addMessage(tmp); return;
		 * 
		 * }
		 */
		Cl.addMessage(new Message(0, tmp, "PlaceHolder", 0, 0));
		// System.out.println("Unknown opcode " + opcode);
	}

	/*
	 * Parse the data from the communicator, see what is, and move the data on
	 * its way to the UI. Buffer size is 1024 Bytes
	 */

}
