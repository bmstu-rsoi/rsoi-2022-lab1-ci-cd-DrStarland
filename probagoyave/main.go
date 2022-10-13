package main

import (
	"log"
	"os"

	"github.com/DrStarland/probagoyave/http/route"
	_ "github.com/DrStarland/probagoyave/http/validation"

	"goyave.dev/goyave/v4"
	// Import the appropriate GORM dialect for the database you're using.
	// _ "goyave.dev/goyave/v4/database/dialect/mysql"

	_ "goyave.dev/goyave/v4/database/dialect/postgres"
	// _ "goyave.dev/goyave/v4/database/dialect/sqlite"
	// _ "goyave.dev/goyave/v4/database/dialect/mssql"
)

func main() {
	// This is the entry point of your application.

	log.Println("Я запускаюсь.")

	if err := goyave.Start(route.Register); err != nil {
		os.Exit(err.(*goyave.Error).ExitCode)
	}
}
