package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Falcer/cart/server"
	"github.com/dgraph-io/badger/v3"
	"github.com/gofiber/fiber/v2"
)

type app struct {
	service server.Service
}

func main() {
	db, err := badger.Open(badger.DefaultOptions("./database"))
	if err != nil {
		log.Fatal(err)
	}
	fiberApp := fiber.New()
	repo := server.NewRepository(db)
	service := server.NewService(repo)

	app := &app{service: service}

	fiberApp.Get("", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(&fiber.Map{
			"message": "App is running",
		})
	})

	// Authentication
	fiberApp.Post("/login", app.login)
	fiberApp.Post("/register", app.register)

	api := fiberApp.Group("/api/v1")
	api.Get("/users", app.getAllUser)
	api.Get("/carts", app.getAllCart)
	api.Get("/cart", app.getUserCart)
	api.Post("/cart", app.addProductToCart)
	api.Put("/cart", app.updateCart)
	api.Post("/cart/paid", app.paidCart)

	fmt.Println("Server running at http://127.0.0.1:8080")
	go func() {
		if err := fiberApp.Listen(":8080"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	fmt.Println("Gracefully shutting down...")
	defer db.Close()
	_ = fiberApp.Shutdown()

	fmt.Println("Running cleanup tasks...")
}

// HTTP handler
func (p *app) login(c *fiber.Ctx) error {
	login := new(server.Login)
	if err := c.BodyParser(login); err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	res, err := p.service.Login(login)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Login successfully",
		"data":    *res,
	})
}

func (p *app) register(c *fiber.Ctx) error {
	register := new(server.Register)
	if err := c.BodyParser(register); err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	res, err := p.service.Register(register)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Register successfully",
		"data":    *res,
	})
}

func (p *app) getAllUser(c *fiber.Ctx) error {
	res, err := p.service.GetUsers()
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get all user",
		"data":    *res,
	})
}

func (p *app) getAllCart(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get all cart",
	})
}

func (p *app) getUserCart(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get User Cart",
	})
}

func (p *app) addProductToCart(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Add product to cart",
	})
}

func (p *app) updateCart(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Update product to cart",
	})
}

func (p *app) paidCart(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Cart paid",
	})
}
