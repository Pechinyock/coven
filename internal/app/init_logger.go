package app

import (
	"coven/internal/app/config"
	"coven/internal/log"
	"coven/internal/utils"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func setupLogger(conf *config.LogConfig) {
	if conf == nil {
		conf = config.DefaultLogSettings()
		fmt.Println("log options was not provided via config file, setting it to default:")
		fmt.Printf("write to stdout, level: '%s'", conf.ConsoleLogger.Level)
	}
	handlers := []slog.Handler{}

	if conf.ConsoleLogger != nil {
		lvl := parseLogLevel(conf.ConsoleLogger.Level)
		consoleHandler := &log.ColoreLabelHandler{
			Level:  lvl,
			Writer: os.Stdout,
		}
		handlers = append(handlers, consoleHandler)
	}
	if conf.FileLogger != nil {
		file, err := getLogFile(conf.FileLogger.FilePath)
		if err != nil {
			fmt.Printf("failed to create file logger: '%s'", err.Error())
			fmt.Println("service will start without file logger")
		} else {
			lvl := parseLogLevel(conf.FileLogger.Level)
			fileHandler := &log.SimpleLogHandler{
				Level:  lvl,
				Writer: file,
			}
			handlers = append(handlers, fileHandler)
		}
	} else {
		fmt.Println("file logger is disabled")
	}

	mlthandler := log.NewMultiHandler(handlers...)
	logger := slog.New(mlthandler)
	slog.SetDefault(logger)
	slog.Info("logger succesfully initialized")
}

func parseLogLevel(logLevelStr string) slog.Level {
	lower := strings.ToLower(logLevelStr)
	switch lower {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "err", "error":
		return slog.LevelError
	default:
		{
			fmt.Printf("failed to parse log level: %s\n", logLevelStr)
			fmt.Printf("log level will be set to default: 'info'\n")
			return slog.LevelInfo
		}
	}
}

func getLogFile(path string) (*os.File, error) {
	fullPath, err := utils.GetFullPath(path)
	if err != nil {
		return nil, err
	}
	info, err := os.Stat(fullPath)

	if err == nil && info.IsDir() {
		fileName := fmt.Sprintf("%s.%s", time.Now().Format("2006_01_02"), "log")
		fullPath := filepath.Join(fullPath, fileName)
		return os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	}

	if os.IsNotExist(err) {
		if !utils.IsFilePath(fullPath) {
			if err := os.MkdirAll(fullPath, 0755); err != nil {
				return nil, err
			}
			fileName := fmt.Sprintf("%s.%s", time.Now().Format("2006_01_02"), "log")
			fullPath := filepath.Join(fullPath, fileName)
			return os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		} else {
			dir := filepath.Dir(fullPath)
			if err := os.MkdirAll(dir, 0755); err != nil {
				return nil, err
			}
			return os.OpenFile(fullPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		}
	}

	if err == nil && !info.IsDir() {
		return os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	}

	return nil, err
}
