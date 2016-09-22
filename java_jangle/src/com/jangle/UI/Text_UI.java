package com.jangle.UI;

import com.sun.xml.internal.ws.api.message.ExceptionHasMessage;
import javafx.application.Application;
import javafx.application.Platform;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.TextArea;
import javafx.scene.control.TextField;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;

import java.io.IOException;

import com.jangle.communicate.Client_Communicator;
import com.jangle.communicate.network_Connection;
import com.jangle.test_Client_Server.Client;
import com.jangle.test_Client_Server.Server;

public class Text_UI extends Application {

	private boolean isServer = true;
	
	private Client_Communicator comms;

	private TextArea chatArea = new TextArea();
	// private network_Connection connection = isServer ? createServer() :
	// createClient();

	private Parent createContent() {
		// Setting pref height of UI on .show() call
		chatArea.setPrefHeight(550);

		// Making a network connection that connects to the server
		
		try {
			comms = new Client_Communicator("localhost", 9090);
		} catch (IOException e1) {
			System.out.println("FAILED TO CONNECT TO SERVER");
			e1.printStackTrace();
		}

		TextField messageStage = new TextField();

		// On event listener for submitting entered text in text box
		messageStage.setOnAction(event -> {
			String message = "";// = isServer ? "Server: " : "Client: ";
			message += "stuff" + messageStage.getText();
			chatArea.appendText(message + "\n");

			/*
			 * //Sending a message WALTERS try { connection.send(message); }
			 * catch (Exception e) {
			 * chatArea.appendText("Failed to send message! Error code: " + e +
			 * "\n"); }
			 */
			// Send the string to the server CONROY
			if (comms != null) {
				try {
					System.out.println(message);
					comms.sendToServer(message.getBytes("UTF-8"));
				} catch (Exception e) {
					chatArea.appendText("Failed to send message! Error code: " + e + "\n");
				}
				messageStage.clear();
			}
		});

		VBox root = new VBox(20, chatArea, messageStage);
		root.setPrefSize(600, 600);
		return root;
	}

	public static void main(String[] args) {
		launch(args);
	}

	/*
	@Override
	public void init() throws Exception {
		connection.startConnection();
	}
	*/

	@Override
	public void start(Stage primaryStage) {
		primaryStage.setScene(new Scene(createContent()));
		primaryStage.show();
	}

	/*
	@Override
	public void stop() throws Exception {
		connection.closeConnection();
	}
	

	private Server createServer() {
		return new Server(7878, data -> {
			Platform.runLater(() -> {
				addMessage(data.toString() + "\n");
			});
		});
	}

	private Client createClient() {
		return new Client("127.0.0.1", 7878, data -> {
			Platform.runLater(() -> {
				addMessage(data.toString() + "\n");
			});
		});
		
	}
*/
	public void addMessage(String message) {
		chatArea.appendText(message);
	}
}
// Simple text UI. Needs to get implemented for demos / testing
