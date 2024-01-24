package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tommylay1902/gateway/internal/customtype"
	"github.com/tommylay1902/gateway/internal/customtype/dto"
	"github.com/tommylay1902/gateway/internal/customtype/encoder"
	"github.com/tommylay1902/gateway/internal/error/apperror"
	"github.com/tommylay1902/gateway/internal/error/errorhandler"
	"github.com/tommylay1902/gateway/internal/helper"
)

type PrescriptionHistoryHandler struct {
	BaseUrl string
}

func InitializePrescriptionHistory(baseUrl string) *PrescriptionHistoryHandler {
	return &PrescriptionHistoryHandler{BaseUrl: baseUrl}
}

func (h *PrescriptionHistoryHandler) CreateHistory(c *fiber.Ctx) error {
	var rxHistory dto.PrescriptionHistoryDTO
	err := c.BodyParser(&rxHistory)

	if err != nil {
		return errorhandler.HandleError(&apperror.BadRequestError{Message: "error"}, c)
	}

	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email, ok := claims["sub"].(string)

	if !ok {
		return errorhandler.HandleError(&apperror.BadRequestError{Message: "email conversion went wrong"}, c)
	}

	rxHistory.Owner = &email

	updatedJSON, err := json.Marshal(rxHistory)

	if err != nil {
		return errorhandler.HandleError(&apperror.BadRequestError{Message: err.Error()}, c)
	}

	resultBody := string(updatedJSON)
	resp, err := helper.MakeRequest("POST", h.BaseUrl, &resultBody)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	defer resp.Body.Close()

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

func (h *PrescriptionHistoryHandler) GetHistory(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email, ok := claims["sub"].(string)

	if !ok {
		return errorhandler.HandleError(&apperror.BadRequestError{Message: "email conversion went wrong"}, c)
	}

	resp, err := helper.MakeRequest("GET", h.BaseUrl+"/all"+"/"+email, nil)

	if err != nil {
		return errorhandler.HandleError(err, c)
	}

	defer resp.Body.Close()

	var rxHistories []customtype.PrescriptionHistory
	json.NewDecoder(resp.Body).Decode(&rxHistories)
	return c.Status(fiber.StatusOK).JSON(rxHistories)
}

func (h *PrescriptionHistoryHandler) GetHistoryById(c *fiber.Ctx) error {
	pId := c.Params("pId")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email, ok := claims["sub"].(string)

	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email conversion went wrong"})
	}

	resp, err := helper.MakeRequest("GET", h.BaseUrl+"/"+email+"/"+pId, nil)

	if err != nil {
		return errorhandler.HandleError(err, c)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var bodyErr encoder.Error
		json.NewDecoder(resp.Body).Decode(&bodyErr)
		return c.Status(resp.StatusCode).JSON(fiber.Map{
			"error": bodyErr.Error,
		})
	}

	var rxHistory customtype.PrescriptionHistory
	json.NewDecoder(resp.Body).Decode(&rxHistory)
	return c.Status(fiber.StatusOK).JSON(rxHistory)
}

func (h *PrescriptionHistoryHandler) UpdateRxHistory(c *fiber.Ctx) error {
	pId := c.Params("pId")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email, ok := claims["sub"].(string)

	if !ok {
		return errorhandler.HandleError(&apperror.BadRequestError{Message: "email conversion went wrong"}, c)
	}

	var rxHistory dto.PrescriptionHistoryDTO
	err := c.BodyParser(&rxHistory)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	rxHistory.Owner = &email
	updatedJSON, err := json.Marshal(rxHistory)

	if err != nil {
		return errorhandler.HandleError(err, c)
	}

	resultBody := string(updatedJSON)
	resp, err := helper.MakeRequest("PUT", h.BaseUrl+"/"+email+"/"+pId, &resultBody)

	if err != nil {
		return errorhandler.HandleError(err, c)
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

	var success encoder.Success
	json.NewDecoder(resp.Body).Decode(&success)
	return c.Status(resp.StatusCode).JSON(success)
}

func (h *PrescriptionHistoryHandler) DeleteByEmailAndRx(c *fiber.Ctx) error {
	pId := c.Params("pId")
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	email := claims["sub"].(string)

	resp, err := helper.MakeRequest("DELETE", h.BaseUrl+"/"+email+"/"+pId, nil)

	if err != nil {
		return errorhandler.HandleError(err, c)
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

	var success encoder.Success
	json.NewDecoder(resp.Body).Decode(&success)
	return c.Status(resp.StatusCode).JSON(success)
}
