package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rabboni171/transactions-demo/configs"
	"github.com/sirupsen/logrus"
)

func NewDatabase(cfg *configs.Config, logger *logrus.Logger) (*sqlx.DB, error) {
	pgParams := cfg.DBParams

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		pgParams.Server, pgParams.User, pgParams.Password, pgParams.DataBase, pgParams.Port)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Fatal("couldn't connect to database", err)
		return nil, err
	}
	return db, nil
}
