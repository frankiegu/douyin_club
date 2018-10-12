package logs

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	var logPath string = "./log_web"
	logFd, err := os.OpenFile( logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644 )
	//defer logFd.Close()
	if err == nil {
		Log.Out = logFd
	} else {
		Log.Infof("open logFilePath err:%s", err.Error())
	}
	Log.Out = os.Stdout
	Log.SetLevel( logrus.DebugLevel )
}
