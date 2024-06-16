package slog

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/go-kratos/kratos/v2/log"
)

// see [log/slog.Level] on why 15 was chosen here
const LevelFatal slog.Level = 15

var _ log.Logger = (*Logger)(nil)

type Logger struct {
	log    *slog.Logger
	msgKey string
}

func NewLogger(handlerOpts slog.Handler) *Logger {
	return &Logger{
		log:    slog.New(handler),
		msgKey: log.DefaultMessageKey,
	}
}

// Log implements log.Logger.
func (l *Logger) Log(level log.Level, keyvals ...interface{}) error {
	var (
		msg    = ""
		keylen = len(keyvals)
	)
	if keylen == 0 || keylen%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}

	data := make([]slog.Attr, 0, (keylen/2)+1)
	for i := 0; i < keylen; i += 2 {
		if key, ok := keyvals[i].(string); ok && key == l.msgKey {
			continue
		}
		data = append(data, slog.Any(fmt.Sprint(keyvals[i]), keyvals[i+1]))
	}
	ctx := context.Background()

	switch level {
	case log.LevelDebug:
		l.log.LogAttrs(ctx, slog.LevelDebug, msg, data...)
	case log.LevelInfo:
		l.log.LogAttrs(ctx, slog.LevelInfo, msg, data...)
	case log.LevelWarn:
		l.log.LogAttrs(ctx, slog.LevelWarn, msg, data...)
	case log.LevelError:
		l.log.LogAttrs(ctx, slog.LevelError, msg, data...)
	case log.LevelFatal:
		// There is no Fatal level in slog
		l.log.LogAttrs(ctx, LevelFatal, msg, data...)
		os.Exit(1)
	}
	return nil
}
