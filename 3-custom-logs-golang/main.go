package main

import (
	"log"
	"os"
)

var (
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func init() {
	applicationLogFile, err := os.OpenFile("application-log-file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("could not open application-log-file.txt")
	}

	InfoLog = log.New(applicationLogFile, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(applicationLogFile, "WARNING : ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(applicationLogFile, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)

	/* infoLogFile, err := os.OpenFile("info-log-file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("could not open info-log-file.txt")
	}
	InfoLog = log.New(infoLogFile, "INFO : ",log.Ldate|log.Ltime|log.Lshortfile)

	warningLogFile, err := os.OpenFile("warning-log-file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("could not open warning-log-file.txt")
	}
	WarningLog = log.New(warningLogFile, "WARNING : ", log.Ldate|log.Ltime|log.Lshortfile)

	errorLogFile, err := os.OpenFile("error-log-file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("could not open error-log-file.txt")
	}
	ErrorLog = log.New(errorLogFile, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile) */
}

func main() {
	standardLog := "email sent successfully"
	InfoLog.Printf("email status: %s", standardLog)

	warningLog := "request with 0 size input array recieved. handled successfully"
	WarningLog.Printf("request details: %s", warningLog)

	errorLog := "connection timed out"
	ErrorLog.Printf("operation failed due to : %s", errorLog)
}
