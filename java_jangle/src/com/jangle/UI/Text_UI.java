package com.jangle.UI;

import com.jangle.client.Client;
import com.jangle.communicate.Client_ParseData;
import javafx.application.Application;
import javafx.fxml.FXMLLoader;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.TextArea;
import javafx.scene.control.TextField;
import javafx.scene.layout.AnchorPane;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;

import java.io.IOException;



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

		return loader.<VBox>load();
	}

	private Parent createLoginDialog() throws IOException {
		mClient = new Client();
		mClientParseData = new Client_ParseData(mClient);

		FXMLLoader loader = new FXMLLoader(getClass().getResource("res/fxml/login.fxml"));
		AnchorPane dialog = loader.load();
		loginController mLogin = loader.getController();
		mLogin.setmClient_parseData(mClientParseData);
		mLogin.initializeThread();

		return dialog;
	}

	public static void main(String[] args) {
		launch(args);
	}

	@Override
	public void start(Stage primaryStage) throws IOException {
		Stage loginStage = new Stage();
		loginStage.setScene(new Scene(createLoginDialog()));
		loginStage.showAndWait();

		primaryStage.setScene(new Scene(createContent()));
		primaryStage.show();

	}

}
