package app

import "fmt"

var (
	version     = "dev"
	gitShortSha = "none"
)

func Run() error {
	printIntialInfo()
	err := initializeLoging()
	return err
}

func PrintVersion() {
	fmt.Printf("version: %s\ngit commit sha: %s", version, gitShortSha)
}

func printIntialInfo() {
	fmt.Println("running coven web app")
	PrintVersion()
}
