package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "15:04:05",
		FullTimestamp:   true,
	})

	return log
}

func LoggerServer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()

		log := Logger()
		log.WithFields(logrus.Fields{
			"status":     c.Response().StatusCode(),
			"method":     c.Method(),
			"path":       c.OriginalURL(),
			"remoteAddr": c.IP(),
			"took":       time.Since(start),
		})
		return err
	}
}
