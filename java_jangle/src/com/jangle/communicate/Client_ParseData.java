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
		Comm.sendToServer(mess.getByteArray());
	}

	
	/**
	 * Figure out what the data that was received is.
	 * 
	 * @param data the character array to parse, and figure out what it is
	 */
	public void parseData(byte[] data) {

		if (data[0] == MESSAGE_FROM_SERVER){
			Cl.addMessage(new Message (data));
		}
	}

	/*
	 * Parse the data from the communicator, see what is, and move the data on
	 * its way to the UI. Buffer size is 1024 Bytes
	 */

}
