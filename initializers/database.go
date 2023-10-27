package initializers

import (
	"fmt"
	"log"

	"github.com/kompit-recruitment/backend/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDatabase() *gorm.DB {
	host := viper.GetString("DB_HOST")
	user := viper.GetString("DB_USER")
	password := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	port := viper.GetString("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Error init database: ", err.Error())
	} else {
		log.Println("Connected successfully to the database")
	}

	doMigration()

	return db
}

func doMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Competition{})
}
