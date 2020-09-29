package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/models"
)

// GetRecords - gets all posts
func GetRecords(c *fiber.Ctx) error {
	// db connection
	InitDatabase()
	// records type
	type Records struct {
		ID      uint
		Website string
		WPass   string
	}
	var res []Records
	// manager model
	var records []models.Manager
	// token payload
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenID := claims["id"].(float64)
	var IDtoInt int = int(tokenID)
	// query
	data := database.Model(&records).Where("user_id = ?", IDtoInt).Find(&res)
	if data == nil {
		return c.SendStatus(400)
	}
	return c.JSON(&res)
}
