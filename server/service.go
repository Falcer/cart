package server

// Service interface
type Service interface {
	Login(login *Login) (*User, error)
	Register(register *Register) (*User, error)
	GetUsers() (*[]User, error)
	GetProducts() (*[]Product, error)
	GetProductByID(id string) (*Product, error)
	GetCarts() (*[]Cart, error)
	GetCart(userID string) (*Cart, error)
	AddCart(userID string, productID string) error
	ChangeAmountCart(userID, productID string, amount uint8) error
	PaidCart(userID string) error
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
func (s *service) GetUsers() (*[]User, error) {
	return s.repo.GetUsers()
}
func (s *service) GetProducts() (*[]Product, error) {
	return s.repo.GetProducts()
}
func (s *service) GetProductByID(id string) (*Product, error) {
	return s.repo.GetProductByID(id)
}
func (s *service) GetCarts() (*[]Cart, error) {
	return s.repo.GetCarts()
}
func (s *service) GetCart(userID string) (*Cart, error) {
	return s.repo.GetCart(userID)
}
func (s *service) AddCart(userID string, productID string) error {
	return s.repo.AddCart(userID, productID)
}
func (s *service) ChangeAmountCart(userID, productID string, amount uint8) error {
	return s.repo.ChangeAmountCart(userID, productID, amount)
}
func (s *service) PaidCart(userID string) error {
	return s.repo.PaidCart(userID)
}
