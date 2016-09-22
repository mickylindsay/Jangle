package com.jangle.UI;

import javafx.application.Application;
import javafx.scene.Parent;
import javafx.scene.Scene;
import javafx.scene.control.TextArea;
import javafx.scene.control.TextField;
import javafx.scene.layout.VBox;
import javafx.stage.Stage;
import com.jangle.communicate.Client_Communicator;


public class Text_UI extends Application {


    private TextArea chatArea = new TextArea();

    private Parent createContent() {
        //Setting pref height of UI on .show() call
        chatArea.setPrefHeight(550);

        //Making a network connection that connects to the server
        Client_Communicator comms = new Client_Communicator("localhost", 9090);

        TextField messageStage = new TextField();

        //On event listener for submitting entered text in text box
        messageStage.setOnAction(event -> {
            String message = "Me: ";
            message += messageStage.getText();
            chatArea.appendText(message + "\n");
            //Send the string to the server
            comms.sendToServer(message);
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
    public void start(Stage primaryStage) {
        primaryStage.setScene(new Scene(createContent()));
        primaryStage.show();
    }

    public void addMessage(String message){
        chatArea.appendText(message);
    }
}
	//Simple text UI. Needs to get implemented for demos / testing

