package log

import (
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

var logger = logrus.New()

var (
	loggerShareInstance IMLoggerInterface
	once sync.Once
)

func IMLoggerShareInstance() IMLoggerInterface{
	once.Do(func ()  {
		loggerShareInstance = &imLoggerImplement{}
	})
	return loggerShareInstance
}

type IMLoggerInterface interface {
	Info(str string)
	Debug(str string)
	Error(str string)
	Warn(str string)
}

type imLoggerImplement struct{}

func (log *imLoggerImplement) Info(str string) {
	_, file, line, _ := runtime.Caller(1)
	logger.Infof("[%s:%d] %s", file, line, str)
}

func (log *imLoggerImplement) Debug(str string) {
	_, file, line, _ := runtime.Caller(1)
	logger.Debugf("[%s:%d] %s", file, line, str)
}

func (log *imLoggerImplement) Warn(str string) {
	_, file, line, _ := runtime.Caller(1)
	logger.Warningf("[%s:%d] %s", file, line, str)
}

func (log *imLoggerImplement) Error(str string) {
	_, file, line, _ := runtime.Caller(1)
	logger.Errorf("[%s:%d] %s", file, line, str)
}


