package models

type Account struct {
	ID      int `json:"id"`
	Balance int `json:"balance"` // для простоты берем целые суммы
}
