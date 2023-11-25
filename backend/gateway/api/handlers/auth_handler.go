package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/internal/types"
)

type AuthHandler struct {
	BaseUrl string
}

func InitializeAuthHandler(baseUrl string) *AuthHandler {
	return &AuthHandler{BaseUrl: baseUrl}
}

func (ah *AuthHandler) RegisterHandler(c *fiber.Ctx) error {

	req, err := http.NewRequest("POST", ah.BaseUrl+"/register", strings.NewReader(string(c.Body())))
	if err != nil {
		// Handle error
		fmt.Println("req err", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := http.Client{
		Timeout: time.Second * 30,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	var token types.AccessToken
	json.NewDecoder(resp.Body).Decode(&token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (ah *AuthHandler) LoginHandler(c *fiber.Ctx) error {

	req, err := http.NewRequest("POST", ah.BaseUrl+"/login", strings.NewReader(string(c.Body())))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	req.Header.Set("Content-Type", "application/json")

	// Create an HTTP client
	client := http.Client{
		Timeout: time.Second * 30,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		// Handle error
		fmt.Println("res err", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()
	var token types.AccessToken
	json.NewDecoder(resp.Body).Decode(&token)

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
