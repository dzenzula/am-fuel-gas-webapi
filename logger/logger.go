package logger

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const log_suffix string = ".log"
const logs_directory string = "logs" + string(os.PathSeparator)

var logger *slog.Logger
var file File

type File struct {
	name, path, fullPath string
}

func LogInit(logLevel slog.Level) {
	createLogFile()
	logger = slog.New(slog.Default().Handler())
	slog.SetLogLoggerLevel(logLevel)
	logger.Info("Logger initialized")
	logger.Debug("DebugMode")
}

func Info(msg string, args ...any) {
	if !checkLogExist(file.fullPath) {
		createLogFile()
	}
	logger.Info(msg, args...)
}

func Debug(msg string, args ...any) {
	if !checkLogExist(file.fullPath) {
		createLogFile()
	}
	logger.Debug(msg, args...)
}

func Error(msg string, args ...any) {
	if !checkLogExist(file.fullPath) {
		createLogFile()
	}
	logger.Error(msg, args...)
}

func Warn(msg string, args ...any) {
	if !checkLogExist(file.fullPath) {
		createLogFile()
	}
	logger.Warn(msg, args...)
}

func createLogFile() {
	file = newFile()

	if _, err := os.Stat(file.path); os.IsNotExist(err) {
		err := os.Mkdir(file.path, 0777)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	logFile, err := os.OpenFile(file.fullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}

	log.SetOutput(logFile)
}

func newFile() File {
	var f = File{}
	date := time.Now().Format("20060102")
	f.name = os.Args[0]
	f.name = filepath.Base(f.name)
	if runtime.GOOS == "windows" {
		var ext string = filepath.Ext(f.name)
		f.name = strings.TrimSuffix(f.name, ext)
	}
	f.path = filepath.Dir(f.name)
	f.path, _ = filepath.Abs(f.path)
	f.path = f.path + string(os.PathSeparator) + logs_directory
	f.fullPath = f.path + date + "_" + f.name + log_suffix
	return f
}

func checkLogExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}
