package logger

import (
	"flag"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	logpath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	logpath = logpath + "/pkg/logger/info.log"

	flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Printf("Log path is %s", logpath)
}
