package com.jangle.UI;

import javafx.application.Platform;

import javax.swing.*;
import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.PrintWriter;
import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;
import java.util.Scanner;

/**
 * Created by sable_000 on 10/20/2016.
 */
public class ConfigUtil {

    private final String SERVER_IP = "[Server IP]";
    private final String BACKGROUND_PATH = "[Background Path]";
    private final String REMEMBER_USER = "[Remember User]";

    private String path;
    private File configFile;
    private boolean isNewConfig;
    private boolean failed;
    private HashMap<String, String> configValues;

    public ConfigUtil() {
        this.failed = false;
        this.path = rootDirectory() + "Jangle\\";
        System.out.print(path + "\n");
        File configDir = new File(path);
        configValues = new HashMap<>();

        if (!configDir.exists()) {
            boolean i = configDir.mkdir();
            System.out.print(i);
        }
        else {
            configFile = new File(path + "config.cfg");
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
        String ip = (String) JOptionPane.showInputDialog(null, "No Cofig file found please enter the server IP");

        if (isValidIP(ip)){
            setServerIP(ip);
            printConfig();
        }
        else {
            failed = true;
        }
    }

    private void printConfig() {
        PrintWriter pw = null;
        try {
            pw = new PrintWriter(configFile);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
        if (pw != null) {
            for (Object o : configValues.entrySet()) {
                Map.Entry me = (Map.Entry) o;
                pw.println(me.getKey());
                pw.println(me.getValue());
            }
            pw.close();
        }
        else
            failed = true;
    }

    private boolean isValidIP(String ip) {
        if (ip.equals("localhost:9090"))
            return true;

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
        Scanner configScanner = null;
        try {
            configScanner = new Scanner(configFile);
        } catch (FileNotFoundException e) {
            e.printStackTrace();
            return;
        }

        configScanner.useDelimiter("\\[|\\]|\\n");
        String key = "";
        String value = "";

        while(configScanner.hasNext()) {
            String s = configScanner.nextLine();
            if (s.equals("")) {
                continue;
            }
            else if (s.charAt(0) == '[') {
                key = s;
                value = configScanner.nextLine();
                configValues.put(key, value);
            }
        }

        for (Object o : configValues.entrySet()) {
            Map.Entry me = (Map.Entry) o;
            System.out.print(me.getKey() + ": ");
            System.out.println(me.getValue());
        }
        System.out.println(getServerIP());
        System.out.println(getFormattedServerIP()[0]);
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
        if (configValues.get(SERVER_IP) != null)
            return configValues.get(SERVER_IP);
        else
            return null;
    }

    public void setServerIP(String serverIP) {
        configValues.put(SERVER_IP, serverIP);
        printConfig();
    }

    public String getUserName() {
        if (configValues.get(REMEMBER_USER) != null)
            return configValues.get(REMEMBER_USER);
        else
            return null;
    }

    public void setUserName(String userName) {
        configValues.put(REMEMBER_USER, userName);
        printConfig();
    }


    public String[] getFormattedServerIP() {
        if (configValues.get(SERVER_IP) != null)
            return configValues.get(SERVER_IP).split(":");
        else
            return null;
    }

    public String getBackgroundPath() {
        if (configValues.get(BACKGROUND_PATH) != null) {
            return configValues.get(BACKGROUND_PATH);
        }
        else
            return null;
    }

    public void setBackgroundPath(String path) {
        configValues.put(BACKGROUND_PATH, path);
        printConfig();
    }
}

