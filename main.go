package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Bill struct {
	Id     int       `json:"id" db:"id"`         // Bill Id
	ExpDay int       `json:"expDay" db:"expDay"` // Expiration day
	Date   time.Time `json:"date" db:"date"`     // Time of last payment
	Path   string    `json:"path" db:"path"`     // Where to find the files
}


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
