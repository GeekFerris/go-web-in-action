package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"runtime/debug"
	"time"
)

const LogFilePath string = "./runtime/log"

func init() {
	// 设置日志格式为json
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetReportCaller(false)
}

func Write(msg string, filename string) {
	setOutputFile(logrus.InfoLevel, filename)
	logrus.Info(msg)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.DebugLevel, "debug")
	logrus.WithFields(fields).Debug(args)
}

func Info(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.InfoLevel, "debug")
	logrus.WithFields(fields).Info(args)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.WarnLevel, "debug")
	logrus.WithFields(fields).Warn(args)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.FatalLevel, "debug")
	logrus.WithFields(fields).Fatal(args)
}

func Error(fields logrus.Fields, args ...interface{}) {
	setOutputFile(logrus.ErrorLevel, "debug")
	logrus.WithFields(fields).Error(args)
}

func setOutputFile(level logrus.Level, logName string) {
	if _, err := os.Stat(LogFilePath); os.IsNotExist(err) {
		err = os.MkdirAll(LogFilePath, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", LogFilePath, err))
		}
	}

	timeStr := time.Now().Format("2006-01-02")
	filename := path.Join(LogFilePath, logName+"_"+timeStr+".log")

	var err error
	os.Stderr, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open log file error", err)
	}
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(level)
	return
}

func LoggerToFile() gin.LoggerConfig {
	if _, err := os.Stat(LogFilePath); os.IsNotExist(err) {
		err = os.MkdirAll(LogFilePath, 0777)
		if err != nil {
			panic(fmt.Errorf("create log dir '%s' error: %s", LogFilePath, err))
		}
	}
	timeStr := time.Now().Format("2006-01-02")
	filename := path.Join(LogFilePath, "success_"+timeStr+".log")

	os.Stderr, _ = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	var conf = gin.LoggerConfig{
		Formatter: func(params gin.LogFormatterParams) string {
			return fmt.Sprintf("%s - %s \" %s %s %s %d %s \"%s\" %s\"\n",
				params.TimeStamp.Format("2006-01-02 15:04:05"),
				params.ClientIP,
				params.Method,
				params.Path,
				params.Request.Proto,
				params.StatusCode,
				params.Latency,
				params.Request.UserAgent(),
				params.ErrorMessage,
			)
		}, Output: io.MultiWriter(os.Stdout, os.Stderr),
	}

	return conf
}

func Recover(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			if _, errDir := os.Stat(LogFilePath); os.IsNotExist(errDir) {
				errDir = os.MkdirAll(LogFilePath, 0777)
				if err != nil {
					panic(fmt.Errorf("create log dir '%s' error: %s", LogFilePath, err))
				}
			}
			timeStr := time.Now().Format("2006-01-02")
			filename := path.Join(LogFilePath, "error_"+timeStr+".log")

			f, errFile := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			if errFile != nil {
				fmt.Println(errFile)
			}

			timeFileStr := time.Now().Format("2006-01-02 15:04:05")
			f.WriteString("panic error time: " + timeFileStr + "\n")
			f.WriteString(fmt.Sprintf("%v", err) + "\n")
			f.WriteString("stacktrace from panic: " + string(debug.Stack()) + "\n")
			f.Close()
			c.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  fmt.Sprintf("%v", err),
			})
			// 终止后续接口调用，不加的话recover到异常后，还会执行接口里后续的代码
			c.Abort()
		}
	}()
	c.Next()
}
