package loggerext

type LoggerInterface interface {
	Info(mm ...interface{})
	Infof(s string, mm ...interface{})
	Debug(mm ...interface{})
	Debugf(s string, mm ...interface{})
	Error(mm ...interface{})
	Errorf(s string, mm ...interface{})
}
