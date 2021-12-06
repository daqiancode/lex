package lex

import (
	"io"
	"os"
	"runtime/debug"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return string(debug.Stack())
	}
}

// maxSize:unit megabytes, maxAge:unit days
func NewRollingWriter(fileName string, maxSize, maxAge, maxBackups int, compress, localtime bool) io.Writer {
	return &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge,   //days
		Compress:   compress, // disabled by default
		LocalTime:  localtime,
	}
}

func NewFileWriter(fileName string) (io.Writer, error) {
	return os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
}

func NewStdoutWriter() io.Writer {
	return zerolog.ConsoleWriter{Out: os.Stdout}
}

func NewStderrWriter() io.Writer {
	return zerolog.ConsoleWriter{Out: os.Stderr}
}

func NewLogger(tags map[string]string, writers ...io.Writer) zerolog.Logger {
	multi := zerolog.MultiLevelWriter(writers...)
	ctx := zerolog.New(multi).With()
	for k, v := range tags {
		ctx = ctx.Str(k, v)
	}
	return ctx.Timestamp().Logger()
}
