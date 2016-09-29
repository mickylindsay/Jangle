package com.jangle.UI;

import javafx.application.Application;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.TextArea;
import javafx.scene.control.TextField;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;

import java.io.IOException;

import com.jangle.communicate.Client_Communicator;


public class Text_UI extends Application {
	
	private Client_Communicator mClientCommunicator;

	private TextArea chatArea = new TextArea();

	private Parent createContent() {
		// Setting pref height of UI on .show() call
		chatArea.setPrefHeight(550);

		// Making a network connection that connects to the server
		
		try {
			mClientCommunicator = new Client_Communicator("localhost", 9090);
		} catch (IOException e1) {
			System.out.println("FAILED TO CONNECT TO SERVER");
			e1.printStackTrace();
		}

		TextField messageStage = new TextField();

		// On event listener for submitting entered text in text box
		messageStage.setOnAction(event -> {
			String message = "";
			message += "stuff" + messageStage.getText();
			chatArea.appendText(message + "\n");

			// Send the string to the server
			if (mClientCommunicator != null) {
				try {
					System.out.println(message);
					mClientCommunicator.sendToServer(message.getBytes("UTF-8"), {0000}, {0000});
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

	@Override
	public void start(Stage primaryStage) {
		primaryStage.setScene(new Scene(createContent()));
		primaryStage.show();
	}

	public void addMessage(String message) {
		chatArea.appendText(message);
	}
}
// Simple text UI. Needs to get implemented for demos / testing
