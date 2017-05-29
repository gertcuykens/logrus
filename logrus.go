package logrus

import (
	"path"
	"runtime"
	"strings"

	"github.com/Sirupsen/logrus"
)

type LocationHook struct{}

func (hook LocationHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook LocationHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)
	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		name := fu.Name()
		if !strings.Contains(name, "github.com/Sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data["file"] = path.Base(file)
			entry.Data["func"] = path.Base(name)
			entry.Data["line"] = line
			break
		}
	}
	return nil
}
