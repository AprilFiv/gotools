package aflog

import (
	"os"
	"time"
	"log"
)

type Aflog struct {
	appName string
	logger  *log.Logger
}

func (log *Aflog) Info(content string) {
	log.logger.SetPrefix("[Info] ")
	log.logger.Println(content)
}
func (log *Aflog) Error(content string) {
	log.logger.SetPrefix("[Error] ")
	log.logger.Println(content)
}
func (log *Aflog) Debug(content string) {
	log.logger.SetPrefix("[Debug] ")
	log.logger.Println(content)
}
func (log *Aflog) Warn(content string) {
	log.logger.SetPrefix("[Warn] ")
	log.logger.Println(content)
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
	var logger *log.Logger
	var logFile *os.File
	filePath := name + ".log." + time.Now().Format("2006-01-02")
	if !pathExists(filePath) {
		logFile, _ = os.Create(filePath)
		logger = log.New(logFile, "[Info] ",log.LstdFlags)
	}
	logFile,_ = os.OpenFile(filePath,os.O_APPEND,0)
	logger = log.New(logFile, "[Info] ",log.LstdFlags)
	aflog := Aflog{
		appName: name,
		logger:logger,
	}
	return aflog
}

