package main

import (
	"log"

	"github.com/Falcer/cart/server"
	"github.com/dgraph-io/badger/v3"
	"github.com/gofiber/fiber/v2"
)

var (
	db      *badger.DB
	repo    server.Repository
	service server.Service
)

func init() {
	db, err := badger.Open(badger.DefaultOptions("./database"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo = server.NewRepository(db)
	service = server.NewService(repo)
}

func main() {
	app := fiber.New()

	// Authentication
	app.Post("/login", login)
	app.Post("/register", register)

	api := app.Group("/api/v1")
	api.Get("/users", getAllUser)
	api.Get("/carts", getAllCart)
	api.Get("/cart", getUserCart)
}

// HTTP handler
func login(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Login",
	})
}

func register(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Register",
	})
}

func getAllUser(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get all user",
	})
}

func getAllCart(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get all cart",
	})
}

func getUserCart(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get User Cart",
	})
}
