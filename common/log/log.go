package log

import (
	"fmt"
	"log"
	"sync"
)

var (
	fileWriter  *dailyFileWriter
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func Config(fileName string) {
	fileWriter = &dailyFileWriter{
		fileName:       fileName,
		lastYearDay:    -1,
		switchFileLock: &sync.Mutex{},
	}

	infoLogger = log.New(
		fileWriter,
		"[ INFO ]",
		log.Lshortfile|log.Lmicroseconds|log.Ldate,
	)

	errorLogger = log.New(
		fileWriter,
		"[ ERROR ]",
		log.Lshortfile|log.Lmicroseconds|log.Ldate,
	)
}

func Info(format string, valArray ...interface{}) {
	_ = infoLogger.Output(
		2,
		fmt.Sprintf(format, valArray...),
	)
}

func Error(formatter string, dataArray ...interface{}) {
	_ = errorLogger.Output(
		2,
		fmt.Sprintf(formatter, dataArray...),
	)
}
