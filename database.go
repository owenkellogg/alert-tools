package main

import (
	"database/sql"
	"time"
)

func ConnectToDB(connString string) (*sql.DB, error) {
	return sql.Open("postgres", connString)
}

func GetLatestInvoiceTime(db *sql.DB, accountName string) (time.Time, error) {
	var latestInvoiceTime time.Time
	query := "SELECT MAX(invoice_time) FROM invoices WHERE account_name = $1"
	err := db.QueryRow(query, accountName).Scan(&latestInvoiceTime)
	return latestInvoiceTime, err
}

func GetNewDeposits(db *sql.DB, startTime time.Time) ([]string, error) {
	var newDeposits []string
	query := "SELECT deposit_id FROM deposits WHERE deposit_time >= $1"
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

