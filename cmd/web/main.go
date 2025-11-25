package main

import (
	coven "coven/internal/app"
	"fmt"
)

func main() {
	err := coven.Run()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to run conve web: %s", err.Error()))
	}
}
