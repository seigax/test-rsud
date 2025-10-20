package config

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib"
	"gitlab.com/erloom.id/libraries/go/backend-skeleton/lib/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPG() (*lib.Database, error) {
	username := viper.GetString("POSTGRES_USERNAME")
	password := viper.GetString("POSTGRES_PASSWORD")
	host := viper.GetString("POSTGRES_HOST")
	port := viper.GetString("POSTGRES_PORT")
	dbname := viper.GetString("POSTGRES_DATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.GetLogger(),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Error(context.Background(), fmt.Sprint("failed to get sql.DB from gorm.DB: %v", err), nil)
	}
	// Set connection pool parameters
	sqlDB.SetMaxIdleConns(viper.GetInt("DB_POOL_MAX_IDLE_CON"))
	sqlDB.SetMaxOpenConns(viper.GetInt("DB_POOL_MAX_OPEN_CON"))
	sqlDB.SetConnMaxLifetime(time.Duration(viper.GetInt("DB_POOL_CON_LIFETIME_IN_MINUTE")) * time.Minute)

	// go func() {
	// 	for {
	// 		stats := sqlDB.Stats()
	// 		log.Printf("Open Connections: %d\n", stats.OpenConnections)
	// 		log.Printf("Idle Connections: %d\n", stats.Idle)
	// 		log.Printf("In Use Connections: %d\n", stats.InUse)
	// 		log.Printf("Wait Count: %d\n", stats.WaitCount)
	// 		log.Printf("Wait Duration: %s\n", stats.WaitDuration)
	// 		time.Sleep(5 * time.Second)
	// 	}
	// }()

	logger.Info(context.Background(), "successfully connected to postgres", make(map[string]interface{}))
	return &lib.Database{DB: db}, nil
}
