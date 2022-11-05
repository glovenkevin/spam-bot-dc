package core

import (
	"fmt"
	"log"
	"os"
)

type CoreLogger struct {
	level    LogLevel
	infoLog  *log.Logger
	debugLog *log.Logger
	errorLog *log.Logger
}

func NewCoreLogger(l LogLevel) *CoreLogger {
	cl := &CoreLogger{
		level:    l,
		infoLog:  log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile),
		debugLog: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile),
		errorLog: log.New(os.Stdout, "[ERROR] ", log.LstdFlags|log.Lshortfile),
	}
	return cl
}

func (c *CoreLogger) Info(mm ...interface{}) {
	c.infoLog.Output(2, fmt.Sprintln(mm...))
}
func (c *CoreLogger) Infof(s string, mm ...interface{}) {
	c.infoLog.Output(2, fmt.Sprintf(s, mm...))
}
func (c *CoreLogger) Debug(mm ...interface{}) {
	if c.level >= LevelDebug {
		c.debugLog.Output(2, fmt.Sprintln(mm...))
	}
}
func (c *CoreLogger) Debugf(s string, mm ...interface{}) {
	if c.level >= LevelDebug {
		c.debugLog.Output(2, fmt.Sprintf(s, mm...))
	}
}
func (c *CoreLogger) Error(mm ...interface{}) {
	c.errorLog.Output(2, fmt.Sprintln(mm...))
}
func (c *CoreLogger) Errorf(s string, mm ...interface{}) {
	c.errorLog.Output(2, fmt.Sprintf(s, mm...))
}
