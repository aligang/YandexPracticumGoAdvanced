package logging

import (
	"io"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func Configure(dst io.Writer, level zerolog.Level) {
	Logger = zerolog.New(dst).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(level)
}

func Crit(msg string, args ...any) {
	if len(args) > 0 {
		Logger.Error().Msgf(msg, args...)
	} else {
		Logger.Error().Msg(msg)
	}
}

func Warn(msg string, args ...any) {
	if len(args) > 0 {
		Logger.Warn().Msgf(msg, args...)
	} else {
		Logger.Warn().Msg(msg)
	}
}

func Debug(msg string, args ...any) {
	if len(args) > 0 {
		Logger.Debug().Msgf(msg, args...)
	} else {
		Logger.Debug().Msg(msg)
	}
}
