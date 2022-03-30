package mylog

import (
	"log"
	"os"
)

func init() {
	log.SetPrefix("[ld-logger]")
	log.SetFlags(log.Llongfile | log.Lshortfile | log.Ltime | log.Ldate)
	logFile, err := os.OpenFile("/mylog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	log.SetOutput(logFile)
}
