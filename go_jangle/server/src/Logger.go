package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Init_Logger() {
	if jangle.logging {
		var err error
		t := time.Now()
		dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		dir = dir + "/" + t.Format(time.ANSIC)
		dir = strings.Replace(dir, " ", "_", -1)
		dir = strings.Replace(dir, ":", "_", -1)
		dir = strings.Replace(dir, "__", "_", -1)
		jangle.log_file, err = os.OpenFile(dir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}

		log.SetOutput(jangle.log_file)
	}

}

func Log(s string) {
	if jangle.logging {
		log.Println(s)
	}
}

func Logln(a ...interface{}) {
	if jangle.logging {
		log.Println(a...)
	}
}

func Warn(s string) {
	if jangle.logging_warn {
		log.Panicln(s)
	}
}

func Warnln(a ...interface{}) {
	if jangle.logging_warn {
		log.Panicln(a...)
	}
}

func Fatal(s string) {
	if jangle.logging {
		log.Fatalln(s)
	}
}

func Fatalln(a ...interface{}) {
	if jangle.logging {
		log.Fatalln(a...)
	}
}