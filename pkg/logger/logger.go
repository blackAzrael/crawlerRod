package logger

import (
	"bytes"
	"crawlerRod/pkg/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
)


const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type Fields map[string]interface{}

var logLevelMap = map[string]logrus.Level{
	//"Trace":           logrus.TraceLevel,
	"Debug": logrus.DebugLevel,
	"Info":  logrus.InfoLevel,
	"Warn":  logrus.WarnLevel,
	"Error": logrus.ErrorLevel,
	"Fatal": logrus.FatalLevel,
	//"Panic":           logrus.PanicLevel,
}
type MyFormatter struct {
	DisableColors bool
}
func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error){
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	case logrus.InfoLevel:
		levelColor = blue
	default:
		levelColor = blue
	}

	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [\u001B[%dm%s\u001B[0m] [%s:%d %s] %s\n",
			timestamp, levelColor,entry.Level, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else{
		newLog = fmt.Sprintf("[%s] \u001B[%dm%s\u001B[0m %s\n", timestamp, levelColor,entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}



var Logger *logrus.Logger

func init() {
	Logger = logrus.New()
	level := config.Config.LogLevel
	Logger.SetLevel(logLevelMap[level])
	//Logger.SetFormatter(&logrus.TextFormatter{
	//	ForceQuote:true,    //键值对加引号
	//	TimestampFormat:"2006-01-02 15:04:05",  //时间格式
	//	FullTimestamp:true,
	//})

	Logger.SetReportCaller(true)
	Logger.SetFormatter(&MyFormatter{})


}