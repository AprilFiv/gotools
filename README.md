aflog
=====
Usages
-------
First import the dependency
    import "github.com/AprilFiv/gotools/aflog"
    And then you can use like this :
    log :=aflog.GetLog("myapp") //your logfile name . it will generate a file like "xxx.log.2006-01-02" every day will new one
    log.Info("test")
    log.Debug("test")
    log.Warn("test")
    log.Error("test")
