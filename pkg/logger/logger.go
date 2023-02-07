// NOTE: SINGELTON
// TODO: Interface
package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	*zerolog.Logger
}

var (
	once   sync.Once
	logger zerolog.Logger
)

func New(level string) *Logger {
	once.Do(func() {
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

		logger = zerolog.New(output).
			With().
			Timestamp().
			Logger()
	})

	return &Logger{&logger}
}
