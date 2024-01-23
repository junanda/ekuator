package main

import (
	"ekuator/costumer/config"
	"ekuator/costumer/controller"
	"ekuator/costumer/repository"
	"ekuator/costumer/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db := config.NewMysqlDatabase().Initialize()

	defer db.Db.Close()

	csRepo := repository.NewCostumerRepository(db)
	csService := services.NewCostumerService(csRepo)
	controller := controller.NewController(csService)

	controller.Route(app)

	app.Listen(":3000")
}
