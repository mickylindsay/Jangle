package main

import(
	"log"
	"os"
	"time"
	"fmt"
)

func Init_Logger(){
	if(jangle.logging){
	var s string;
		t := time.Now();
		fmt.Sprint(&s, t);
		f, err := os.OpenFile("s", os.O_CREATE | os.O_APPEND, 0666)
		jangle.log_file = f;
		if err != nil {
				log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}
	
}

func Log(s string){
	if(jangle.logging){
		log.Println(s);
	}
}

func Logf(a ...interface{}){
	if(jangle.logging){
		log.Println(a...);
	}
}

func Fatal(s string){
	if(jangle.logging){
		log.Fatalln(s);
	}
}

func Fatalf(a ...interface{}){
	if(jangle.logging){
		log.Fatalln(a...);
	}
}