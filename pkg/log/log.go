package log

import (
	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
	"github.com/sirupsen/logrus"
)

// Level for logger
type Level uint32

// String form of a Level
func (l Level) String() string {
	return logrus.Level(l).String()
}

// SetLevel is used to change the level of logging globally
var SetLevel = func(level Level) {
	logrus.SetLevel(logrus.Level(level))
}

// GetLevel is used to get the current level of logging
var GetLevel = func() Level {
	return Level(logrus.GetLevel())
}

// DebugLevel is a threshold level
const DebugLevel Level = Level(logrus.DebugLevel)

// InfoLevel is a threshold level
const InfoLevel Level = Level(logrus.InfoLevel)

// WarnLevel is a threshold level
const WarnLevel Level = Level(logrus.WarnLevel)

// ErrorLevel is a threshold level
const ErrorLevel Level = Level(logrus.ErrorLevel)

// FatalLevel is a threshold level
const FatalLevel Level = Level(logrus.FatalLevel)

// Interface for the logger
type Interface interface {
	Debug(args ...interface{})
	Print(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}

// Create a new logger
func Create(frameworkName string, usingStackdriver bool) Interface {
	logrus.SetLevel(logrus.InfoLevel)
	if usingStackdriver {
		logrus.SetFormatter(stackdriver.NewFormatter(
			stackdriver.WithService(frameworkName),
		))
	}
	return logrus.WithFields(logrus.Fields{
		"framework_name": frameworkName,
	})
}
