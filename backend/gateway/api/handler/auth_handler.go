package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tommylay1902/gateway/internal/customtype/encoder"
	"github.com/tommylay1902/gateway/internal/helper"
)

type AuthHandler struct {
	BaseUrl string
}

func InitializeAuth(baseUrl string) *AuthHandler {
	return &AuthHandler{BaseUrl: baseUrl}
}

func (ah AuthHandler) RegisterHandler(c *fiber.Ctx) error {
	body := string(c.Body())

	resp, err := helper.MakeRequest(
		"POST", ah.BaseUrl+"/register", &body)

	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		var bodyErr encoder.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}

	var token encoder.AccessToken
	json.NewDecoder(resp.Body).Decode(&token)

	return c.Status(fiber.StatusOK).JSON(token)
}

func (ah AuthHandler) LoginHandler(c *fiber.Ctx) error {
	body := string(c.Body())

	// Send the request
	resp, err := helper.MakeRequest(
		"POST", ah.BaseUrl+"/login", &body)

	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		var bodyErr encoder.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}

	var token encoder.AccessToken
	json.NewDecoder(resp.Body).Decode(&token)

	return c.Status(resp.StatusCode).JSON(token)
}

func (ah AuthHandler) RefreshHandler(c *fiber.Ctx) error {
	body := string(c.Body())

	resp, err := helper.MakeRequest(
		"POST", ah.BaseUrl+"/refresh", &body)

	if err != nil {
		fmt.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		var bodyErr encoder.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}

	var token encoder.AccessToken
	json.NewDecoder(resp.Body).Decode(&token)

	return c.Status(resp.StatusCode).JSON(token)
}
