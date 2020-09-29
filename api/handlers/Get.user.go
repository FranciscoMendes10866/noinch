package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/noinch-api/models"
	"golang.org/x/crypto/bcrypt"
)

// Get a single user
func Get(c *fiber.Ctx) error {
	// db connection
	InitDatabase()
	// database response object
	type DBResponse struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// var with the UserData struct
	var res DBResponse
	// parses the user according to the model
	user := new(models.User)
	c.BodyParser(user)
	// stores the user input password value
	pass := user.Password
	// makes a query to validate if the email exists
	data := database.Where(&models.User{Email: user.Email}).Find(user)
	// if the query is null, we will have an error
	// otherwise, we will store every diferent property value to the res
	if data == nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	res = DBResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
	// here we check if the passwords match
	match := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(pass))
	if match != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	// Generate encoded token and send it as response.
	loginToken, err := token.SignedString([]byte("024QwiCG0"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	// Response
	return c.JSON(fiber.Map{
		"token":    loginToken,
		"username": res.Username,
	})
}
