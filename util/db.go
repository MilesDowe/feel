package util

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // database driver
)

// TODO: get GOPATH lookup to work:
// var feelLoc string = os.Getenv("GOPATH") + "/src/github.com/milesdowe/feel"
const databaseLoc = "C:/Users/Miles/go/src/github.com/milesdowe/feel/feel.db"

const createTable = "CREATE TABLE IF NOT EXISTS feel_recording (" +
	"id INTEGER PRIMARY KEY," +
	"score INTEGER," +
	"concern TEXT NULLABLE," +
	"grateful TEXT NULLABLE," +
	"learn TEXT NULLABLE," +
	"entered INTEGER)"

const deleteRecordPerID = "DELETE FROM feel_recording WHERE id = ?"

// OpenDb : returns a connection to the SQLite database
func OpenDb() *sql.DB {
	database, err := sql.Open("sqlite3", databaseLoc)

	verifyTableExists(database)

	if err != nil {
		fmt.Println(err)
	}

	return database
}

// runs a query to create the standard feel_recording table if it doesn't already exist
func verifyTableExists(db *sql.DB) {
	stmt, _ := db.Prepare(createTable)
	defer stmt.Close()
	stmt.Exec()
}

// DeleteRecord : deletes the record with the given ID from the feel_recording table
func DeleteRecord(id int) {
	db := OpenDb()
	defer db.Close()
	db.Exec(deleteRecordPerID, id)
}
