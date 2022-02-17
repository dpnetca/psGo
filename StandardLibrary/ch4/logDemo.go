package main

import (
	"log"
	"os"
)

type messageType int

const (
	INFO messageType = 0 + iota
	WARNING
	ERROR
	FATAL
)

var (
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "INFO: ", log.LUTC|log.Lmicroseconds|log.Llongfile)
	WarningLogger = log.New(file, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	FatalLogger = log.New(file, "Fatal: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	log.Println("This is a log message!")

	// // open log.txt create it if it doesn't exist, append to it if it does
	// // open it as write only, and set permissions to 0666
	// file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.SetOutput(file)
	// log.Println("This is a log message!")
	writeLog(INFO, "this is an information message")
	writeLog(WARNING, "this is an warning message")
	writeLog(ERROR, "this is an error message")
	// writeLog(FATAL, "this is an fatal message")

	InfoLogger.Println("v2 this is an information message")
	WarningLogger.Println("v2 this is an warning message")
	ErrorLogger.Println("v2 this is an error message")
	FatalLogger.Fatal("v2 this is an fatal message")

}

func writeLog(messagetype messageType, message string) {
	// open log.txt create it if it doesn't exist, append to it if it does
	// open it as write only, and set permissions to 0666
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	switch messagetype {
	case INFO:
		logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case WARNING:
		logger := log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case ERROR:
		logger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println(message)
	case FATAL:
		logger := log.New(file, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Fatal(message) // kills application after writing message
	}
}
