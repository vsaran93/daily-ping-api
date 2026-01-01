
package storage
import (
	"log"
	"os"
	"gorm.io/gorm"
)


func OpenDbConnection() *gorm.DB {
	config := &Config{
		Host:	os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DBName: os.Getenv("DB_NAME"),
	}

	db, err := NewConnection(config)

	if(err != nil) {
		log.Fatal("There is an error with database")
	}
	return db
}
