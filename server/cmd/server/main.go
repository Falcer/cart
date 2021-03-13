package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Falcer/cart/server"
	"github.com/dgraph-io/badger/v3"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type app struct {
	service server.Service
}

func init() {

}

func main() {
	badgerDB, err := badger.Open(badger.DefaultOptions("./database"))
	if err != nil {
		log.Fatal(err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fiberApp := fiber.New()
	repo := server.NewRepository(badgerDB, redisClient)
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
	api.Get("/products", app.getAllProduct)
	api.Get("/products/:id", app.getProductByID)
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

	<-c
	fmt.Println("\nGracefully shutting down...")
	defer badgerDB.Close()
	defer redisClient.Close()
	_ = fiberApp.Shutdown()

	fmt.Println("Running cleanup tasks...")
}

// HTTP handler
func (p *app) login(c *fiber.Ctx) error {
	login := new(server.Login)
	if err := c.BodyParser(login); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	res, err := p.service.Login(login)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
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
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	res, err := p.service.Register(register)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
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
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get all user",
		"data":    *res,
	})
}

func (p *app) getAllProduct(c *fiber.Ctx) error {
	res, err := p.service.GetProducts()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get all product",
		"data":    *res,
	})
}

func (p *app) getProductByID(c *fiber.Ctx) error {
	res, err := p.service.GetProductByID(c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": fmt.Sprintf("Get product id : %s", c.Params("id")),
		"data":    *res,
	})
}

func (p *app) getAllCart(c *fiber.Ctx) error {
	result, err := p.service.GetCarts()
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	data := *result
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get all cart",
		"data":    data,
	})
}

func (p *app) getUserCart(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Get User Cart",
	})
}

func (p *app) addProductToCart(c *fiber.Ctx) error {
	body := new(server.AddCart)
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	err := p.service.AddCart(body.UserID, body.ProductID)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Add product to cart successfully",
	})
}

func (p *app) updateCart(c *fiber.Ctx) error {
	body := new(server.ChangeAmountCart)
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	err := p.service.ChangeAmountCart(body.UserID, body.ProductID, body.Amount)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Update product to cart successfully",
	})
}

func (p *app) paidCart(c *fiber.Ctx) error {
	body := new(server.PaidCart)
	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Something went wrong!",
		})
	}
	err := p.service.PaidCart(body.UserID)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(200).JSON(&fiber.Map{
		"message": "Cart paid",
	})
}
