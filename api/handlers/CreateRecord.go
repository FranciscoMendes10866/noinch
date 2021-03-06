package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/models"
)

// CreateRecord - creates a record in the Manager Table
func CreateRecord(c *fiber.Ctx) error {
	// db connection
	InitDatabase()
	// record type
	type RecResponse struct {
		Website string
		WPass   string
	}
	var recRes RecResponse
	// gets the manger model
	record := new(models.Manager)
	c.BodyParser(record)
	// token payload
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	tokenID := claims["id"].(float64)
	// converts the tokenID from float64 (1.0000) to int (1)
	var IDtoInt int = int(tokenID)
	// adds the user_id to the object
	record.UserID = IDtoInt
	// creates
	create := database.Create(record)
	if create == nil {
		return c.SendStatus(500)
	}
	recRes = RecResponse{
		Website: record.Website,
		WPass:   record.WPass,
	}
	return c.JSON(recRes)
}
