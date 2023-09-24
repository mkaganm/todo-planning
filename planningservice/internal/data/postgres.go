package data

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"planningservice/internal/config"
)

var PostgresDSN *string

// init is a function that initializes the data source name
func init() {
	dsn := createDSN()
	PostgresDSN = &dsn
}

// createDSN is a function that creates a data source name
func createDSN() string {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.EnvConfigs.PostgresHost,
		config.EnvConfigs.PostgresPort,
		config.EnvConfigs.PostgresUser,
		config.EnvConfigs.PostgresPass,
		config.EnvConfigs.PostgresDb,
		config.EnvConfigs.PostgresSSLMode)

	return dsn
}

// InitDB is a function that initializes the database
func InitDB() *gorm.DB {

	log.Default().Println("Connecting to database...")

	db, err := gorm.Open(postgres.Open(*PostgresDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to database: ", err)
	}

	log.Default().Println("Connected to database.")

	return db
}

// Close is a function that closes the database connection
func Close(db *gorm.DB) {

	log.Default().Println("Closing database connection...")

	dbSQL, err := db.DB()
	if err != nil {
		log.Fatal("Error while closing the database connection: ", err)
	}

	err = dbSQL.Close()
	if err != nil {
		log.Fatal("Error while closing the database connection: ", err)
	}

	log.Default().Println("Database connection closed.")
}
