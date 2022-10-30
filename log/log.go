package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github/xiaoda-ye/web-gin/config"
	"io"
	"os"
	"path"
)

var (
	//日志地址
	logFilePath = "./"
	//日志文件名
	logFileName = "gin.log"
)
var Log *logrus.Entry

func init() {
	Log = loadLog()
}
func loadLog() *logrus.Entry {

	lg := logrus.New()
	lg.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05", //时间格式
	})
	env := config.Conf().Env
	if env == "dev" {
		lg.SetOutput(os.Stdout) //输出控制台
	} else if env == "pro" {
		// 日志文件
		fileName := path.Join(logFilePath, logFileName)
		// 写入文件
		src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("err", err)
		}
		lg.SetOutput(src) // 输出日志文件
		//lg.SetOutput(io.MultiWriter(src,os.Stdout)) 同时输出控制台和日志文件
	}
	lg.SetLevel(logrus.DebugLevel)
	Log := lg.WithFields(logrus.Fields{
		"version": "1.0",
		"env":     config.Conf().Env, // 日志的公共属性
	})
	return Log
}
func LoggerWriter() io.Writer {
	return Log.Logger.Out
}
