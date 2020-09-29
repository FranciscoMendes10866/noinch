package handlers

import (
	"encoding/hex"
	"log"

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
	// encrypts the website password
	key := []byte("2tHbuAi7NnEttccVY5vIZ16sil5zbfhM") // 32 bytes
	plaintext := []byte(record.WPass)
	ciphertext, err := Encrypt(key, plaintext)
	if err != nil {
		log.Fatal(err)
	}
	// adds the the encrypted password
	record.WPass = hex.EncodeToString(ciphertext)
	// Update
	update := database.Model(&models.Manager{}).Where("id = ?", recordID).Update("w_pass", record.WPass)
	if update == nil {
		log.Fatal(err)
	}
	// Response
	recRes = RecResponse{
		WPass: record.WPass,
	}
	return c.JSON(recRes)
}
