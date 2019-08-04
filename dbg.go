package dbg

import (
	"io"

	"os"

	"github.com/sirupsen/logrus"
)

var (
	Logger = New("default", os.Stdout)
)

func SetDebugLevel() {
	Logger.SetDebugLevel()
}

func (l *L) Debugln(args ...interface{}) {
	Logger.Debugln(args...)
}

func Debug(format string, args ...interface{}) {
	Logger.Debug(format, args...)
}

func New(name string, o io.Writer) *L {
	l := logrus.New()
	l.SetOutput(o)
	return &L{l}
}

type L struct {
	logger *logrus.Logger
	name string
}

func (l *L) SetDebugLevel() {
	l.logger.SetLevel(logrus.DebugLevel)
}

func (l *L) Debugln(args ...interface{}) {
	l.logger.WithField({"app":l.name}).Debugln(args...)
}

func (l *L) Debug(format string, args ...interface{}) {
	l.logger.WithField({"app":l.name}).Debugf(format, args...)
}
