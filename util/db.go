package util

import (
    _ "github.com/mattn/go-sqlite3" // database driver
	"database/sql"
	"fmt"
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

// OpenDb : returns a connection to the SQLite database
func OpenDb() *sql.DB {
	database, err := sql.Open("sqlite3", databaseLoc)

	if err != nil {
		fmt.Println(err)
	}

	return database
}

// VerifyDbExists : runs a query to create the standard table if it doesn't already exist
func VerifyDbExists() {
	db := OpenDb()
	stmt, _ := db.Prepare(createTable)
	stmt.Exec()
}
