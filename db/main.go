package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

const (
	COMMAND_UP   = "up"
	COMMAND_DOWN = "down"
)

func main() {
	if envErr := godotenv.Load(); envErr != nil {
		log.Fatal("error while loading environment file")
		os.Exit(1)
	}

	m, err := migrate.New("file://db/migrations", os.Getenv("DB_CONNECTION_STRING"))

	var command = COMMAND_UP

	if len(os.Args) > 1 {
		command = os.Args[1]
	}

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	if command == COMMAND_UP {
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
	} else if command == COMMAND_DOWN {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
	}
}
