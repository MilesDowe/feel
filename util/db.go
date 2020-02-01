package util

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // database driver
	"log"
	"os/user"
)

const createTable = "CREATE TABLE IF NOT EXISTS feel_recording (" +
	"id INTEGER PRIMARY KEY," +
	"score INTEGER," +
	"concern TEXT NULLABLE," +
	"grateful TEXT NULLABLE," +
	"learn TEXT NULLABLE," +
	"milestone TEXT NULLABLE," +
	"entered INTEGER)"

const deleteRecordPerID = "DELETE FROM feel_recording WHERE id = ?"

var databaseLoc = ""

// OpenDb : returns a connection to the SQLite database
func OpenDb() *sql.DB {
	if databaseLoc == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		databaseLoc = usr.HomeDir + "/go/src/github.com/MilesDowe/feel/feel.db"
	}
	database, err := sql.Open("sqlite3", databaseLoc)

	verifyTableExists(database)

	if err != nil {
		log.Fatal(err)
	}

	return database
}

// runs a query to create the standard feel_recording table if it doesn't already exist
func verifyTableExists(db *sql.DB) {
	stmt, err := db.Prepare(createTable)
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}

	stmt.Exec()
}

// DeleteRecord : deletes the record with the given ID from the feel_recording table
func DeleteRecord(id int) {
	db := OpenDb()
	defer db.Close()
	db.Exec(deleteRecordPerID, id)
}
