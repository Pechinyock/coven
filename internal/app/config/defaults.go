package config

func DefaultLogSettings() *LogConfig {
	return &LogConfig{
		LogLevel: "info",
	}
}
