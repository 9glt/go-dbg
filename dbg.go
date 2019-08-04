package dbg

import (
	"io"

	"os"

	"github.com/sirupsen/logrus"
)

var (
	logger = New("default", os.Stdout)
)

func SetDebugLevel() {
	logger.SetDebugLevel()
}

func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func New(name string, o io.Writer) *L {
	l := logrus.New()
	l.SetOutput(o)
	l.SetFormatter(&logrus.TextFormatter{})
	l.SetLevel(logrus.InfoLevel)
	return &L{l, name}
}

type L struct {
	logger *logrus.Logger
	name   string
}

func (l *L) SetDebugLevel() {
	l.logger.SetLevel(logrus.DebugLevel)
}

func (l *L) Debugln(args ...interface{}) {
	l.logger.WithFields(logrus.Fields{"app": l.name}).Debugln(args...)
}

func (l *L) Debugf(format string, args ...interface{}) {
	l.logger.WithFields(logrus.Fields{"app": l.name}).Debugf(format, args...)
}
