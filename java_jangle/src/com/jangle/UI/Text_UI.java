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

		FXMLLoader loader = new FXMLLoader(getClass().getResource("res/fxml/mainUI.fxml"));
		VBox root = loader.load();

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
