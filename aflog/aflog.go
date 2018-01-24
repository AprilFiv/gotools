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
var currentLog Aflog
func (log *Aflog) Info(content string) {
	currentLog.checkTimeOut()
	log.logger.SetPrefix("[Info] ")
	log.logger.Println(content)
}
func (log *Aflog) Error(content string) {
	currentLog.checkTimeOut()
	log.logger.SetPrefix("[Error] ")
	log.logger.Println(content)
}
func (log *Aflog) Debug(content string) {
	currentLog.checkTimeOut()
	log.logger.SetPrefix("[Debug] ")
	log.logger.Println(content)
}
func (log *Aflog) Warn(content string) {
	currentLog.checkTimeOut()
	log.logger.SetPrefix("[Warn] ")
	log.logger.Println(content)
}
func (log *Aflog) checkTimeOut(){
	 var1 := strings.Split(log.logName,".")
	 suffix :=var1[len(var1)-1]
	 if suffix!=time.Now().Format("2006-01-02"){
	 	log.fd.Close();
	 	currentLog = GetLog(currentLog.appName)
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
		logName: filePath,
		logger:logger,
		fd :logFile,
	}
	currentLog = aflog
	return aflog
}

