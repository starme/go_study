package log

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"star/pkg/config"
	"sync"
)

var (
	manager *Manager
	logOnce sync.Once
)

type Config struct {
	Driver     string
	Name       string
	Level      string
	Path       string
	Days       int
	Formatter  string
	DateFormat string
	Format     map[string]interface{}
}

func init() {
	conf := config.Sub("logging")
	if conf == nil {
		panic("log config not found")
	}
	manager = NewManager()
	manager.DefaultChannel = conf.GetString("default")
	channels := parseChannelConfig(conf)
	for _, channel := range channels {
		manager.AddChannel(channel.Name, channel)
	}
}

func NewManager() *Manager {
	return &Manager{
		Channels: make(map[string]*zap.Logger),
	}
}

// Manager structure
type Manager struct {
	DefaultChannel string
	Channels       map[string]*zap.Logger
}

func (m *Manager) AddChannel(name string, conf Config) {
	m.Channels[name] = factory(conf)
}

func (m *Manager) Channel(name string) *zap.Logger {
	return m.Channels[name]
}

func parseChannelConfig(conf *viper.Viper) []Config {
	var c []Config
	channels := conf.GetStringMap("channels")
	if channels == nil {
		panic("log config channels not found")
	}

	for name := range channels {
		var channel Config
		if err := conf.Sub("channels." + name).Unmarshal(&channel); err != nil {
			continue
		}
		channel.Name = name
		c = append(c, channel)
	}
	return c
}

func Channel(s string) *zap.Logger {
	return manager.Channel(s)
}

func Debug(msg string, fields ...zap.Field) {
	Channel(manager.DefaultChannel).WithOptions(zap.AddCallerSkip(1)).Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Channel(manager.DefaultChannel).WithOptions(zap.AddCallerSkip(1)).Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Channel(manager.DefaultChannel).WithOptions(zap.AddCallerSkip(1)).Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Channel(manager.DefaultChannel).WithOptions(zap.AddCallerSkip(1)).Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	Channel(manager.DefaultChannel).WithOptions(zap.AddCallerSkip(1)).DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Channel(manager.DefaultChannel).WithOptions(zap.AddCallerSkip(1)).Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Channel(manager.DefaultChannel).WithOptions(zap.AddCallerSkip(1)).Fatal(msg, fields...)
}
