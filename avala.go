package main

import (
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.dev")

	if err != nil {
		panic("Env is not loaded")
	}
}
