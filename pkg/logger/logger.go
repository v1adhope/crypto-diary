package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

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

func New(level string) *Log {
	var l zerolog.Level

	switch level {
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

	var output io.Writer = zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC822Z,
	}

	logger := zerolog.New(output).
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
