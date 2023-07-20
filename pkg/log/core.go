package log

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

type Core interface {
	core() zapcore.Core
}

func NewRotateCore(formatter Formatter, conf Config) Core {
	return &RotateCore{
		formatter: formatter,
		level:     zapLevel(conf.Level),
		handler: RotateHandler{
			filename: conf.Path,
			daily:    conf.Days,
		},
	}
}

type RotateCore struct {
	handler   RotateHandler
	formatter Formatter
	level     zapcore.Level
}

func (c *RotateCore) core() zapcore.Core {
	return zapcore.NewCore(c.formatter.encoder(), c.handler.syncer(), c.level)
}

func NewConsoleCore(formatter Formatter, level zapcore.Level) Core {
	return &ConsoleCore{
		formatter: formatter,
		level:     level,
		stdout:    StdoutHandler{},
		stderr:    StderrHandler{},
	}
}

type ConsoleCore struct {
	stdout    StdoutHandler
	stderr    StderrHandler
	formatter Formatter
	level     zapcore.Level
}

func (c *ConsoleCore) core() zapcore.Core {
	return zapcore.NewTee(
		zapcore.NewCore(c.formatter.encoder(), c.stdout.syncer(), zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l <= c.level
		})),
		zapcore.NewCore(c.formatter.encoder(), c.stderr.syncer(), zap.LevelEnablerFunc(func(l zapcore.Level) bool {
			return l > c.level
		})),
	)
}

type Handler interface {
	syncer() zapcore.WriteSyncer
}

type RotateHandler struct {
	filename string
	daily    int
}

func (h *RotateHandler) syncer() zapcore.WriteSyncer {
	if h.daily == 0 {
		h.daily = 1
	}
	w, err := rotateLogs.New(
		strings.Replace(h.filename, ".log", "-%Y-%m-%d.log", 1),
		rotateLogs.WithLinkName(h.filename),
		rotateLogs.WithMaxAge(time.Duration(h.daily)*24*time.Hour),
		rotateLogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return zapcore.AddSync(w)
}

type StdoutHandler struct {
}

func (h *StdoutHandler) syncer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

type StderrHandler struct {
}

func (h *StderrHandler) syncer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stderr)
}

type Formatter interface {
	encoder() zapcore.Encoder
}

type JSONFormatter struct {
	format map[string]interface{}
}

func (f *JSONFormatter) encoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(buildEncoderConfig(f.format))
}

type ConsoleFormatter struct {
	format map[string]interface{}
}

func (f *ConsoleFormatter) encoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(buildEncoderConfig(f.format))
}

func buildEncoderConfig(conf map[string]interface{}) zapcore.EncoderConfig {
	encodeConf := zapcore.EncoderConfig{}
	if mod, ok := conf["mod"]; ok {
		switch mod {
		case "production":
			encodeConf = zap.NewProductionEncoderConfig()
		case "development":
			encodeConf = zap.NewDevelopmentEncoderConfig()
		}
	}
	if levelKey, ok := conf["level_key"]; ok {
		encodeConf.LevelKey = levelKey.(string)
	}
	if timeKey, ok := conf["time_key"]; ok {
		encodeConf.TimeKey = timeKey.(string)
	}
	if nameKey, ok := conf["name_key"]; ok {
		encodeConf.NameKey = nameKey.(string)
	}
	if callerKey, ok := conf["caller_key"]; ok {
		encodeConf.CallerKey = callerKey.(string)
	}
	if stacktraceKey, ok := conf["stacktrace_key"]; ok {
		encodeConf.StacktraceKey = stacktraceKey.(string)
	}
	if lineEnding, ok := conf["line_ending"]; ok {
		encodeConf.LineEnding = lineEnding.(string)
	}
	if levelEncoder, ok := conf["level_encoder"]; ok {
		switch levelEncoder {
		case "capital":
			encodeConf.EncodeLevel = zapcore.CapitalLevelEncoder
		case "capitalColor":
			encodeConf.EncodeLevel = zapcore.CapitalColorLevelEncoder
		case "lowercase":
			encodeConf.EncodeLevel = zapcore.LowercaseLevelEncoder
		case "lowercaseColor":
			encodeConf.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		}
	}
	if timeEncoder, ok := conf["time_encoder"]; ok {
		encodeConf.EncodeTime = zapcore.TimeEncoderOfLayout(timeEncoder.(string))
	}
	if durationEncoder, ok := conf["duration_encoder"]; ok {
		switch durationEncoder {
		case "string":
			encodeConf.EncodeDuration = zapcore.StringDurationEncoder
		case "nanos":
			encodeConf.EncodeDuration = zapcore.NanosDurationEncoder
		case "ms":
			encodeConf.EncodeDuration = zapcore.MillisDurationEncoder
		case "seconds":
			encodeConf.EncodeDuration = zapcore.SecondsDurationEncoder
		}
	}
	if callerEncoder, ok := conf["caller_encoder"]; ok {
		switch callerEncoder {
		case "short":
			encodeConf.EncodeCaller = zapcore.ShortCallerEncoder
		case "full":
			encodeConf.EncodeCaller = zapcore.FullCallerEncoder
		}
	}
	if nameEncoder, ok := conf["name_encoder"]; ok {
		switch nameEncoder {
		case "short":
			encodeConf.EncodeName = zapcore.FullNameEncoder
		case "full":
			encodeConf.EncodeName = zapcore.FullNameEncoder
		}
	}
	if consoleSeparator, ok := conf["console_separator"]; ok {
		encodeConf.ConsoleSeparator = consoleSeparator.(string)
	}
	return encodeConf
}
