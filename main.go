package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Bill struct {
	Id     int       `json:"id" db:"id"`         // Bill Id
	ExpDay int       `json:"expDay" db:"expDay"` // Expiration day
	Date   time.Time `json:"date" db:"date"`     // Time of last payment
	Path   string    `json:"path" db:"path"`     // Where to find the files
}

func db() (*sql.DB, error) {
	HOST := os.Getenv("PG_HOST")
	PORT := os.Getenv("PG_PORT")
	USER := os.Getenv("PG_USER")
	PASSWORD := os.Getenv("PG_PASSWORD")
	DBNAME := os.Getenv("PG_DBNAME")

	conn, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			HOST, PORT, USER, PASSWORD, DBNAME),
	)

	if err != nil {
		return nil, err
	}
	err = conn.Ping()

	return conn, err
}

func getBillToCheck() (*[]Bill, error) {
	conn, err := db()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

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

	result, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var d Bill
		if err = result.Scan(&d.Id, &d.Description, &d.ExpDay, &d.LastDate, &d.Path); err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	return &data, nil
}

func searchFile(bill Bill) (bool, error) {
	// not today
	return true, nil
}

func updateFile(bill Bill) error {
	conn, err := db()
	if err != nil {
		return err
	}
	defer conn.Close()

	var data []Bill
	query := `
	update bills
	set lastDate = $1
	where id = $2
	`
	if _, err := conn.Exec(query, time.Now(), (*bill).Id); err != nil {
		return err
	}

	return nil
}

func updateReceipt() error {
	bills, err := getBillToCheck()
	if err != nil {
		return err
	}

	for b := range *bills {
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
