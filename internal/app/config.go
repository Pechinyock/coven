package app

import (
	"coven/internal/app/config"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var configPath = flag.String("config", "config.json", "path to config file")

func readConfig() (*config.CovenWebConfig, error) {
	if !flag.Parsed() {
		flag.Parse()
	}
	cfgPath := *configPath

	fmt.Printf("reading configuration from file %s\n", cfgPath)

	data, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var config config.CovenWebConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
