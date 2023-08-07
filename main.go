package main

import (
  "fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
  "strconv"
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
    accountId, err := strconv.Atoi(os.Getenv("account_id"))
		latestInvoiceTime, err := GetLatestInvoiceTime(db, accountId)
    log.Println("Latest invoice time fetched")
		if err != nil {
			log.Println("Error fetching latest invoice:", err)
		} else {
			// Calculate the time one hour after the latest invoice
			oneHourAfterInvoice := latestInvoiceTime.Add(time.Hour)
      log.Println(oneHourAfterInvoice)

			// Check for new deposits within one hour after the latest invoice
			newDeposits, err := GetNewDeposits(db, oneHourAfterInvoice)
			if err != nil {
				log.Println("Error fetching new deposits:", err)
			} else {
				// If there are no new deposits, send an alert to Rocket.Chat
        log.Println(fmt.Sprintf("%d new deposits detected since latest invoice paid ", len(newDeposits)))
				if len(newDeposits) == 0 {
          log.Println("No new deposits, send rocketchat alert now")
					SendRocketChatAlert("No new deposits in Kraken one hour after last paid invoice")
				}
			}
		}

    CheckPrices(db)

		// Sleep for a minute before checking again
		time.Sleep(time.Minute)
	}
}

