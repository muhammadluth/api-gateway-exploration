package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

//Event is public function to create logging
func Event(text ...string) {
	logText := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	logrus.SetLevel(logrus.InfoLevel)
	logrus.Info("[EVE]" + timenow + logText)
}

//Message is public function to create logging
func Message(text ...string) {
	logText := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("[MSG]" + timenow + logText)
}

//Warning is public function to create logging
func Warning(text ...string) {
	logText := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	logrus.SetLevel(logrus.WarnLevel)
	logrus.Info("[WARN]" + timenow + logText)
}

//Error is public function to create logging
func Error(err error, text ...string) {
	logText := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	err1 := errors.Wrap(err, err.Error())
	stackTrace := "[" + strings.Replace(strings.Replace(fmt.Sprintf("%+v", err1), "\n\t", " ", -1),
		"\n", " | ", -1) + "]"
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.Error("[ERR]" + timenow + logText + stackTrace)

}

//Fatal is public function to create logging
func Fatal(err error, text ...string) {
	sLog := "[" + strings.Join(text, "][") + "]"
	timenow := "[" + time.Now().Format("2006-01-02 15:04:05.000") + "]"
	err1 := errors.Wrap(err, err.Error())
	stackTrace := "[" + strings.Replace(strings.Replace(fmt.Sprintf("%+v", err1), "\n\t", " ", -1),
		"\n", " | ", -1) + "]"
	logrus.SetLevel(logrus.FatalLevel)
	logrus.Fatal("[FTL]" + timenow + sLog + stackTrace)

}

func SetupLogging() {
	logrus.SetOutput(ioutil.Discard) // Send all logs to nowhere by default
	logrus.SetLevel(logrus.FatalLevel)
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetLevel(logrus.ErrorLevel)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
		Writer: os.Stderr,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
		},
	})
	logrus.AddHook(&writer.Hook{ // Send info and debug logs to stdout
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
			logrus.TraceLevel,
		},
	})
}
