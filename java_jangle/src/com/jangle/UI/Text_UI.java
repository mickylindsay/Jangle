package com.jangle.UI;

import com.jangle.client.Channel;
import com.jangle.client.Client;
import com.jangle.client.Server;
import com.jangle.client.User;
import com.jangle.communicate.Client_ParseData;
import com.jangle.voice.VoiceChat;

import javafx.application.Application;
import javafx.application.Platform;
import javafx.fxml.FXMLLoader;
import javafx.scene.Parent;
import javafx.scene.Scene;

import javafx.scene.layout.AnchorPane;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;

import java.io.IOException;



public class Text_UI extends Application {

	private Client_ParseData mClientParseData;
	private Client mClient;
    private ConfigUtil mConfigUtil;
    private String[] serverIP;
    private VoiceChat mVoice;

	private Parent createContent() throws IOException {

		FXMLLoader loader = new FXMLLoader(getClass().getResource("res/fxml/mainUI.fxml"));
        VBox mainUI = loader.load();
		FXMLController controller = loader.getController();
		controller.setmClientParseData(mClientParseData);
        controller.setConfigUtil(mConfigUtil);
        controller.setVoiceChat(mVoice);

		return mainUI;
	}

	private Parent createLoginDialog() throws IOException {
        serverIP = mConfigUtil.getFormattedServerIP();
		this.mClient = new Client(1, 1);
        Server server = new Server(1);
        Channel channel = new Channel(1);
        mClient.addServer(server);
        mClient.getServer(1).addChannel(channel);
        
        this.mVoice = new VoiceChat(7800, false, mClient, mClientParseData);

		try {
			this.mClientParseData = new Client_ParseData(mClient, serverIP[0], new Integer(serverIP[1]));
		}catch (Exception e) {
			e.printStackTrace();
		}

		FXMLLoader loader = new FXMLLoader(getClass().getResource("res/fxml/login.fxml"));
		AnchorPane dialog = loader.load();
		loginController mLogin = loader.getController();
		mLogin.setmClient_parseData(mClientParseData);
        mLogin.setmConfigUtil(mConfigUtil);

		return dialog;
	}

	public static void main(String[] args) {
		launch(args);
	}

	@Override
	public void start(Stage primaryStage) throws IOException {
        mConfigUtil = new ConfigUtil();

		Stage loginStage = new Stage();
		loginStage.setScene(new Scene(createLoginDialog()));
		loginStage.showAndWait();

        mClientParseData.request50MessagesWithOffset(0);
        mClientParseData.requestAllUsersTiedToServer();
        //TODO: Fix requesting servers
        mClientParseData.requestAllServers(new User("", mClient.getUserID()));
        primaryStage.setScene(new Scene(createContent()));
        primaryStage.show();

	}

}
