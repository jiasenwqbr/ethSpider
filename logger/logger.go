package logger

import (
	logging "github.com/ipfs/go-log/v2"
)

// SetupLog 设置Logger级别
func SetupLog(name string) logging.StandardLogger {
	var logger = logging.Logger(name)
	_ = logging.SetLogLevel("*", "INFO")
	//_ = logging.SetLogLevel("*", "WARN")

	return logger
}
