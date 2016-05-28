package glog

import (
	"testing"
)

func TestGLogger(t *testing.T) {
	Logger.Debug("Debug value")
	Logger.Info("Info value")
	Logger.Warn("Warn value")
	Logger.Error("Error value")
	Logger.Fatal("Fatal value")
	Logger.Panic("Panic value")
}
