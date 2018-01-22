aflog
=====
Usages
-------
<h3>First import the dependency</h3></br>
import "github.com/AprilFiv/gotools/aflog"</br>
<h3>And then you can use like this :</h3></br>
log :=aflog.GetLog("myapp") //your logfile name . it will generate a file like "xxx.log.2006-01-02" every day will new one</br>
log.Info("test")</br>
log.Debug("test")</br>
log.Warn("test")</br>
log.Error("test")</br>
