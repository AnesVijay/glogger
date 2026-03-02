package glogger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LogLVL int

const (
	Debug LogLVL = iota
	Info
	Warning
	Error
)

type Logger struct {
	file   string
	logLVL LogLVL
}

var logger Logger

func GetLogger() *Logger {
	return &logger
}

func getNowTimestamp() string {
	return string(time.Now().Format("02 Jan 2006 15:04:05"))
}

/*
logLevel:

	0 - Debug,
	1 - Info,
	2 - Warning,
	3 - Error
*/
func InitLogger(path string, logLevel int) {
	file := strings.Join([]string{
		string(time.Now().Format("02-01-2006")),
		"txt"},
		".")
	path = strings.Join(
		[]string{path, file},
		"/")

	if _, err := os.Stat(path); err != nil {
		pathParts := strings.Split(path, "/")
		pathToDir := strings.Join(pathParts[:len(pathParts)-1], "/")
		err := os.MkdirAll(pathToDir, 0764)
		if err != nil {
			logger.SendError(fmt.Sprintf("failed to create dir for log file: %v\n", err))
			log.Fatal(err)
		}
		_, err = os.Create(path)
		if err != nil {
			logger.SendError(fmt.Sprintf("failed to create log file: %v\n", err))
			log.Fatal(err)
		}
	}
	logger = Logger{
		file:   path,
		logLVL: LogLVL(logLevel),
	}
}

func (l *Logger) writeToLogFile(log string) {
	file, err := os.OpenFile(l.file, os.O_RDWR|os.O_APPEND, 0o644)
	if err != nil {
		fmt.Printf("ERROR: failed to open log file to write: %v\n", err)
	}
	file.WriteString(log + "\n")
}

//------------------------ Функции уровней логирования ------------------------

func (l *Logger) SendDebug(msg string) {
	if l.logLVL == LogLVL(0) {
		log := fmt.Sprintf("DEBUG: %s", msg)
		fmt.Println(log)
		l.writeToLogFile("[" + getNowTimestamp() + "]" + " " + log)
	}
}

func (l *Logger) SendInfo(msg string) {
	if l.logLVL <= LogLVL(1) {
		log := fmt.Sprintf("INFO: %s", msg)
		fmt.Println(log)
		l.writeToLogFile("[" + getNowTimestamp() + "]" + " " + log)
	}
}

func (l *Logger) SendWarning(msg string) {
	if l.logLVL <= LogLVL(2) {
		log := fmt.Sprintf("WARNING: %s", msg)
		fmt.Println(log)
		l.writeToLogFile("[" + getNowTimestamp() + "]" + " " + log)
	}
}

func (l *Logger) SendError(msg string) {
	if l.logLVL <= LogLVL(3) {
		log := fmt.Sprintf("ERROR: %s", msg)
		fmt.Println(log)
		l.writeToLogFile("[" + getNowTimestamp() + "]" + " " + log)
	}
}
