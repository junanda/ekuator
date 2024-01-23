package controller

import (
	"ekuator/product/models"
	"ekuator/product/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

type BaseController interface {
	Route(app *fiber.App)
}

type baseControllerImpl struct {
	prodServ services.ProductServiceInterface
}

func NewController(cs services.ProductServiceInterface) BaseController {
	return &baseControllerImpl{
		prodServ: cs,
	}
}

func (b *baseControllerImpl) Route(app *fiber.App) {
	app.Get("/product", func(c *fiber.Ctx) error {
		result := b.prodServ.GetAllProduct()
		return c.JSON(result)
	})

	app.Get("/product/:id", func(c *fiber.Ctx) error {
		idprod := c.Params("id")
		result := b.prodServ.DetailProduct(idprod)
		return c.JSON(models.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   result,
		})
	})

	app.Post("/product", func(c *fiber.Ctx) error {
		var (
			datareq  models.Product
			response models.WebResponse
			err      error
		)
		err = c.BodyParser(&datareq)
		if err != nil {
			log.Println("error parsing data: ", err.Error())
		}

		err = b.prodServ.SaveProduct(datareq)
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

	app.Delete("/product/:id", func(c *fiber.Ctx) error {
		var (
			response models.WebResponse
		)
		idProd := c.Params("id")

		err := b.prodServ.DeleteProduct(idProd)
		if err != nil {
			response = models.WebResponse{
				Code:   501,
				Status: "Failed Delete Product",
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
