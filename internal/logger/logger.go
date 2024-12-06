package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
	"log"
	"path/filepath"
	"search-nova/internal/config"
	"search-nova/internal/constant"
)

var (
	L *logrus.Logger
)

const (
	maxSize    = 500
	maxBackups = 3
	maxAge     = 7
	compress   = true
	file       = "search-nova.log"
)

func init() {
	L = logrus.New()
	level, err := logrus.ParseLevel(config.C.GetString(constant.LoggerLevel))
	if err != nil {
		log.Fatalf("Fatal error logger : %v\n", err)
	}
	fn := filepath.Join(config.C.GetString(constant.LoggerDir), file)
	L.Out = &lumberjack.Logger{
		Filename:   fn,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	L.SetLevel(level)
	L.Formatter = &logrus.JSONFormatter{}
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetOutput(L.WriterLevel(logrus.InfoLevel))
}
