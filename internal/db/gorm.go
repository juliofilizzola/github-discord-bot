package db

import (
	"fmt"
	"github-discord-bot/internal/config"
	"github-discord-bot/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func connection(dsn string) error {
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	return nil
}

func startMigrations() error {
	err := DB.AutoMigrate(model.GetRegisteredModels()...)
	if err != nil {
		return fmt.Errorf("erro ao migrar banco de dados: %v", err)
	}
	return nil
}

func InitializeDatabase() error {
	cfg := config.Load()
	err := connection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return err
	}
	err = startMigrations()
	if err != nil {
		log.Fatalf("Failed to start migrations: %v", err)
		return err
	}

	if DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	return nil
}
