package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("app-logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("could not open app-logs.txt")
	}
	defer file.Close()
	log.SetOutput(file)

	standardError := "connection timed out"
	log.Printf("operation failed due to: %s.This will be logged", standardError)
}
