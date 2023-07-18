package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func factory(conf Config) *zap.Logger {
	var core Core
	switch conf.Driver {
	case "daily":
		core = NewRotateCore(formatter(conf), conf)
	case "console":
		core = NewConsoleCore(formatter(conf), zapLevel(conf.Level))
	}

	return zap.New(
		core.core(),
		zap.AddCaller(),
		//zap.AddCallerSkip(2),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}

func formatter(conf Config) Formatter {
	switch conf.Formatter {
	case "json":
		return jsonFormatter(conf)
	case "console":
		return consoleFormatter(conf)
	}
	return nil
}

func consoleFormatter(conf Config) Formatter {
	return &ConsoleFormatter{
		format: conf.Format,
	}
}

func jsonFormatter(conf Config) Formatter {
	return &JSONFormatter{
		format: conf.Format,
	}
}

func zapLevel(level string) zapcore.Level {
	var lv zapcore.Level

	if err := lv.UnmarshalText([]byte(level)); err != nil {
		panic("get log lv error: " + fmt.Sprintf("%#v", err))
	}
	return lv
}
