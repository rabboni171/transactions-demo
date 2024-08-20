package service

type AccountService interface {
	Transfer(int, int, int) error
}
