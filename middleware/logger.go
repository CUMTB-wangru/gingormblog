package middleware

// 项目日志处理---日志分割，这里不用gin框架自带的日志文件是为了方便自定义日志文件需要写入的内容
import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	// 日志文件的保存路径
	filePath := "log/log"
	//linkName := "latest_log.log"
	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		fmt.Println("err:", err)
	}
	// 实例化logger
	logger := logrus.New()

	// 保存到src中
	logger.Out = scr

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 分割日志
	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		// 最长保存时间
		retalog.WithMaxAge(7*24*time.Hour),
		// 24小时分割一次
		retalog.WithRotationTime(24*time.Hour),
		//retalog.WithLinkName(linkName),
	)

	// 写入日志文件的内容
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.AddHook(Hook)

	return func(c *gin.Context) {
		startTime := time.Now()
		//go中间件：c.Next() 调用下一个中间件或者接口处理函数
		c.Next()
		// 在这里可以处理请求返回给用户之前的逻辑
		stopTime := time.Since(startTime).Milliseconds()
		spendTime := fmt.Sprintf("%d ms", stopTime)
		// 客户端请求过来的名字
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		// 客户端ip
		clientIp := c.ClientIP()
		// 请求过来的设备名
		userAgent := c.Request.UserAgent()
		// 请求过来的文件长度
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		// 请求方法
		method := c.Request.Method
		// 请求路劲
		path := c.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
