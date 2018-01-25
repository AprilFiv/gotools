package aflog

import (
	"os"
	"time"
	"log"
	"strings"
)

type Aflog struct {
	appName string
	logName string
	logger  *log.Logger
	fd *os.File
}

func (currentlog *Aflog) Info(content string) {
	currentlog.flushLog()
	if !pathExists(currentlog.logName) {
		logFile, _ := os.Create(currentlog.logName)
		logger := log.New(logFile, "[Info] ",log.LstdFlags)
		logger.Println(content)
		defer logFile.Close()
	} else {
		logFile, _ := os.OpenFile(currentlog.logName, os.O_APPEND, 0)
		logger := log.New(logFile, "[Info] ", log.LstdFlags)
		logger.Println(content)
		defer logFile.Close()
	}
}
func (currentlog *Aflog) Error(content string) {
	currentlog.flushLog()
	if !pathExists(currentlog.logName) {
		logFile, _ := os.Create(currentlog.logName)
		logger := log.New(logFile, "[Error] ",log.LstdFlags)
		logger.Println(content)
		defer logFile.Close()
	} else {
		logFile, _ := os.OpenFile(currentlog.logName, os.O_APPEND, 0)
		logger := log.New(logFile, "[Error] ", log.LstdFlags)
		logger.Println(content)
		defer logFile.Close()
	}
}
func (currentlog *Aflog) Debug(content string) {
	currentlog.flushLog()
	if !pathExists(currentlog.logName) {
		logFile, _ := os.Create(currentlog.logName)
		logger := log.New(logFile, "[Debug] ",log.LstdFlags)
		logger.Println(content)
		defer logFile.Close()
	} else {
		logFile, _ := os.OpenFile(currentlog.logName, os.O_APPEND, 0)
		logger := log.New(logFile, "[Debug] ", log.LstdFlags)
		logger.Println(content)
		defer logFile.Close()
	}
}
func (currentlog *Aflog) Warn(content string) {
	currentlog.flushLog()
	if !pathExists(currentlog.logName) {
		logFile, _ := os.Create(currentlog.logName)
		logger := log.New(logFile, "[Warn] ",log.LstdFlags)
		logger.Println(content)
		defer logFile.Close()
	} else {
		logFile, _ := os.OpenFile(currentlog.logName, os.O_APPEND, 0)
		logger := log.New(logFile, "[Warn] ", log.LstdFlags)
		logger.Println(content)
		defer logFile.Close()
	}
}
func pathExists(path string) (bool) {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func GetLog(name string) Aflog {
	filePath := name + ".log." + time.Now().Format("2006-01-02")
	aflog := Aflog{
		appName: name,
		logName: filePath,
	}
	return aflog
}
func (currentlog *Aflog) flushLog(){
	filePath := currentlog.appName + ".log." + time.Now().Format("2006-01-02")
	currentlog.logName = filePath
}

