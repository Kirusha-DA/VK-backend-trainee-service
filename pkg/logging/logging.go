package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writers   []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}

	for _, w := range hook.Writers {
		w.Write([]byte(line))
	}

	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var entry *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{entry}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

func init() {
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		Writers:   []io.Writer{allFile, os.Stdout},
		LogLevels: logrus.AllLevels,
	})

	l.SetLevel(logrus.TraceLevel)

	entry = logrus.NewEntry(l)
}
