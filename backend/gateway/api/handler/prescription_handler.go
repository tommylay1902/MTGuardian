package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tommylay1902/gateway/internal/customtype"
	"github.com/tommylay1902/gateway/internal/customtype/dto"
	"github.com/tommylay1902/gateway/internal/customtype/encoder"
	"github.com/tommylay1902/gateway/internal/helper"
)

type PrescriptionHandler struct {
	BaseUrl string
}

func InitializePrescription(baseUrl string) *PrescriptionHandler {
	return &PrescriptionHandler{BaseUrl: baseUrl}
}

func (ph *PrescriptionHandler) GetPrescriptionById(c *fiber.Ctx) error {
	idParam := c.Params("id")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email, ok := claims["sub"].(string)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	resp, err := helper.MakeRequest("GET", ph.BaseUrl+"/"+email+"/"+idParam, nil)
	if err != nil {
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

	var prescription customtype.Prescription

	json.NewDecoder(resp.Body).Decode(&prescription)

	return c.Status(resp.StatusCode).JSON(prescription)
}

func (ph *PrescriptionHandler) GetPrescriptions(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email := claims["sub"].(string)
	viewHistory := c.Query("present")

	url := fmt.Sprintf("%s/all/%s?present=%s", ph.BaseUrl, email, viewHistory)

	if viewHistory == "" {
		url = fmt.Sprintf("%s/all/%s", ph.BaseUrl, email)
	}

	resp, err := helper.MakeRequest("GET", url, nil)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var bodyErr encoder.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}

	var prescription []customtype.Prescription
	json.NewDecoder(resp.Body).Decode(&prescription)

	return c.Status(resp.StatusCode).JSON(prescription)
}

func (ph *PrescriptionHandler) CreatePrescription(c *fiber.Ctx) error {
	//figure out why it wont take BodyParser
	prescription := string(c.Body())
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email := claims["sub"]

	var data map[string]interface{}

	err := json.Unmarshal([]byte(prescription), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	data["owner"] = email

	updatedJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	resultBody := string(updatedJSON)

	resp, err := helper.MakeRequest("POST", ph.BaseUrl, &resultBody)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Println("ERROR")
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

func (ph *PrescriptionHandler) UpdatePrescription(c *fiber.Ctx) error {
	var prescription dto.PrescriptionDTO
	err := c.BodyParser(&prescription)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	idParam := c.Params("id")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email, ok := claims["sub"].(string)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	prescription.Owner = &email
	updatedJSON, err := json.Marshal(prescription)

	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	resultBody := string(updatedJSON)
	resp, err := helper.MakeRequest("PUT", ph.BaseUrl+"/"+email+"/"+idParam, &resultBody)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		var bodyErr encoder.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}

	defer resp.Body.Close()

	var success encoder.Success
	json.NewDecoder(resp.Body).Decode(&success)
	return c.Status(resp.StatusCode).JSON(success)
}

func (ph *PrescriptionHandler) DeletePrescription(c *fiber.Ctx) error {
	idParam := c.Params("id")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email, ok := claims["sub"].(string)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Something went wrong"})
	}

	resp, err := helper.MakeRequest("DELETE", ph.BaseUrl+"/"+email+"/"+idParam, nil)

	if err != nil {
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

	return c.Status(resp.StatusCode).JSON(fiber.Map{"success": "succesfully deleted prescription"})
}
