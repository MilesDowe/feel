package main

import (
    "bufio"
    "database/sql"
    "fmt"
    "os"
    "time"
    _ "github.com/mattn/go-sqlite3"
)

const createTable = "CREATE TABLE IF NOT EXISTS feel_recording (" +
    "id INTEGER PRIMARY KEY," +
    "score INTEGER," +
    "concern TEXT NULLABLE," +
    "grateful TEXT NULLABLE," +
    "learn TEXT NULLABLE," +
    "entered INTEGER)"
const addRecord = "INSERT INTO feel_recording " +
    "(score, concern, grateful, learn, entered) VALUES (?, ?, ?, ?, ?)"

func readUserInput() (string, string, string, string) {
    reader := bufio.NewReader(os.Stdin)

    fmt.Println()
    fmt.Println("How happy do you feel right now? Choose from 1 (awful) to 10 (elated):")
    score, _ := reader.ReadString('\n')

    fmt.Println("Anything have you concerned? (`enter` to skip)")
    concern, _ := reader.ReadString('\n')

    fmt.Println("Do you feel grateful for anything? (`enter` to skip)")
    grateful, _ := reader.ReadString('\n')

    fmt.Println("Did you learn anything new today? (`enter` to skip)")
    learn, _ := reader.ReadString('\n')
    fmt.Println()

    return score, concern, grateful, learn
}

func recordToDatabase(databaseLoc, score, concern, grateful, learn string) {
    database, _ := sql.Open("sqlite3", databaseLoc)
    statement, _ := database.Prepare(createTable)
    statement.Exec()

    statement, _ = database.Prepare(addRecord)
    statement.Exec(score, concern, grateful, learn, time.Now().Unix())
}
