package database

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		"localhost", "postgres", "password", "e-pemilu", "5432", "prefer",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err = connection.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Info("Connected and pinged database successfully")

	connection.SetMaxIdleConns(30)
	connection.SetMaxOpenConns(100)
	connection.SetConnMaxLifetime(time.Second * time.Duration(300))

	// db.Debug().AutoMigrate(
	// 	domain.Role{},
	// 	domain.User{},
	// 	domain.Province{},
	// 	domain.City{},
	// 	domain.District{},
	// 	domain.Merchant{},
	// 	domain.MerchantTable{},
	// 	domain.MerchantFloor{},
	// 	domain.MerchantLocation{},
	//  domain.Village{},
	// )

	DB = db
	return db
}
