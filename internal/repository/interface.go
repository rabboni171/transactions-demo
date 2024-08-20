package repository

type AccountRepository interface {
	GetBalance(int) (int, error)
	UpdateBalance(int, int) error
}
