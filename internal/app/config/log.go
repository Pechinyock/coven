package config

import "fmt"

type ConoselLogger struct {
	Level string `json:"level"`
}

type FileLogger struct {
	Level    string `json:"level"`
	FilePath string `json:"filePath"`
}

type LogConfig struct {
	ConsoleLogger *ConoselLogger `json:"consoleLogger"`
	FileLogger    *FileLogger    `json:"fileLogger"`
}

func (l *LogConfig) Validate() error {
	fmt.Println("log config validation is not implemented")
	return nil
}
