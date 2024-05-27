package db

import (
	"fmt"
	"ka/config"
	"ka/pkg/logging"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbOnce sync.Once

// InitDB initializes the database connection
func InitDB() (*gorm.DB, error) {
	var err error
	dbOnce.Do(func() {
		envVars := config.GetEnvVars()
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", envVars.GetPostgresHost(), envVars.GetPostgresUser(), envVars.GetPostgresPassword(), envVars.GetPostgresDB(), envVars.GetPostgresPort())
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: true, Logger: NewGormLogger(logging.InitialiseLogger())})

		if err != nil {
			return
		}

		//set max conn
		sqlDB, err := db.DB()
		if err != nil {
			return
		}

		sqlDB.SetMaxIdleConns(3)
		sqlDB.SetMaxOpenConns(10)
	})
	return db, err
}
