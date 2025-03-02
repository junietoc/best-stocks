package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}

	database, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	fmt.Println("Connected to CockroachDB successfully!")
	DB = database
}

type Stock struct {
	ID        string `json:"id"`
	Ticker    string `json:"ticker"`
	Company   string `json:"company"`
	Action    string `json:"action"`
	Brokerage string `json:"brokerage"`
	TargetTo  string `json:"target_to"`
}

func GetStocks() ([]Stock, error) {
	rows, err := DB.Query(`
        SELECT id, ticker, company, action, brokerage, target_to
        FROM stocks
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []Stock
	for rows.Next() {
		var s Stock
		if err := rows.Scan(&s.ID, &s.Ticker, &s.Company, &s.Action, &s.Brokerage, &s.TargetTo); err != nil {
			return nil, err
		}
		stocks = append(stocks, s)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for i, s := range stocks {
		if i >= 10 {
			break
		}
		log.Println(s)
	}
	return stocks, nil
}
