package logrus

import (
	"fmt"
	debug "log"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
)

var testLog *debug.Logger

func init() {
	debug.SetFlags(debug.LstdFlags | debug.Lshortfile)
	debug.SetOutput(new(debugWriter))
	testLog = debug.New(new(debugWriter), "test: ", debug.LstdFlags|debug.Lshortfile) // os.Stderr
}

func TestLog(t *testing.T) {
	log := logrus.New()
	log.Hooks.Add(new(LocationHook))
	log.Info("Test location")
	debug.Println("Test location")
	testLog.Println("Test location")
}

type debugWriter struct{}

func (writer debugWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format("2006-01-02T15:04:05.999Z") + " [DEBUG] " + string(bytes))
}
