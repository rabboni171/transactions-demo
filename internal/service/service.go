package service

import (
	"database/sql"
	"errors"
	"github.com/rabboni171/transactions-demo/internal/repository"
)

type Service struct {
	repo repository.AccountRepository
	db   *sql.DB
}

func NewService(repo repository.AccountRepository, db *sql.DB) AccountService {
	return &Service{repo, db}
}

func (s *Service) Transfer(fromID, toID, amount int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	fromBalance, err := s.repo.GetBalance(fromID)
	if err != nil {
		return err
	}

	if fromBalance < amount {
		return errors.New("insufficient funds")
	}

	if err = s.repo.UpdateBalance(fromID, -amount); err != nil {
		return err
	}

	if err = s.repo.UpdateBalance(toID, amount); err != nil {
		return err
	}

	return nil
}
