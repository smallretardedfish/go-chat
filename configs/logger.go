package configs

import "github.com/sirupsen/logrus"

type Logger interface {
	Info(args ...interface{})
	Warn(args ...interface{})
	Fatal(args ...interface{})
}

//TODO check params of logger
func NewLogger() *logrus.Logger {
	return logrus.New()
}
