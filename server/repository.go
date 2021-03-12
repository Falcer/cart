package server

import (
	"errors"

	"github.com/dgraph-io/badger/v3"
	"github.com/segmentio/ksuid"
)

// Repository interface
type Repository interface {
	Login(login *Login) (*User, error)
	Register(register *Register) (*User, error)
	GetUsers() (*[]User, error)
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
func NewRepository(db *badger.DB) Repository {
	return &repo{
		badgerDB: db,
	}
}
func (r *repo) Login(login *Login) (*User, error) {
	var user *User
	err := r.badgerDB.View(func(txn *badger.Txn) error {
		iopt := badger.DefaultIteratorOptions
		itr := txn.NewIterator(iopt)
		defer itr.Close()
		for itr.Rewind(); itr.Valid(); itr.Next() {
			err := itr.Item().Value(func(val []byte) error {
				tmp := Decode(val)
				if tmp.Username == login.Username {
					user = tmp
					return nil
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("User not found")
	}
	return user, nil
}
func (r *repo) Register(register *Register) (*User, error) {
	// Check if Exist
	err := r.badgerDB.View(func(txn *badger.Txn) error {
		iopt := badger.DefaultIteratorOptions
		itr := txn.NewIterator(iopt)
		defer itr.Close()
		for itr.Rewind(); itr.Valid(); itr.Next() {
			err := itr.Item().Value(func(val []byte) error {
				tmp := Decode(val)
				if tmp.Username == register.Username {
					return errors.New("Username Exist")
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	user := User{
		ID:       ksuid.New().String(),
		Username: register.Username,
		Fullname: register.Fullname,
	}
	// Set New List into key `users`
	err = r.badgerDB.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(user.ID), []byte(user.Encode()))
		return err
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *repo) GetUsers() (*[]User, error) {
	var users []User
	err := r.badgerDB.View(func(txn *badger.Txn) error {
		iopt := badger.DefaultIteratorOptions
		itr := txn.NewIterator(iopt)
		defer itr.Close()
		for itr.Rewind(); itr.Valid(); itr.Next() {
			err := itr.Item().Value(func(val []byte) error {
				users = append(users, *Decode(val))
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &users, nil
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
