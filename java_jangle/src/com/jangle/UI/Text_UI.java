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
import com.jangle.communicate.Client_Communicator;
import com.jangle.communicate.network_Connection;
import com.jangle.test_Client_Server.Client;
import com.jangle.test_Client_Server.Server;


public class Text_UI extends Application {

    private boolean isServer = true;

    private TextArea chatArea = new TextArea();
    private network_Connection connection = isServer ? createServer() : createClient();

    private Parent createContent() {
        //Setting pref height of UI on .show() call
        chatArea.setPrefHeight(550);

        //Making a network connection that connects to the server
        //Client_Communicator comms = new Client_Communicator("localhost", 9090);

        TextField messageStage = new TextField();

        //On event listener for submitting entered text in text box
        messageStage.setOnAction(event -> {
            String message = isServer ? "Server: " : "Client: ";
            message += messageStage.getText();
            chatArea.appendText(message + "\n");

            //Sending a message WALTERS
            try {
                connection.send(message);
            } catch (Exception e) {
                chatArea.appendText("Failed to send message! Error code: " + e + "\n");
            }

            //Send the string to the server CONROY
            /*try {
                comms.sendToServer(message);
            }
            catch (Exception e) {
                chatArea.appendText("Failed to send message! Error code: " + e + "\n");
            } */
            messageStage.clear();
        });

        VBox root = new VBox(20, chatArea, messageStage);
        root.setPrefSize(600, 600);
        return root;
    }

    public static void main(String[] args) {
        launch(args);
    }

    @Override
    public void init() throws Exception {
        connection.startConnection();
    }

    @Override
    public void start(Stage primaryStage) {
        primaryStage.setScene(new Scene(createContent()));
        primaryStage.show();
    }

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

    private Client createClient(){
        return new Client("127.0.0.1", 7878, data -> {
            Platform.runLater(() -> {
                addMessage(data.toString() + "\n");
            });
        });
    }

    public void addMessage(String message){
        chatArea.appendText(message);
    }
}
	//Simple text UI. Needs to get implemented for demos / testing

