package main

import "github.com/mrborghini/go-logger"

func main() {
	log := logger.NewLogger("main")
	log.Info("Running...")
}