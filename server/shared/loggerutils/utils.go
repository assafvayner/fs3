package loggerutils

import (
	"fmt"
	"log"
	"os"
)

func InitLogger(service string) *log.Logger {
	path := "/log/" + service + ".log"
	logFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to open log file")
		os.Exit(1)
	}
	return log.New(logFile, service+": ", log.LstdFlags|log.Llongfile|log.Lmsgprefix)
}
