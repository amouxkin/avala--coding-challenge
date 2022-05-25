package main

import (
	"github.com/joho/godotenv"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(10)
	err := godotenv.Load(".env.dev")

	if err != nil {
		panic("Env is not loaded")
	}
}
