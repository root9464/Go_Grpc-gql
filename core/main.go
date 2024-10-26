package main

import (
	"root/shared/database"
	"root/shared/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func main() {
	log := utils.Logger()
	_, err := database.ConnectDb()
	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Не удалось подключиться к базе данных")
	} else {
		log.Info("Подключение к базе данных прошло успешно")
	}
	app := fiber.New()
	app.Use(utils.LoggerServer())
	log.Info("Starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		logrus.Fatal("Error starting server:", err)
	}
}
