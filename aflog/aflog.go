package aflog

import (
	"os"
	"time"
	"log"
)

type Aflog struct {
	appName string
	logName string
	logger  *log.Logger
	fd *os.File
}

func (currentlog *Aflog) Info(content string) {
	currentlog.flushLog()
	currentlog.logger.SetPrefix("[Info] ")
	currentlog.logger.Println(content)
	currentlog.fd.Close()
}
func (currentlog *Aflog) Error(content string) {
	currentlog.flushLog()
	currentlog.logger.SetPrefix("[Error] ")
	currentlog.logger.Println(content)
	currentlog.fd.Close()
}
func (currentlog *Aflog) Debug(content string) {
	currentlog.flushLog()
	currentlog.logger.SetPrefix("[Debug] ")
	currentlog.logger.Println(content)
	currentlog.fd.Close()
}
func (currentlog *Aflog) Warn(content string) {
	currentlog.flushLog()
	currentlog.logger.SetPrefix("[Warn] ")
	currentlog.logger.Println(content)
	currentlog.fd.Close()
}
//func (log *Aflog) checkTimeOut(){
//	 var1 := strings.Split(log.logName,".")
//	 suffix :=var1[len(var1)-1]
//	 if suffix!=time.Now().Format("2006-01-02"){
//	 	log.fd.Close();
//	 	currentLog = GetLog(currentLog.appName)
//	 }
//}
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
		logger.Println("test")
		defer logFile.Close()
	} else {
		logFile, _ = os.OpenFile(filePath, os.O_APPEND, 0)
		logger = log.New(logFile, "[Info] ", log.LstdFlags)
	}
	aflog := Aflog{
		appName: name,
		logName: filePath,
		logger:logger,
		fd :logFile,
	}
	return aflog
}
func (currentlog *Aflog) flushLog(){
	var logger *log.Logger
	var logFile *os.File
	filePath := currentlog.appName + ".log." + time.Now().Format("2006-01-02")
	if !pathExists(filePath) {
		logFile, _ = os.Create(filePath)
		logger = log.New(logFile, "[Info] ",log.LstdFlags)
	}
	logFile,_ = os.OpenFile(filePath,os.O_APPEND,0)
	logger = log.New(logFile, "[Info] ",log.LstdFlags)
	currentlog.logger = logger
	currentlog.fd = logFile
	currentlog.logName = filePath

}

