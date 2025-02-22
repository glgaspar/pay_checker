package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func updateReceipt() error {
	// do something
	return nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %s\n", err)
	}

	sleepMinutes, err := strconv.Atoi(os.Getenv("SLEEP_MINUTE"))
	if err != nil {
		panic(err)
	}

	for {
		if err := updateReceipt(); err != nil {
			panic(err)
		}
		time.Sleep(time.Minute * time.Duration(sleepMinutes))
	}
}
