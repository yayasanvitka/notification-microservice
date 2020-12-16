package main

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"log"
	"os"
	"whatsapp-microservice/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"whatsapp-microservice/pkg/config"
)

func main() {
	godotenv.Load()
	cfg := config.Get()
	dir, _ := os.Getwd()

	log.Printf("Working Dir: %s", dir)

	direction := cfg.GetMigration()

	if direction != "down" && direction != "up" {
		log.Fatal("-migrate accepts [up, down] values only")
	}

	//log.Println(cfg.GetDBConnStr())

	// open connection
	log.Println("Opening Connection...")
	db, _ := sql.Open("mysql", cfg.GetDBConnStr())

	// prepare driver
	log.Println("Preparing Driver...")
    driver, _ := mysql.WithInstance(db, &mysql.Config{})

    // prepare files to migrate
    m, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s/db/migrations", dir),
        "mysql",
        driver,
    )

	if err != nil {
		log.Fatal(err)
	}

	// init logger
	m.Log = &logger.MigrationLogger{}

	if direction == "up" {
		m.Log.Printf("Applying changes...")
		err = m.Up()
	}

	if direction == "down" {
		m.Log.Printf("Rolling Back changes...")
		err = m.Down()
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}

	// log if error detected
	if err != nil {
		m.Log.Printf(err.Error())
	}

	// checking database version
	version, _, err := m.Version()

	if err != nil {
		m.Log.Printf(err.Error())
	}

	m.Log.Printf("Active DB Version: %d", version)
}
