package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/models"
)

// DeleteRecord - deletes a given record
func DeleteRecord(c *fiber.Ctx) error {
	// db connection
	InitDatabase()
	// gets the record id
	recordID := c.Params("id")
	// gets the manager model
	var manager models.Manager
	// deletes the record
	delete := database.Where("id = ?", recordID).Delete(&manager)
	if delete == nil {
		return c.SendStatus(400)
	}
	// response
	return c.SendStatus(200)
}
