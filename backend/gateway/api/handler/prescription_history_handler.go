package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tommylay1902/gateway/internal/customtype/encoder"
	"github.com/tommylay1902/gateway/internal/helper"
)

type PrescriptionHistoryHandler struct {
	BaseUrl string
}

func InitializePrescriptionHistory(baseUrl string) *PrescriptionHistoryHandler {
	return &PrescriptionHistoryHandler{BaseUrl: baseUrl}
}

func (h *PrescriptionHistoryHandler) CreateHistory(c *fiber.Ctx) error {

	// resultBody := string(updatedJSON)

	body := string(c.Body())
	token := c.Locals("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	email := claims["sub"].(string)

	var data map[string]interface{}

	// Unmarshal the JSON string into the map
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// Add the additional field
	data["owner"] = email

	updatedJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	resultBody := string(updatedJSON)

	resp, err := helper.MakeRequest("POST", h.BaseUrl, &resultBody)

	if err != nil {
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

	var success encoder.Success
	json.NewDecoder(resp.Body).Decode(&success)
	return c.Status(resp.StatusCode).JSON(success)
}
