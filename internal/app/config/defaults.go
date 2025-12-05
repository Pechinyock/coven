package config

func DefaultLogSettings() *LogConfig {
	return &LogConfig{
		ConsoleLogger: &ConoselLogger{
			Level: "debug",
		},
	}
}
