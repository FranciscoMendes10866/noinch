package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser - create a new user handler
func CreateUser(c *fiber.Ctx) error {
	// db connection
	InitDatabase()
	// gets the user input fields based on the model
	user := new(models.User)
	c.BodyParser(user)
	// hashes the given password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 4)
	if err != nil {
		return c.SendStatus(415)
	}
	// exchanges the given password with the hashed password
	user.Password = string(hash)
	// creates the user
	create := database.Create(user)
	if create == nil {
		return c.SendStatus(415)
	}
	// Response
	return c.SendStatus(fiber.StatusCreated)
}
