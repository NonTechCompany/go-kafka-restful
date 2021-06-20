package db

import (
	"fmt"
	"github.com/NonTechCompany/go-kafka-restful/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Connection *gorm.DB

func Init() {
	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.GetAppConfig().DB.Username, config.GetAppConfig().DB.Password, config.GetAppConfig().DB.Host, config.GetAppConfig().DB.Port, config.GetAppConfig().DB.DbName)
	database, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Connection = database
}
