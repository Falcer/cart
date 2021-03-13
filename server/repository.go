package server

import (
	"context"
	"errors"
	"time"

	"github.com/dgraph-io/badger/v3"
	"github.com/go-redis/redis/v8"
	"github.com/segmentio/ksuid"
)

var (
	products = []Product{
		{
			ID:       ksuid.New().String(),
			Name:     "Asus ROG",
			ImageURL: "https://cdn.medcom.id/dynamic/content/2019/07/03/1038320/fO7HtXghpT.jpg?w=480",
			Price:    20_000_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Macbook Pro m1",
			ImageURL: "https://www.apple.com/v/macbook-pro-13/g/images/overview/hero_endframe__bsza6x4fldiq_large_2x.jpg",
			Price:    40_000_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Alienware",
			ImageURL: "https://images-na.ssl-images-amazon.com/images/I/71hhY4ikVwL._AC_SL1500_.jpg",
			Price:    50_000_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Asus Vivobook",
			ImageURL: "https://images-na.ssl-images-amazon.com/images/I/81oVSSITRQL._AC_SL1500_.jpg",
			Price:    6_000_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Lenovo Legion",
			ImageURL: "https://pisces.bbystatic.com/image2/BestBuy_US/images/products/6398/6398969_sd.jpg",
			Price:    17_000_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Xiaomi Mi 10 Ultra",
			ImageURL: "https://www.gizmochina.com/wp-content/uploads/2020/08/a-500x500.jpg",
			Price:    9_000_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Samsung A51",
			ImageURL: "https://images.samsung.com/is/image/samsung/id-galaxy-a51-a515-sm-a515fzbwxid-209736060?$720_576_PNG$",
			Price:    5_000_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Iphone 12",
			ImageURL: "https://cdn.tmobile.com/content/dam/t-mobile/en-p/cell-phones/apple/Apple-iPhone-12/Blue/Apple-iPhone-12-Blue-backimage.png",
			Price:    17_000_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Stealseries Arctis Pro DTS",
			ImageURL: "https://lh3.googleusercontent.com/proxy/FTdEwgA1FBIRBEI2b3i6_G7lHE9d_RCHbDJCAUZCvbm07P2LR6hgASdnXketnuBkKRZi1Gb-vIpslf6y3f2VW5T43khylBzUOj2dhYHQK3A7mhTnR-LCmttB6xDyE_VwMXMm1I2HOkQLS-9FOUzpRV6oqP0PiWMTcrgJIilD0Ts2sSzop1UnkF6BPO41RLmWCe6OSwsQ19skQCdzVfMb9sJbRzohdfpnBCH9JU-WAWs8hkPrItQ",
			Price:    2_710_000,
		},
		{
			ID:       ksuid.New().String(),
			Name:     "Razer Blackwidow - Green Switch",
			ImageURL: "https://assets2.razerzone.com/images/blackwidow-2019/BlackWidow2019_OGimage-1200x630.jpg",
			Price:    1_730_000,
		},
	}
)

// Repository interface
type Repository interface {
	Login(login *Login) (*User, error)
	Register(register *Register) (*User, error)
	GetUsers() (*[]User, error)
	GetUserByID(userID string) (*User, error)
	GetProducts() (*[]Product, error)
	GetProductByID(id string) (*Product, error)
	GetCart(userID string) (*Cart, error)
	AddCart(userID string, productID string) error
	ChangeAmountCart(userID, cartID, productID string, amount uint8) error
	PaidCart(userID, cartID string) error
}

type repo struct {
	badgerDB *badger.DB
	redis    *redis.Client
}

// NewRepository instance
func NewRepository(badgerDB *badger.DB, redisClient *redis.Client) Repository {
	return &repo{
		badgerDB: badgerDB,
		redis:    redisClient,
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
				tmp := DecodeUser(val)
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
				tmp := DecodeUser(val)
				if tmp.Username == register.Username {
					return errors.New("User already exist")
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
	// Set New user with key is UserID
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
				users = append(users, *DecodeUser(val))
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
func (r *repo) GetUserByID(userID string) (*User, error) {
	var user *User
	err := r.badgerDB.View(func(txn *badger.Txn) error {
		iopt := badger.DefaultIteratorOptions
		itr := txn.NewIterator(iopt)
		defer itr.Close()
		for itr.Rewind(); itr.Valid(); itr.Next() {
			err := itr.Item().Value(func(val []byte) error {
				tmp := DecodeUser(val)
				if tmp.ID == userID {
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
func (r *repo) GetProducts() (*[]Product, error) {
	return &products, nil
}
func (r *repo) GetProductByID(id string) (*Product, error) {
	for _, v := range products {
		if id == v.ID {
			return &v, nil
		}
	}
	return nil, errors.New("Product not found")
}
func (r *repo) GetCart(userID string) (*Cart, error) {
	return nil, nil
}
func (r *repo) AddCart(userID string, productID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	product, err := r.GetProductByID(productID)
	if err != nil {
		return err
	}
	val, err := r.redis.Get(ctx, userID).Result()
	if err == redis.Nil {
		// Add New Key Value
		user, err := r.GetUserByID(userID)
		if err != nil {
			return err
		}
		items := []ItemCart{
			{
				ID:      ksuid.New().String(),
				Product: product,
				Amount:  1,
			},
		}
		cart := Cart{
			ID:     ksuid.New().String(),
			IsPaid: false,
			User:   user,
			Items:  &items,
		}
		err = r.redis.Set(ctx, userID, cart, 0).Err()
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	} else {
		// Add into cart
		cart, err := DecodeCart(val)
		if err != nil {
			return err
		}
		products := *(cart).Items
		products = append(products, ItemCart{
			ID:      ksuid.New().String(),
			Product: product,
			Amount:  1,
		})
		cart.Items = &products
		err = r.redis.Set(ctx, userID, *cart, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *repo) ChangeAmountCart(userID, cartID, productID string, amount uint8) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	val, err := r.redis.Get(ctx, userID).Result()
	if err == redis.Nil {
		return errors.New("Cart not found")
	} else if err != nil {
		return err
	} else {
		// Change Amount of Cart
		cart, err := DecodeCart(val)
		if err != nil {
			return err
		}
		// Find ItemCart with ProductID
		products := *(cart).Items
		find := false
		// Linear Search ItemCart
		for _, v := range products {
			// If ItemCart match with ProductID
			if v.Product.ID == productID {
				// Change Amount of ItemCart
				v.Amount = amount
				// And change `find` to true
			}
		}
		// If `find` is false ItemCart is not found
		if !find {
			return errors.New("Product not found")
		}
		cart.Items = &products
		err = r.redis.Set(ctx, userID, *cart, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
func (r *repo) PaidCart(userID, cartID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	val, err := r.redis.Get(ctx, userID).Result()
	if err == redis.Nil {
		return errors.New("Cart not found")
	} else if err != nil {
		return err
	} else {
		cart, err := DecodeCart(val)
		if err != nil {
			return err
		}
		cart.IsPaid = true
		err = r.redis.Set(ctx, userID, *cart, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
