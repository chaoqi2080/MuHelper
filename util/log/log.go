package log

import (
	"fmt"
	"log"
	"sync"
)

var (
	fileWriter    *dailyFileWriter
	infoLogger    *log.Logger
	errorLogger   *log.Logger
	warnLogger    *log.Logger
	extractLogger *log.Logger
)

func Config(fileName string) {
	fileWriter = &dailyFileWriter{
		fileName:       fileName,
		lastYearDay:    -1,
		switchFileLock: &sync.Mutex{},
	}

	infoLogger = log.New(
		fileWriter,
		"[ INFO ] ",
		log.Lshortfile|log.Lmicroseconds|log.Ltime,
	)

	warnLogger = log.New(
		fileWriter,
		"[ WARN ] ",
		log.Lshortfile|log.Lmicroseconds|log.Ltime,
	)

	errorLogger = log.New(
		fileWriter,
		"[ ERROR ] ",
		log.Lshortfile|log.Lmicroseconds|log.Ltime,
	)

	//为了方便提取日志文件添加
	extractLogger = log.New(fileWriter, "", 0)
}

func Info(format string, valArray ...interface{}) {
	_ = infoLogger.Output(
		2,
		fmt.Sprintf(format, valArray...),
	)
}

func Warn(formatter string, dataArray ...interface{}) {
	_ = warnLogger.Output(
		2,
		fmt.Sprintf(formatter, dataArray...),
	)
}

func Error(formatter string, dataArray ...interface{}) {
	_ = errorLogger.Output(
		2,
		fmt.Sprintf(formatter, dataArray...),
	)
}

// ExtractLog 专门用于提取日志文件的API，正常情况不要使用
func ExtractLog(line string) {
	_ = extractLogger.Output(2, line)
}
