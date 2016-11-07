package main

import(
	"log"
	"os"
)

func Init_Logger(){
	if(!jangle.no_logging){
		f, err := os.OpenFile("testlogfile", os.O_CREATE | os.O_APPEND, 0666)
		jangle.log_file = f;
		if err != nil {
				log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
		log.SetOutput(f)
	}
	
}

func Log(s string){
	if(!jangle.no_logging){
		log.Println(s);
	}
}

func Logf(a ...interface{}){
	if(!jangle.no_logging){
		log.Println(a...);
	}
}

func Fatal(s string){
	if(!jangle.no_logging){
		log.Fatalln(s);
	}
}

func Fatalf(a ...interface{}){
	if(!jangle.no_logging){
		log.Fatalln(a...);
	}
}