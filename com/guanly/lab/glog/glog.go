package glog

import (
	"log"
	"os"
)

const (
	Debug = iota
	Info
	Warn
	Error
	Fatal
	Panic
)

type GLogger struct {
	logger   *log.Logger
	logLevel int
}

var Logger *GLogger

func init() {
	Logger = &GLogger{logger: log.New(os.Stdout, "[GLogger] ", log.LstdFlags), logLevel: Debug}
}

func (g *GLogger) Debug(value interface{}) {
	if g.logLevel >= Debug {
		g.logger.Println(value)
	}
}

func (g *GLogger) Info(value interface{}) {
	if g.logLevel >= Info {
		g.logger.Println(value)
	}
}

func (g *GLogger) Warn(value interface{}) {
	if g.logLevel >= Warn {
		g.logger.Println(value)
	}
}

func (g *GLogger) Error(value interface{}) {
	if g.logLevel >= Error {
		g.logger.Println(value)
	}
}

func (g *GLogger) Fatal(value interface{}) {
	if g.logLevel >= Fatal {
		g.logger.Fatalln(value)
	}
}

func (g *GLogger) Panic(value interface{}) {
	if g.logLevel == Panic {
		g.logger.Panicln(value)
	}
}
