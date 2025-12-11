package main

import (
	"fmt"

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

// @Summary		Create User
// @Description	create User by id, name, age
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			Body	body		User	true	"User Object"
// @Success		200		{object}	User
// @Failure		400		{object}	InvalidBodyError		"Invalid Body"
// @Failure		409		{object}	UserAlreadyExistsError	"User already exists"
// @Router			/users [post]
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

// @Summary		Get All Users
// @Description	get all users data
// @Tags			User
// @Accept			json
// @Produce		json
// @Success		200	{object}	[]User
// @Router			/users [get]
func getAllUsers(c fiber.Ctx) error {
	result := []User{}
	for _, u := range users {
		result = append(result, u)
	}
	return c.JSON(result)
}

// @Summary		Get User
// @Description	get user by ID
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200	{object}	User
// @Failure		404	{object}	NotFoundError	"User not found"
// @Router			/users/{id} [get]
func getUser(c fiber.Ctx) error {
	id := c.Params("id")

	fmt.Println(id)

	u, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(u)
}

// @Summary		Update User
// @Description	update user by ID
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id		path		string	true	"User ID"
// @Param			Body	body		User	true	"User Object"
// @Success		200		{object}	User
// @Failure		400		{object}	InvalidBodyError	"Invalid body"
// @Failure		404		{object}	NotFoundError		"User not found"
// @Router			/users/{id} [put]
func updateUser(c fiber.Ctx) error {
	id := c.Params("id")
	_, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"Error": "not found"})
	}

	u := new(User)
	if err := c.Bind().Body(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "invalid Body"})
	}
	u.ID = id

	users[id] = *u
	return c.JSON(u)
}

// @Summary		Delete User
// @Description	delete user by ID
// @Tags			User
// @Accept			json
// @Produce		json
// @Param			id	path		string			true	"User ID"
// @Success		200	{object}	DeleteSuccess	"User deleted"
// @Failure		404	{object}	NotFoundError	"User not found"
// @Router			/users/{id} [delete]
func deleteUser(c fiber.Ctx) error {
	id := c.Params("id")
	_, exists := users[id]
	if !exists {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "not found"})
	}

	delete(users, id)
	return c.JSON(fiber.Map{"status": "deleted"})
}
