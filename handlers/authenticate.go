package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mohithchintu/final_year_project/models"
	"github.com/mohithchintu/final_year_project/sss"
)

type AuthenticateRequest struct {
	DeviceIDs []string `json:"device_ids"`
}

func AuthenticateDevices(c *fiber.Ctx) error {
	req := new(AuthenticateRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	if len(req.DeviceIDs) < len(devices[0].Peers)/2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Not enough devices to meet threshold",
		})
	}

	authenticated := true
	for _, deviceID := range req.DeviceIDs {
		device := findDeviceByID(deviceID)
		if device == nil {
			authenticated = false
			break
		}
	}

	if !authenticated {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authentication failed",
		})
	}

	// Simulate HMAC authentication (replace with actual logic)
	groupKey := sss.HandleDeviceFailure(devices, len(req.DeviceIDs))

	return c.JSON(fiber.Map{
		"message":  "Authentication successful",
		"groupkey": groupKey,
	})

	// return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//     "error": "Authentication failed",
	// })
}

func findDeviceByID(deviceID string) *models.Device {
	for _, device := range devices {
		if device.ID == deviceID {
			return device
		}
	}
	return nil
}
