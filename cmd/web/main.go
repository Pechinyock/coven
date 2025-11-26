package main

import (
	"coven/internal/app"

	"fmt"
)

func main() {
	config, err := app.Init()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to run conven web: %s", err.Error()))
	}
	app.Run(config)
}
