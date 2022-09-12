package logging

import (
	"github.com/rs/zerolog"
	"io"
)

var Logger zerolog.Logger

//type Logger struct {
//
//}

//func Warning(msg string) {
//	Logger.Warn().Msg()
//}

func Configure(dst io.Writer, level zerolog.Level) {
	Logger = zerolog.New(dst).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(level)
}
