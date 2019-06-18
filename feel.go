package main

import (
	"flag"
	"fmt"
	"os"
)

// TODO: get GOPATH lookup to work:
// var feelLoc string = os.Getenv("GOPATH") + "/src/github.com/milesdowe/feel"
var databaseLoc = "C:/Users/Miles/go/src/github.com/milesdowe/feel/feel.db"

func main() {

	/* Subcommand: `now` */
	nowCommand := flag.NewFlagSet("now", flag.ExitOnError)
	nowHelpPtr := nowCommand.Bool("help", false, "Print help")
	//nowAmmendPtr := flag.Bool("ammend", false, "Change values already provided today.")
	//nowRemovePtr := flag.Bool("delete", false, "Delete today's entry")

	/* Subcommand: `log` */
	logCommand := flag.NewFlagSet("log", flag.ExitOnError)

	/* Subcommand: `stat` */

	/* Get the arguments */

	flag.Parse()

	/* Check for required subcommand */
	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	/** Check which subcommand was used **/

	switch os.Args[1] {
	case "now":
		nowCommand.Parse(os.Args[2:])
	case "log":
		logCommand.Parse(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	/** Run subcommand **/

	if nowCommand.Parsed() {
		if *nowHelpPtr {
			fmt.Fprintf(os.Stderr, "Usage of subcommand %s:\n", os.Args[1])
			nowCommand.PrintDefaults()
			os.Exit(0)
		}
		score, concern, grateful, learn := readUserInput()
		recordToDatabase(databaseLoc, score, concern, grateful, learn)
	}

	if logCommand.Parsed() {
		printLog(databaseLoc)
	}

	os.Exit(0)
}
