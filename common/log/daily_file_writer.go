package log

import (
	"github.com/pkg/errors"
	"os"
	"path"
	"sync"
	"time"
)

type dailyFileWriter struct {
	fileName       string
	lastYearDay    int
	fileWriter     *os.File
	switchFileLock *sync.Mutex
}

func (w *dailyFileWriter) Write(p []byte) (int, error) {
	if p == nil || len(p) <= 0 {
		return 0, nil
	}

	_, _ = os.Stderr.Write(p)

	fileWriter, err := w.getFileWriter()
	if err != nil {
		_, _ = os.Stderr.Write([]byte(err.Error()))
		return 0, err
	}

	_, _ = fileWriter.Write(p)

	return len(p), nil
}

func (w *dailyFileWriter) getFileWriter() (*os.File, error) {
	yearDay := time.Now().YearDay()

	if yearDay == w.lastYearDay && w.fileWriter != nil {
		return w.fileWriter, nil
	}

	w.switchFileLock.Lock()
	defer w.switchFileLock.Unlock()

	//首先检查，然后再做
	if yearDay == w.lastYearDay && w.fileWriter != nil {
		return w.fileWriter, nil
	}

	//创建目录
	err := os.MkdirAll(path.Dir(w.fileName), os.ModePerm|os.ModeSticky)
	if err != nil {
		return nil, errors.Errorf("创建目录失败 fileName:%v, err:%v\n", w.fileName, err.Error())
	}
	//拼接文件名
	newDailyFile := w.fileName + "_" + time.Now().Format("20060102")

	//打开文件
	outputFile, err := os.OpenFile(
		newDailyFile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil || outputFile == nil {
		return nil, err
	}
	//关掉旧的文件
	if w.fileWriter != nil {
		w.fileWriter.Close()
	}

	w.lastYearDay = yearDay
	w.fileWriter = outputFile

	return outputFile, nil
}
