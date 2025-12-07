package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/swagger/v2"
	_ "github.com/n4mchun/swagger-in-go/docs"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = map[string]User{}

// @title			Fiber Example API
// @version		1.0
// @description	This is a sample swagger for Fiber
// @termsOfService	http://swagger.io/terms/
// @contact.name	n4mchun
// @contact.email	n4mchun@gmail.com
// @license.name	MIT
// @license.url	None
// @host			localhost:8080
// @BasePath		/
func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(":8080")
}

func setupRoutes(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/users", createUser)
	app.Get("/users", getAllUsers)
	app.Get("/users/:id", getUser)
	app.Put("/users/:id", updateUser)
	app.Delete("/users/:id", deleteUser)
}

func createUser(c fiber.Ctx) error {
	u := new(User)
	if err := c.Bind().Body(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	if _, exists := users[u.ID]; exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "user already exists"})
	}

	users[u.ID] = *u
	return c.JSON(u)
}

func getAllUsers(c fiber.Ctx) error {
	result := []User{}
	for _, u := range users {
		result = append(result, u)
	}
	return c.JSON(result)
}

func getUser(c fiber.Ctx) error {
	id := c.Params("id")
	u, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(u)
}

func updateUser(c fiber.Ctx) error {
	id := c.Params("id")
	_, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}

	u := new(User)
	if err := c.Bind().Body(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	u.ID = id

	users[id] = *u
	return c.JSON(u)
}

func deleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	_, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}

	delete(users, id)
	return c.JSON(fiber.Map{"status": "deleted"})
}
