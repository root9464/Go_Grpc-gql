package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Logger(log *logrus.Logger) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()
		err := ctx.Next()
		latency := time.Since(start)
		latencyStr := fmt.Sprintf("%.2fms", float64(latency.Microseconds())/1000)

		entry := log.WithFields(logrus.Fields{
			"\nClient IP    ": ctx.IP(),
			"\nStatus       ": ctx.Response().StatusCode(),
			"\nLatency      ": latencyStr,
			"\nUser-Agent   ": ctx.Get("User-Agent"),
			"\nContent-Type ": ctx.Get("Content-Type"),
			"\nPath         ": ctx.Path(),
			"\nProtocol     ": ctx.Protocol(),
			"\nMethod       ": ctx.Method(),
		})

		if err != nil || ctx.Response().StatusCode() >= 400 {
			entry.WithField("Error", err).Error("❌ Request error")
		} else {
			entry.Info("✅ Request successful")
		}

		return err
	}
}
