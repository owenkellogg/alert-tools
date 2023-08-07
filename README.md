Sure! Here's a full `README.md` with instructions on building and operating the program:

# Alert Tools

Alert Tools is a Go (Golang) program that connects to a PostgreSQL database to query for new Kraken deposits. If there are no new deposits within one hour after the latest invoice paid by a particular account, it sends an alert to a Slack channel using a Slack webhook.

## Prerequisites

Before running the program, make sure you have the following prerequisites installed:

- Go (Golang) 1.17 or later: [https://golang.org/dl/](https://golang.org/dl/)
- PostgreSQL: [https://www.postgresql.org/download/](https://www.postgresql.org/download/)

## Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/owenkellogg/alert-tools.git
   ```

2. Navigate to the project directory:

   ```bash
   cd alert-tools
   ```

3. Initialize the Go module:

   ```bash
   go mod init github.com/owenkellogg/alert-tools
   ```

4. Install the required Go packages:

   ```bash
   go get github.com/lib/pq
   go get github.com/nlopes/slack
   go get github.com/joho/godotenv
   ```

5. Create a `.env` file in the root of the project and add your sensitive information:

   ```plaintext
   SLACK_WEBHOOK_URL=your_slack_webhook_url
   PG_CONN_STRING=postgresql://user:password@localhost/dbname?sslmode=disable
   ```

   Replace `your_slack_webhook_url` with your actual Slack webhook URL and `postgresql://user:password@localhost/dbname?sslmode=disable` with your PostgreSQL connection string.

6. Build the program:

   ```bash
   go build
   ```

7. Run the program:

   ```bash
   ./alert-tools
   ```

   The program will connect to the PostgreSQL database and monitor for new deposits. If there are no new deposits within one hour after the latest invoice for the specified account, it will send an alert to the Slack channel using the specified webhook URL.

## Operating the Program

The Kraken Deposit Monitor is designed to run continuously and periodically check for new deposits. It follows the one-minute sleep interval between checks.

To stop the program, you can use `Ctrl+C` in the terminal where it is running.

## Configuring the Program

You can modify the following constants in the `main.go` file to customize the behavior of the program:

- `slackWebhookURL`: Replace this with your actual Slack webhook URL.
- `pgConnString`: Replace this with your PostgreSQL connection string.

You can also adjust the sleep interval in the main loop to check for new deposits more frequently or less frequently, depending on your requirements.

## Contributing

If you find any issues or have ideas for improvements, feel free to open an issue or submit a pull request. Your contributions are welcome!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- The program uses the [github.com/lib/pq](https://github.com/lib/pq) package for PostgreSQL connectivity.
- The program uses the [github.com/nlopes/slack](https://github.com/nlopes/slack) package for sending Slack webhooks.
- The program uses the [github.com/joho/godotenv](https://github.com/joho/godotenv) package for handling environment variables from a `.env` file.

---
Please make sure to replace `owenkellogg` and `your_slack_webhook_url` with your actual GitHub username and Slack webhook URL in the `README.md` file.

This README.md file provides users with instructions on how to set up and use the Kraken Deposit Monitor program, along with details on configuration, contributing, and licensing. Feel free to add more information as needed to make it more comprehensive for your specific use case.

