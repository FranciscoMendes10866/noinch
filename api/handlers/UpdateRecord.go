package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/models"
)

// UpdateRecord - updates a single record
func UpdateRecord(c *fiber.Ctx) error {
	// db connection
	InitDatabase()
	// record type
	type RecResponse struct {
		WPass string
	}
	var recRes RecResponse
	// gets the manger model
	record := new(models.Manager)
	c.BodyParser(record)
	// record id
	recordID := c.Params("id")
	// Update
	update := database.Model(&models.Manager{}).Where("id = ?", recordID).Update("w_pass", record.WPass)
	if update == nil {
		return c.SendStatus(400)
	}
	// Response
	recRes = RecResponse{
		WPass: record.WPass,
	}
	return c.JSON(recRes)
}
