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
		Website string
		WPass   string
	}
	var recRes RecResponse
	// gets the manger model
	record := new(models.Manager)
	c.BodyParser(record)
	// record id
	recordID := c.Params("id")
	// Update
	updatePassword := database.Model(&models.Manager{}).Where("id = ?", recordID).Update("w_pass", record.WPass)
	if updatePassword == nil {
		return c.SendStatus(400)
	}
	updateWebsite := database.Model(&models.Manager{}).Where("id = ?", recordID).Update("website", record.Website)
	if updateWebsite == nil {
		return c.SendStatus(400)
	}
	// Response
	recRes = RecResponse{
		Website: record.Website,
		WPass:   record.WPass,
	}
	return c.JSON(recRes)
}
