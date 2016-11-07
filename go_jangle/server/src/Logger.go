package main

import(
	"log"
	"os"
	"time"
	"path/filepath"
)

func Init_Logger(){
	if(jangle.logging){
		t := time.Now();
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]));
		dir = dir + "/" + t.Format(time.UnixDate);
		f, err := os.OpenFile(dir, os.O_CREATE | os.O_APPEND, 0666)
		jangle.log_file = f;
		if err != nil {
				log.Fatalf("error opening file: %v", err)
		}
		log.SetOutput(f)
		defer f.Close()
	}
	
}

func Log(s string){
	if(jangle.logging){
		log.Println(time.Now().Format(time.UnixDate), s);
	}
}

func Logln(a ...interface{}){
	if(jangle.logging){
		log.Print(time.Now().Format(time.UnixDate));
		log.Println(a...);
	}
}

func Warn(s string){
	if(jangle.logging_warn){
		log.Println(time.Now().Format(time.UnixDate), s);
	}
}

func Warnln(a ...interface{}){
	if(jangle.logging_warn){
		log.Print(time.Now().Format(time.UnixDate));
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