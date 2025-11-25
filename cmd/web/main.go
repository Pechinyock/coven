package main

import (
	coven "coven/internal/app"
	"fmt"
)

func main() {
	err := coven.Init()
	if err != nil {
		fmt.Println(fmt.Errorf("failed to run conven web: %s", err.Error()))
	}
}
