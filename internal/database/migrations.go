package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "./auctionhouse.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	createAuctionTable := `
	CREATE TABLE IF NOT EXISTS auctions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		item_name TEXT NOT NULL,
		description TEXT,
		end_time TIMESTAMP NOT NULL,
		starting_bid REAL NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := DB.Exec(createAuctionTable)
	if err != nil {
		log.Fatal(err)
	}

	createBidsTable := `
	CREATE TABLE IF NOT EXISTS bids (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		auction_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		amount REAL NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (auction_id) REFERENCES auctions (id),
		FOREIGN KEY (user_id) REFERENCES users (id)
	);
	`
	_, err = DB.Exec(createBidsTable)
	if err != nil {
		log.Fatal(err)
	}

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT UNIQUE NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
	`
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}
}
