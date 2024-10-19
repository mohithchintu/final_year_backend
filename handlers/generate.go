package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mohithchintu/final_year_project/helpers"
	"github.com/mohithchintu/final_year_project/models"
	"github.com/mohithchintu/final_year_project/sss"
	"github.com/mohithchintu/final_year_project/utils"
)

type GenerateRequest struct {
	NumDevices int `json:"num_devices"`
	Threshold  int `json:"threshold"`
}

var devices []*models.Device

func GenerateDevices(c *fiber.Ctx) error {
	req := new(GenerateRequest)

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	devices = []*models.Device{}

	for i := 1; i <= req.NumDevices; i++ {
		deviceID := "Device" + strconv.Itoa(i)
		device := helpers.InitializeDevice(deviceID, req.Threshold)
		devices = append(devices, device)
	}

	// for i, device := range devices {
	// 	for j, peer := range devices {
	// 		if i != j {
	// 			device.Peers[peer.ID] = peer
	// 		}
	// 	}
	// }

	for _, device := range devices {
		coefficients, _ := sss.GenerateAndSharePolynomial(device, req.Threshold-1)
		shares := sss.GenerateShares(coefficients, req.NumDevices)
		utils.DistributeShares(device, shares)
	}

	// reconstructedKey := sss.HandleDeviceFailure(devices, req.Threshold)

	// fmt.Println(reconstructedKey)

	return c.JSON(fiber.Map{
		"devices": devices,
	})
}
