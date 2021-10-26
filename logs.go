package lex

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = func(err error) interface{} {
		return string(debug.Stack())
	}
}

const (
	Method = "method"
	Module = "module"
	App    = "app"
	IP     = "ip"
)

func NewLog(filename string, writers ...io.Writer) zerolog.Logger {
	if filename != "" {
		logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
		}
		writers = append(writers, logFile)
	}

	multi := zerolog.MultiLevelWriter(writers...)
	log := zerolog.New(multi).With().Timestamp().Logger()
	return log
}

func GetLogger(logFile string, level string, withStdout bool) zerolog.Logger {
	if logFile != "" {
		if !filepath.IsAbs(logFile) {
			wd, err := os.Getwd()
			if err == nil {
				logFile = filepath.Join(wd, logFile)
			}
		}
		logDir := filepath.Dir(logFile)
		if _, err := os.Stat(logDir); err != nil {
			os.MkdirAll(logDir, 0755)
		}
	}
	var writers []io.Writer
	if withStdout {
		writers = append(writers, os.Stdout)
	}
	r := NewLog(logFile, writers...)
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		fmt.Println(err)
	}
	return r.Level(lvl)
}
