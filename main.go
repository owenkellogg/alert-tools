package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	pgConnStringEnvKey = "PG_CONN_STRING"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Get environment variables
	pgConnString := os.Getenv(pgConnStringEnvKey)

	// Connect to PostgreSQL database
	db, err := ConnectToDB(pgConnString)
	if err != nil {
		log.Fatal("Error connecting to PostgreSQL:", err)
	}
	defer db.Close()

	// Run the main loop to check for new deposits
	for {
		// Fetch the latest invoice paid by a particular account
		latestInvoiceTime, err := GetLatestInvoiceTime(db, "account_name")
		if err != nil {
			log.Println("Error fetching latest invoice:", err)
		} else {
			// Calculate the time one hour after the latest invoice
			oneHourAfterInvoice := latestInvoiceTime.Add(time.Hour)

			// Check for new deposits within one hour after the latest invoice
			newDeposits, err := GetNewDeposits(db, oneHourAfterInvoice)
			if err != nil {
				log.Println("Error fetching new deposits:", err)
			} else {
				// If there are no new deposits, send an alert to Rocket.Chat
				if len(newDeposits) == 0 {
					SendRocketChatAlert()
				}
			}
		}

		// Sleep for a minute before checking again
		time.Sleep(time.Minute)
	}
}

