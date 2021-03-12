package server

// Service interface
type Service interface {
	Login(login *Login) (*User, error)
	Register(register *Register) (*User, error)
	GetProducts() (*[]Product, error)
	GetProductByID(id string) (*Product, error)
	GetCart(userID string) (*Cart, error)
	AddCart(userID string, productID string) error
	ChangeAmountCart(userID, cartID, productID string, amount uint8) error
	PaidCart(userID, cartID string) error
}

type service struct {
	repo Repository
}

// NewService instance
func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Login(login *Login) (*User, error) {
	return s.repo.Login(login)
}
func (s *service) Register(register *Register) (*User, error) {
	return s.repo.Register(register)
}
func (s *service) GetProducts() (*[]Product, error) {
	return s.repo.GetProducts()
}
func (s *service) GetProductByID(id string) (*Product, error) {
	return s.repo.GetProductByID(id)
}
func (s *service) GetCart(userID string) (*Cart, error) {
	return s.repo.GetCart(userID)
}
func (s *service) AddCart(userID string, productID string) error {
	return s.repo.AddCart(userID, productID)
}
func (s *service) ChangeAmountCart(userID, cartID, productID string, amount uint8) error {
	return s.repo.ChangeAmountCart(userID, cartID, productID, amount)
}
func (s *service) PaidCart(userID, cartID string) error {
	return s.repo.PaidCart(userID, cartID)
}
