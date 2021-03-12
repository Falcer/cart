package server

import "github.com/dgraph-io/badger/v3"

// Repository interface
type Repository interface {
	Login(login *Login) (*User, error)
	Register(register *Register) (*User, error)
	GetProducts() (*[]Product, error)
	GetProductByID(id string) (*Product, error)
	GetCart(userID string) (*Cart, error)
	AddCart(userID string, productID string) error
	ChangeAmountCart(userID, cartID, productID string, amount uint8) error
	PaidCart(userID, cartID string) error
}

type repo struct {
	badgerDB *badger.DB
}

// NewRepository instance
func NewRepository(badgerDB *badger.DB) Repository {
	return &repo{
		badgerDB: badgerDB,
	}
}

func (r *repo) Login(login *Login) (*User, error) {
	return nil, nil
}
func (r *repo) Register(register *Register) (*User, error) {
	return nil, nil
}
func (r *repo) GetProducts() (*[]Product, error) {
	return nil, nil
}
func (r *repo) GetProductByID(id string) (*Product, error) {
	return nil, nil
}
func (r *repo) GetCart(userID string) (*Cart, error) {
	return nil, nil
}
func (r *repo) AddCart(userID string, productID string) error {
	return nil
}
func (r *repo) ChangeAmountCart(userID, cartID, productID string, amount uint8) error {
	return nil
}
func (r *repo) PaidCart(userID, cartID string) error {
	return nil
}
