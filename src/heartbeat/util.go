package heartbeat

import (
	"fmt"
	"os"
	"runtime"

	log "github.com/cihub/seelog"
)

// 以下函数均为日志函数
func LogFlush() {
	log.Flush()
}

// 使用日志配置文件@file，如果配置文件解析失败，就退出程序
func InitLogAsFile(file string) error {
	log.Infof("Log init as file %s ...", file)
	logger, err := log.LoggerFromConfigAsFile(file)
	ChkErrOnExit(err, fmt.Sprintf("Parse log config file %s Fail", file))

	log.ReplaceLogger(logger)
	ChkErrOnExit(err, "ReplaceLogger")
	return nil
}

// 判断致命错误，一旦发生，不可修复，程序退出
func ChkErrOnExit(err error, errMsg string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Errorf("[%s,%d]%s, unrecoverable error: %s", file, line, errMsg, err.Error())
		log.Flush()
		os.Exit(1)
	}
}