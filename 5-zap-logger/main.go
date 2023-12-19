package main

import (
	"fmt"
	//"os"

	"go.uber.org/zap"
	//"go.uber.org/zap/zapcore"
)

func main() {

		//BASIC SETUP
	   	logger, _ := zap.NewProduction()
	   	//logger, _ := zap.NewDevelopment()

	   	logger.Debug("This is my first zap logger DEBUG message")
	   	fmt.Println()
	   	logger.Info("This is my first zap logger INFO message",zap.String("url","wwww.deepak.com"))
	   	fmt.Println()
	   	logger.Warn("This is my first zap logger WARNING message",zap.String("url","wwww.deepak.com"))
	   	fmt.Println()
	   	logger.Error("This is my first zap logger ERROR message",zap.String("url","wwww.deepak.com"))

/* 	//Configured Setup
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)

	logFile, _ := os.OpenFile("text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel // this will only print the logs which have log level greater than defaultLogLevel

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	logger.Debug("This is my first zap logger DEBUG message")
	fmt.Println()
	logger.Info("This is my first zap logger INFO message", zap.String("url", "wwww.deepak.com"))
	fmt.Println()
	logger.Warn("This is my first zap logger WARNING message", zap.String("url", "wwww.deepak.com"))
	fmt.Println()
	logger.Error("This is my first zap logger ERROR message", zap.String("url", "wwww.deepak.com")) */
}
