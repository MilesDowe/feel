package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const createTable = "CREATE TABLE IF NOT EXISTS feel_recording (" +
	"id INTEGER PRIMARY KEY," +
	"score INTEGER," +
	"concern TEXT NULLABLE," +
	"grateful TEXT NULLABLE," +
	"learn TEXT NULLABLE," +
	"entered INTEGER)"

const addRecord = "INSERT INTO feel_recording (score, concern, grateful, learn, entered) VALUES (?, ?, ?, ?, ?)"

// Min : Lowest a happy score can be
const Min = 1

// Max : Highest a happy score can be
const Max = 10

// MinStr : String representation of Min
var MinStr = strconv.Itoa(Min)

// MaxStr : String representation of Max
var MaxStr = strconv.Itoa(Max)

func readUserInput() (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Printf("How happy do you feel right now? Choose from %s (awful) to %s (great):", MaxStr, MinStr)
	score, _ := reader.ReadString('\n')

	fmt.Println("Anything have you concerned? (<enter> to skip)")
	concern, _ := reader.ReadString('\n')

	fmt.Println("Do you feel grateful for anything? (<enter> to skip)")
	grateful, _ := reader.ReadString('\n')

	fmt.Println("Did you learn anything new today? (<enter> to skip)")
	learn, _ := reader.ReadString('\n')
	fmt.Println()

	return score, concern, grateful, learn
}

func recordToDatabase(databaseLoc, score, concern, grateful, learn string) {
	database := getDatabaseAndPrepareTable()

	score = checkScoreInput(score)

	statement, _ := database.Prepare(addRecord)
	statement.Exec(score, concern, grateful, learn, time.Now().Unix())
}

func getDatabaseAndPrepareTable() *sql.DB {
	database, _ := sql.Open("sqlite3", databaseLoc)
	statement, _ := database.Prepare(createTable)
	statement.Exec()

	return database
}

func checkScoreInput(score string) string {
	score = strings.TrimSpace(score)
	i, err := strconv.Atoi(score)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if i < Min {
		score = MinStr
	} else if i > Max {
		score = MaxStr
	}
	return score
}
