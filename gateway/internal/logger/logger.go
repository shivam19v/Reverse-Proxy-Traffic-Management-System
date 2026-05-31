package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// Log is global logger instance
var Log zerolog.Logger

func init() {

	// Output logs to console
	Log = zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()
}