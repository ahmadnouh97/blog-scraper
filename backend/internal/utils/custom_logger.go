package utils

import (
	"log"
	"os"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Green  = "\033[32m"
)

type CustomLogger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
}

func NewCustomLogger() *CustomLogger {
	return &CustomLogger{
		debugLogger:   log.New(os.Stdout, Blue+"DEBUG: "+Reset, log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:    log.New(os.Stdout, Green+"INFO: "+Reset, log.Ldate|log.Ltime),
		warningLogger: log.New(os.Stdout, Yellow+"WARNING: "+Reset, log.Ldate|log.Ltime),
		errorLogger:   log.New(os.Stderr, Red+"ERROR: "+Reset, log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (c *CustomLogger) Debug(format string, args ...interface{}) {
	c.debugLogger.Printf(format, args...)
}

func (c *CustomLogger) Info(format string, args ...interface{}) {
	c.infoLogger.Printf(format, args...)
}

func (c *CustomLogger) Warning(format string, args ...interface{}) {
	c.warningLogger.Printf(format, args...)
}

func (c *CustomLogger) Error(format string, args ...interface{}) {
	c.errorLogger.Printf(format, args...)
}
