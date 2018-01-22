aflog
=====
Usages
-------
#####First import the dependency
import "github.com/AprilFiv/gotools/aflog"
#####And then you can use like this :
log :=aflog.GetLog("myapp") //your logfile name . it will generate a file like "xxx.log.2006-01-02" every day will new one</br>
log.Info("test")</br>
log.Debug("test")</br>
log.Warn("test")</br>
log.Error("test")</br>
