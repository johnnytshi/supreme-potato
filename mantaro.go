// pulls the latest Specs repo and parse the podspec.json files
package main

import (
	"fmt"
	"os"

	"github.com/supreme-potato/db"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	fmt.Println("It's better to be a pirate than join the navy.")
	(&cli.App{}).Run(os.Args)

	db.DB().ListTabls()

	// create the DB table if needed

	// list folders

	// parse json

	// dump into db

	// repeat
}
