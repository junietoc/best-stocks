package stocks

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func SaveStockItems(db *sql.DB, items []StockItem) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback()

	log.Println("saving this thingg")

	stmt, err := tx.Prepare(`
        INSERT INTO stocks (
            ticker,
            target_from,
            target_to,
            company,
            action,
            brokerage,
            rating_from,
            rating_to,
            time
        )
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        ON CONFLICT (ticker)
        DO UPDATE SET
            target_from = EXCLUDED.target_from,
            target_to = EXCLUDED.target_to,
            company = EXCLUDED.company,
            action = EXCLUDED.action,
            brokerage = EXCLUDED.brokerage,
            rating_from = EXCLUDED.rating_from,
            rating_to = EXCLUDED.rating_to,
            time = EXCLUDED.time
    `)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	for _, item := range items {
		parsedTime, parseErr := time.Parse(time.RFC3339Nano, item.Time)
		if parseErr != nil {
			// Decide how you want to handle a parsing error:
			// return fmt.Errorf("parse time: %w", parseErr)
			// or skip this row, etc.
		}

		_, execErr := stmt.Exec(
			item.Ticker,
			item.TargetFrom,
			item.TargetTo,
			item.Company,
			item.Action,
			item.Brokerage,
			item.RatingFrom,
			item.RatingTo,
			parsedTime,
		)
		if execErr != nil {
			return fmt.Errorf("inserting/updating item: %w", execErr)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}
