package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.client.Message;
import com.jangle.communicate.Client_ParseData;
import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.fxml.FXML;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.TextArea;
import javafx.scene.control.TextField;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;

import java.io.IOException;
import java.net.URL;
import java.text.SimpleDateFormat;
import java.util.Date;

import com.jangle.communicate.Client_Communicator;


public class Text_UI extends Application {

	private Client_ParseData mClientParseData;
	private messageThread messageThread;
	private userThread userThread;
	private Client mClient;
	private FXMLController mFXMLController;

	public TextArea chatArea = new TextArea();
	public TextField messageStage = new TextField();


	private Parent createContent() throws IOException {
		// Setting pref height of UI on .show() call
		//chatArea.setPrefHeight(550);
		//chatArea.setEditable(false);


		//Loading FXML main UI
		FXMLLoader loader = new FXMLLoader(getClass().getResource("res/fxml/mainUI.fxml"));
		VBox root = loader.load();

		// Making a network connection that connects to the server
		//mClient = new Client();

		/*try {
			mClientParseData = new Client_ParseData(mClient, "localhost", 9090);
		} catch (IOException e) {
			e.printStackTrace();
		} */

		// On event listener for submitting entered text in text box
		/*messageStage.setOnAction(event -> {
			String message = messageStage.getText();
			// Send the string to the server
			try {
				mClientParseData.sendMessage(new Message(0, message, System.currentTimeMillis(), 0, 0));
			} catch (IOException e) {
				e.printStackTrace();
			}
			messageStage.clear();
		}); */

		//VBox root = new VBox(20, chatArea, messageStage);
		//root.setPrefSize(600, 600);
		//messageThread = new messageThread(mClient, this);
		//userThread = new userThread(mClient, this);

		return root;
	}

	public static void main(String[] args) {
		launch(args);
	}

	@Override
	public void start(Stage primaryStage) throws IOException {
		primaryStage.setScene(new Scene(createContent()));
		primaryStage.show();
	}

}
