package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var connString = "postgresql://postgres:example@localhost:5432/postgres"

var Connection *gorm.DB

func Init() {
	database, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Connection = database
}
