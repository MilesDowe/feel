package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
	"time"
)

const getAllRecords = "SELECT id, score, concern, grateful, learn, entered FROM feel_recording"

// TODO: catch uninitialized database
func printLog(databaseLoc string) {
	database, _ := sql.Open("sqlite3", databaseLoc)
	rows, _ := database.Query(getAllRecords)

	var id int
	var score int
	var concern string
	var grateful string
	var learn string
	var entered int64

	fmt.Println()
	for rows.Next() {
		rows.Scan(&id, &score, &concern, &grateful, &learn, &entered)

		fmt.Println("Date: " + time.Unix(entered, 0).String())
		fmt.Println("Score: " + strconv.Itoa(score))
		fmt.Println("Concerned: " + format(concern))
		fmt.Println("Grateful: " + format(grateful))
		fmt.Println("Learned: " + format(learn))
		fmt.Println()
	}
}

func format(in string) string {
	result := strings.Trim(in, " \n")

	if result == "" {
		return "<skipped>"
	}
	return result
}
