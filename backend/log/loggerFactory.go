package log

import (
	"go.uber.org/zap"
	"runtime"
	"runtime/debug"
	"serviceAuth/backend/config"
	"time"
)

func getCaller() string {
	var caller string

	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		caller = details.Name()
	}

	return caller
}

func Info(msg interface{}) {
	config.GetLogger().Info("INFO",
		zap.Time("time", time.Now()),
		zap.Any("message", msg))
}

func Warn(msg interface{}) {
	config.GetLogger().Warn("WARN",
		zap.Time("time", time.Now()),
		zap.Any("message", msg))
}

func Error(msg interface{}) {
	config.GetLogger().Error("Error Found",
		zap.Time("time", time.Now()),
		zap.Any("error", msg),
		zap.String("stack", string(debug.Stack())))
}

func Panic(msg interface{}) {
	config.GetLogger().Panic("PANIC",
		zap.Time("time", time.Now()),
		zap.Any("message", msg))
}

func Debug(msg interface{}) {

	config.GetLogger().Debug("DEBUG",
		zap.Time("time", time.Now()),
		zap.Any("message", msg),
		zap.Any("caller", getCaller()))

}
