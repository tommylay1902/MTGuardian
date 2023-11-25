package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/internal/helper"
	"github.com/tommylay1902/gateway/internal/types"
)

type AuthHandler struct {
	BaseUrl string
}

func InitializeAuthHandler(baseUrl string) *AuthHandler {
	return &AuthHandler{BaseUrl: baseUrl}
}

func (ah *AuthHandler) RegisterHandler(c *fiber.Ctx) error {

	resp, err := helper.MakeRequest(
		"POST", ah.BaseUrl+"/register", string(c.Body()))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()
	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		var bodyErr types.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}
	var token types.AccessToken
	json.NewDecoder(resp.Body).Decode(&token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (ah *AuthHandler) LoginHandler(c *fiber.Ctx) error {

	// Send the request
	resp, err := helper.MakeRequest(
		"POST", ah.BaseUrl+"/login", string(c.Body()))
	if err != nil {
		// Handle error
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		var bodyErr types.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}

	var token types.AccessToken
	json.NewDecoder(resp.Body).Decode(&token)

	return c.Status(resp.StatusCode).JSON(fiber.Map{
		"token": token,
	})
}

func (ah *AuthHandler) RefreshHandler(c *fiber.Ctx) error {

	resp, err := helper.MakeRequest(
		"POST", ah.BaseUrl+"/refresh", string(c.Body()))
	if err != nil {
		// Handle error
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		var bodyErr types.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}

	var token types.AccessToken
	json.NewDecoder(resp.Body).Decode(&token)

	return c.Status(resp.StatusCode).JSON(fiber.Map{
		"token": token,
	})
}
