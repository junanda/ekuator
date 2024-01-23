package controller

import (
	"ekuator/costumer/models"
	"ekuator/costumer/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type BaseController interface {
	Route(app *fiber.App)
}

type baseControllerImpl struct {
	costumerServ services.CostumerServiceInterface
}

func NewController(cs services.CostumerServiceInterface) BaseController {
	return &baseControllerImpl{
		costumerServ: cs,
	}
}

func (b *baseControllerImpl) Route(app *fiber.App) {
	app.Get("/costumer", func(c *fiber.Ctx) error {
		result := b.costumerServ.GetAllCostumer()
		return c.JSON(result)
	})

	app.Get("/costumer/:id", func(c *fiber.Ctx) error {
		idCostumer := c.Params("id")
		result := b.costumerServ.DetailCostumer(idCostumer)
		return c.JSON(models.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   result,
		})
	})

	app.Post("/costumer", func(c *fiber.Ctx) error {
		var (
			dataReq  models.Costumer
			response models.WebResponse
		)

		err := c.BodyParser(&dataReq)
		if err != nil {
			log.Println("error parsing data: ", err.Error())
		}

		err = b.costumerServ.SaveCostumer(dataReq)
		if err != nil {
			response = models.WebResponse{
				Code:   501,
				Status: "Failed Insert Data",
				Data:   nil,
			}
		}

		response = models.WebResponse{
			Code:   201,
			Status: "Success",
			Data:   nil,
		}

		return c.JSON(response)
	})

	app.Delete("/costumer/:id", func(c *fiber.Ctx) error {
		var (
			response models.WebResponse
		)
		idCos := c.Params("id")
		err := b.costumerServ.DeleteCostumer(idCos)
		if err != nil {
			response = models.WebResponse{
				Code:   501,
				Status: "Failed Delete Costumer",
				Data:   nil,
			}
		}

		response = models.WebResponse{
			Code:   200,
			Status: "Success",
			Data:   nil,
		}

		return c.JSON(response)
	})
}
