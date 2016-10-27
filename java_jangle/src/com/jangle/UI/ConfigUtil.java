package com.jangle.UI;

import javafx.application.Platform;

import javax.swing.*;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.Scanner;

/**
 * Created by sable_000 on 10/20/2016.
 */
public class ConfigUtil {

    private String path;
    private boolean isNewConfig;
    private boolean failed;
    private String serverIP;
    private String userName;

    public ConfigUtil() {
        this.failed = false;
        this.path = rootDirectory() + "Jangle\\";
        System.out.print(path);
        File configDir = new File(path);

        if (!configDir.exists()) {
            boolean i = configDir.mkdir();
            System.out.print(i);
        }
        else {
            File configFile = new File(path + "config.cfg");
            if (configFile.exists()){
                readConfig(configFile);
            }
            else {
                boolean created = false;
                try {
                    created = configFile.createNewFile();
                } catch (IOException e) {
                    e.printStackTrace();
                }
                if (created) {
                    makeConfig(configFile);
                }
                else {
                    failed = true;
                }
            }
        }
    }

    private void makeConfig(File configFile) {
        PrintWriter pw;
        try {
            pw = new PrintWriter(configFile);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
            return;
        }
        String ip = (String) JOptionPane.showInputDialog(null, "No Cofig file found please enter the server IP");

        if (isValidIP(ip)){
            this.serverIP = ip;
            printServerIP(pw);
        }
        else {
            failed = true;
        }
        pw.close();
    }

    private void printServerIP(PrintWriter pw) {
        pw.println("[Server IP]");
        pw.println(serverIP);
    }

    private boolean isValidIP(String ip) {
        if (ip.contains("[a-zA-Z]+")) {
            return false;
        }
        else if(!ip.contains(".") || !ip.contains(":")) {
            return false;
        }
        else
            return true;
    }

    private void readConfig(File configFile) {

    }


    private String rootDirectory(){
        return File.listRoots()[0].getAbsolutePath();
    }

    public boolean isNewConfig() {
        return isNewConfig;
    }

    public void setNewConfig(boolean newConfig) {
        isNewConfig = newConfig;
    }

    public boolean isFailed() {
        return failed;
    }

    public void setFailed(boolean failed) {
        this.failed = failed;
    }

    public String getServerIP() {
        return serverIP;
    }

    public void setServerIP(String serverIP) {
        this.serverIP = serverIP;
    }

    public String getUserName() {
        return userName;
    }

    public void setUserName(String userName) {
        this.userName = userName;
    }


    public String[] getFormattedServerIP() {
        return serverIP.split(":");
    }
}

