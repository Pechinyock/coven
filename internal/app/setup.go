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

/* file log doesn't work */
/* log options should be map like
-> console logger
-> file logger
could be dissabled both
both could have different level
*/
func setupLog(conf *config.LogConfig) error {
	if conf == nil {
		conf = config.DefaultLogSettings()
	}
	lvl := parseLogLevel(conf.LogLevel)
	logHandler := log.PrettyHandler{
		Level:  lvl,
		Writer: os.Stdout,
	}
	pth := conf.LogPath
	if pth != "" {
		fullPath, err := utils.GetFullPath(pth)
		if err != nil {
			return err
		}
		utils.CreatePath(fullPath)
		var file os.File
		if !utils.IsFilePath(fullPath) {
			nowDate := time.Now().UTC().Format("2006-01-02")
			logFileName := fmt.Sprintf("%s.%s", nowDate, ".log")
			fullPath = filepath.Join(fullPath, logFileName)
			created, err := os.Create(fullPath)
			if err != nil {
				return err
			}
			file = *created
		} else {
			existing, err := os.Open(fullPath)
			if err != nil {
				return err
			}
			file = *existing
		}
		fileHandler := log.FileLogHandler{
			Level:  lvl,
			Writer: &file,
		}

		mltHandler := log.NewMultiHandler(&logHandler, &fileHandler)
		logger := slog.New(mltHandler)
		slog.SetDefault(logger)
		slog.Info("logger has been initialized", "level", conf.LogLevel,
			"log file path", fullPath,
		)

		return nil
	}
	logger := slog.New(&logHandler)
	slog.SetDefault(logger)
	slog.Info("logger has been initialized", "level", conf.LogLevel)
	return nil
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
