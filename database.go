package main

import (
	"database/sql"
	"time"
)

func ConnectToDB(connString string) (*sql.DB, error) {
	return sql.Open("postgres", connString)
}

func GetLatestInvoiceTime(db *sql.DB, accountId int) (time.Time, error) {
	var latestInvoiceTime time.Time
	query := "SELECT MAX(\"createdAt\") FROM payments WHERE account_id = $1"
	err := db.QueryRow(query, accountId).Scan(&latestInvoiceTime)
	return latestInvoiceTime, err
}

func GetNewDeposits(db *sql.DB, startTime time.Time) ([]string, error) {
	var newDeposits []string
	query := "SELECT id FROM \"KrakenDeposits\" WHERE \"createdAt\" >= $1"
	rows, err := db.Query(query, startTime)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var depositID string
		if err := rows.Scan(&depositID); err != nil {
			return nil, err
		}
		newDeposits = append(newDeposits, depositID)
	}

	return newDeposits, nil
}

