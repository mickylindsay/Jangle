package com.jangle.ConroyTest;

import java.io.IOException;

import javax.swing.text.html.HTMLEditorKit.Parser;

import com.jangle.*;
import com.jangle.client.*;
import com.jangle.communicate.Client_ParseData;
import com.jangle.communicate.CommUtil;

public class Test {

	public static void main(String[] args) throws IOException, InterruptedException {

		Client Cl = new Client();
		Client_ParseData Parse = null;
		// TestServer server = new TestServer(9090);

		try {
			Parse = new Client_ParseData(Cl, "localhost", 9090);
			System.out.println("generated client");
		} catch (IOException e1) {
			// TODO Auto-generated catch block
			e1.printStackTrace();
		}

		Thread.sleep(1000);

		// EDIT BELOW HERE

		Message mess = TestUtil.genTest();

		mess.setMessageContent(
				"gr8 b8 m8. i rel8 str8 appreci8 nd congratul8. i r8 dis b8 an 8/8. plz no h8, i'm str8 ir8. cr8 more cant w8. we shood convers8 i wont ber8, my number is 8888888 ask for N8. no calls l8 or out of st8. if on a d8, ask K8 to loc8. even with a full pl8 i always hav time to communic8 so dont hesit8. dont forget to medit8 and particip8 and masturb8 to allevi8 ur ability to tabul8 the f8. we should meet up m8 and convers8 on how we can cre8 more gr8 b8, im sure everyone would appreci8 no h8. i dont mean to defl8 ur hopes, but itz hard to dict8 where the b8 will rel8 and we may end up with out being appreci8d, im sure u can rel8. we can cre8 b8 like alexander the gr8, stretch posts longer than the nile's str8s. well be the captains of b8 4chan our first m8s the growth r8 will spread to reddit and like reel est8 and be a flow r8 of gr8 b8 like a blind d8 well coll8 meet me upst8 where we can convers8 or ice sk8 or lose w8 infl8 our hot air baloons and fly tail g8. we cood land in kuw8, eat a soup pl8 followed by a dessert pl8 the payment r8 wont be too ir8 and hopefully our currency wont defl8. well head to the israeli-St8, taker over like herod the gr8 and b8 the jewish masses 8 million m8. we could interrel8 communism thought it's past it's maturity d8, a department of st8 volunteer st8. reduce the infant mortality r8, all in the name of making gr8 b8 m8﻿");

		// server.sendToClient(TestUtil.TEST_MESS_FROM_SERVER);

		Parse.sendMessage(mess);
		
		Thread.sleep(2000);

		System.out.println("\n\n" + Cl.getMessages().get(0).getMessageContent());
	}

}
