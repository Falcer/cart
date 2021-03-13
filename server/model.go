package server

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type (
	// Login model
	Login struct {
		Username string `json:"username"`
	}
	// Register model
	Register struct {
		Username string `json:"username"`
		Fullname string `json:"fullname"`
	}
	// User model
	User struct {
		ID       string `json:"id"`
		Username string `json:"username"`
		Fullname string `json:"fullname"`
	}

	// Product model
	Product struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Price    uint64 `json:"price"`
		ImageURL string `json:"image_url"`
	}

	// ItemCart model
	ItemCart struct {
		ID      string   `json:"id"`
		Product *Product `json:"product"`
		Amount  uint8    `json:"amount"`
	}

	// Cart model
	Cart struct {
		ID     string      `json:"id"`
		User   *User       `json:"user"`
		Items  *[]ItemCart `json:"items"`
		IsPaid bool        `json:"id_paid"`
	}
)

// Encode Cart
func (c *Cart) Encode() (*string, error) {
	res, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	result := string(res)
	return &result, nil
}

// Decode Cart
func DecodeCart(data string) (*Cart, error) {
	var cart Cart
	if err := json.Unmarshal([]byte(data), &cart); err != nil {
		return nil, err
	}
	return &cart, nil
}

// Encode user
func (u *User) Encode() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(u)
	if err != nil {
		return nil
	}
	return res.Bytes()
}

// DecodeUser user
func DecodeUser(data []byte) *User {
	var res User
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&res)
	if err != nil {
		return nil
	}
	return &res
}
