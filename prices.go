package main

import (
	"database/sql"
	"log"
	"time"
  "fmt"
)

const maxTimeWithoutUpdate = time.Minute * 10

func CheckPrices(db *sql.DB) {

	// Fetch the latest price update time from the prices table
	latestPriceTime, err := GetLatestPriceTime(db)
  log.Println(fmt.Sprintf("latest price time: %s"), latestPriceTime)
	if err != nil {
		log.Println("Error fetching latest price update:", err)
		return
	}

	// Check if the latest price update is more than ten minutes ago
	if time.Since(latestPriceTime) > maxTimeWithoutUpdate {
		SendRocketChatAlert("No price update in more than ten minutes!")
	}
}

func GetLatestPriceTime(db *sql.DB) (time.Time, error) {
	var latestPriceTime time.Time
	query := "SELECT MAX(\"updatedAt\") FROM prices"
	err := db.QueryRow(query).Scan(&latestPriceTime)
	if err != nil {
		log.Println("Error fetching latest price update time:", err)
	}
	return latestPriceTime, err
}

