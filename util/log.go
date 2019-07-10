package util

import (
	"github.com/sirupsen/logrus"
	"os"
)

/**
* @Author: hexing
* @Date: 19-7-10 下午3:57
 */
func init() {
	// 设置日志格式为json格式
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	logrus.SetLevel(logrus.WarnLevel)
}
