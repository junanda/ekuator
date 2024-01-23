package main

import (
	"ekuator/product/config"
	"ekuator/product/controller"
	"ekuator/product/repository"
	"ekuator/product/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := config.NewMysqlDatabase().Initialize()

	defer db.Db.Close()

	proRepo := repository.NewProductRepository(db)
	csService := services.NewProductService(proRepo)
	controller := controller.NewController(csService)

	controller.Route(app)

	app.Listen(":3001")
}
