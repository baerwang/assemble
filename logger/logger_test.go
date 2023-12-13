package logger

import "testing"

func TestLogger(t *testing.T) {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
	Debugf("%s", "debug")
	Infof("%s", "info")
	Warnf("%s", "warn")
	Errorf("%s", "error")
}
