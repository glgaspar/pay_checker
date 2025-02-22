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

func db() (any, error) {
	var conn any
	// probably gonna use postgre
	return conn, nil
}

func getBillToCheck() (*[]Bill, error) {
	conn, err := db()
	if err != nil {
		return nil, err
	}
	//defer close missing

	var data []Bill
	query := `
	select id, expDay lastDate 
	from bills
	where 
	(
		-- current month > month(lastDate)
		-- and current year = year(lastDate)
	)
	or (
		-- current month < month(lastDate)
		-- and current year > year(lastDate)
	)
	or lastDate is null
	`
	// do the select
	return &data, nil
}

func updateReceipt() error {
	bills, err := getBillToCheck()
	if err != nil {
		return err
	}

	for b := 0; b < len(*bills); b++ {
		file, err := searchFile((*bills)[b])
		if err != nil {
			return err
		}
		if file {
			if err := updateFile((*bills)[b]); err != nil {
				return err
			}
		}
	}

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
