package repository

import "database/sql"

type Repository struct {
	db *sql.DB
}

func (r *Repository) GetBalance(id int) (int, error) {
	var balance int
	err := r.db.QueryRow("SELECT balance FROM accounts WHERE id = $1", id).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}

func (r *Repository) UpdateBalance(id int, amount int) error {
	_, err := r.db.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, id)
	return err
}
