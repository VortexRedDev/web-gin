package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
)

var (
	//日志地址
	logFilePath = "./log"
	//日志文件名
	logFileName = "gin.log"
)
var Logger = logrus.New()

func init() {
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("err", err)
	}

	Logger.SetFormatter(&logrus.JSONFormatter{

		TimestampFormat: "2006-01-02 15:04:05", //时间格式

	})
	Logger.SetOutput(io.MultiWriter(src, os.Stdout))
	Logger.SetLevel(logrus.DebugLevel)
}

func LoggerWriter() io.Writer {
	return Logger.Out
}
