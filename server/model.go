package server

type (
	// Login model
	Login struct {
		Username string `json:"username"`
		Fullname string `json:"fullname"`
	}
	// Register model
	Register struct {
		Username string `json:"username"`
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
