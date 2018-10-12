package logs

import (
	"github.com/sirupsen/logrus"
	"flag"
	"os"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	var logPath string
	flag.StringVar(&logPath, "log_path","./log_scheduler","log file path")
	flag.Parse()
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