package base

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Database *gorm.DB
}

func NewPostgres(env Env, logger Logger) Postgres {
	host := env.PostgresHost
	username := env.PostgresUser
	password := env.PostgresPass
	databaseName := env.PostgresDB
	port := env.PostgresPort

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, databaseName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Zap.Error("Error creating postgres client: %s", err)
		panic(err)
	}

	logger.Zap.Info("Connected to postgres database")

	return Postgres{
		Database: db,
	}
}
