package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	_timeFormat = "02 Jan 06 15:04"
	_maxSize    = 100
	_maxAge     = 30
	_maxBackups = 5
)

type Config struct {
	LogLevel           string `mapstructure:"log_level"`
	FileName           string `mapstructure:"file_name"`
	IsConsoleLogEnable bool   `mapstructure:"is_console_log_enable"`
	IsFileLogEnable    bool   `mapstructure:"is_file_log_enable"`
	IsCompress         bool   `mapstructure:"is_compress"`
}

type Logger interface {
	Fatal(err error, msg string, args ...interface{})
	Error(err error, msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Debug(err error, msg string, args ...interface{})
}

type Log struct {
	*zerolog.Logger
}

func New(cfg *Config) *Log {
	var l zerolog.Level

	switch cfg.LogLevel {
	default:
		l = zerolog.InfoLevel
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(l)
	zerolog.ErrorFieldName = "err"

	var writers []io.Writer

	if cfg.IsConsoleLogEnable {
		writers = append(writers, zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: _timeFormat,
		})
	}

	if cfg.IsFileLogEnable {
		writers = append(writers, &lumberjack.Logger{
			Filename:   cfg.FileName,
			MaxSize:    _maxSize,
			MaxAge:     _maxAge,
			MaxBackups: _maxBackups,
			Compress:   cfg.IsCompress,
		})
	}

	multi := zerolog.MultiLevelWriter(writers...)

	logger := zerolog.New(multi).
		With().
		Timestamp().
		Logger()

	return &Log{&logger}
}

func (l *Log) Debug(err error, msg string, args ...interface{}) {
	l.Logger.Debug().Err(err).Msgf(msg, args...)
}

func (l *Log) Info(msg string, args ...interface{}) {
	l.Logger.Info().Msgf(msg, args...)
}

func (l *Log) Warn(msg string, args ...interface{}) {
	l.Logger.Warn().Msgf(msg, args...)
}

func (l *Log) Error(err error, msg string, args ...interface{}) {
	l.Logger.Error().Err(err).Msgf(msg, args...)
}

func (l *Log) Fatal(err error, msg string, args ...interface{}) {
	l.Logger.Fatal().Err(err).Msgf(msg, args...)
}
