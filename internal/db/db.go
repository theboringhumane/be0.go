package db

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"be0/internal/config"
	"be0/internal/models"
	console "be0/internal/utils/logger"
)

var DB *gorm.DB
var log = console.New("DB")

func Connect(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)

	log.Info("Connecting to database...")
	maxRetries := 5
	var err error
	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Info),
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true,
			AllowGlobalUpdate:                        false,
		})
		if err == nil {
			log.Info("DSN: %s", dsn)
			log.Success("Connected to database")

			// Configure connection pool
			sqlDB, err := DB.DB()
			if err != nil {
				return log.Error("Failed to get underlying *sql.DB instance", err)
			}

			// Set connection pool settings
			sqlDB.SetMaxOpenConns(100)                 // Maximum number of open connections to the database
			sqlDB.SetMaxIdleConns(10)                  // Maximum number of idle connections in the pool
			sqlDB.SetConnMaxLifetime(time.Hour)        // Maximum amount of time a connection may be reused
			sqlDB.SetConnMaxIdleTime(time.Minute * 30) // Maximum amount of time a connection may be idle

			// Run migrations
			if err := runMigrations(); err != nil {
				return log.Error("Failed to run migrations", err)
			}

			log.Success("Migrations completed")

			return nil
		}
		log.Warn("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Second * 5)
	}
	return log.Error("failed to connect to database after %d attempts", fmt.Errorf("failed to connect to database after %d attempts", maxRetries))
}

func runMigrations() error {
	log.Info("Running migrations...")
	// Begin transaction for migrations
	tx := DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Defer rollback in case of error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.AutoMigrate(
		// Base models without foreign keys
		&models.User{},
		&models.Team{},
		&models.Resource{},

		// Models with single foreign key dependencies
		&models.PasswordReset{},
		&models.TeamInvite{},
		&models.AuthTransaction{},
		// Permission models
		&models.UserPermission{},
		&models.ResourcePermission{},
	); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func GetDB() *gorm.DB {
	return DB
}
