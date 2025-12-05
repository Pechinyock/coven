package main

import (
	"coven/internal/app"

	"fmt"
)

func main() {
	config, err := app.Init()
	if err != nil {
		fmt.Println("failed to initialize app:")
		fmt.Println(fmt.Errorf("failed to run conven web: %s", err.Error()))
		return
	}
	app.Run(config)
}
