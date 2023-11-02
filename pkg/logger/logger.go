package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"runtime"
)

const (
	dir         = "logs"
	logFileName = "all.log"
)

var l *logrus.Logger

type Logger struct {
	*logrus.Entry
}

type writerHook struct {
	levels  []logrus.Level
	writers []io.Writer
}

func (wh *writerHook) Levels() []logrus.Level {
	return wh.levels
}

func (wh *writerHook) Fire(e *logrus.Entry) error {
	msg, err := e.String()
	if err != nil {
		return err
	}

	for _, w := range wh.writers {
		w.Write([]byte(msg))
	}

	return nil
}

func init() {
	l = logrus.New()
	l.ReportCaller = true

	l.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return frame.Function, fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
	})

	err := os.MkdirAll(dir, 0744)
	if err != nil {
		l.Fatalln(err)
	}

	logFile, err := os.OpenFile(fmt.Sprintf("%s/%s", dir, logFileName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0744)
	if err != nil {
		l.Fatalln(err)
	}
	l.SetOutput(io.Discard)

	l.AddHook(&writerHook{
		levels:  logrus.AllLevels,
		writers: []io.Writer{os.Stdout, logFile},
	})
}

func New(level string) (*Logger, error) {
	parseLvl, err := logrus.ParseLevel(level)
	if err != nil {
		return nil, err
	}
	l.SetLevel(parseLvl)
	return &Logger{logrus.NewEntry(l)}, nil
}
