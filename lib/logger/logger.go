package logger

import (
	"api-go/lib/config"
	"api-go/lib/file"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Setup() {
	var err error
	filePath := config.LoggerConfig.SavePath
	fileName := fmt.Sprintf("%s%s.%s",
		config.LoggerConfig.SaveName,
		time.Now().Format(config.LoggerConfig.TimeFormat),
		config.LoggerConfig.FileExt,
	)

	F, err = file.OpenExistFile(fileName, filePath)
	if err != nil {
		log.Fatalln(err)
	}

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

// setPrefix 日志级别
func setPrefix(level Level) {
	_, logFile, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]",
			levelFlags[level],
			filepath.Base(logFile),
			line,
		)
	} else {
		logPrefix = fmt.Sprintf("[%s]",
			levelFlags[level],
		)
	}
	logger.SetPrefix(logPrefix)
}
