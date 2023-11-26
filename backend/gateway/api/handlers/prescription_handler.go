package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tommylay1902/gateway/internal/helper"
	"github.com/tommylay1902/gateway/internal/types"
)

type PrescriptionHandler struct {
	BaseUrl string
}

func InitializePrescriptionHandler(baseUrl string) *PrescriptionHandler {
	return &PrescriptionHandler{BaseUrl: baseUrl}
}

func (ph *PrescriptionHandler) GetPrescriptionById(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["sub"].(string))

	idParam := c.Params("id")

	resp, err := helper.MakeRequest("GET", ph.BaseUrl+"/"+idParam, nil)
	if err != nil {
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

	var prescription types.Prescription

	json.NewDecoder(resp.Body).Decode(&prescription)

	return c.Status(resp.StatusCode).JSON(prescription)

}

func (ph *PrescriptionHandler) GetPrescriptions(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["sub"].(string))
	resp, err := helper.MakeRequest("GET", ph.BaseUrl, nil)
	if err != nil {
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

	var prescription []types.Prescription

	json.NewDecoder(resp.Body).Decode(&prescription)

	return c.Status(resp.StatusCode).JSON(prescription)

}
