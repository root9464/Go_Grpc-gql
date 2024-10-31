package middleware

import (
	"context"
	"fmt"
	"root/shared/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func LoggerHTTP(log *logrus.Logger) fiber.Handler {
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

func LoggerGRPC(log *logrus.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {

		start := time.Now()
		resp, err = handler(ctx, req)
		latency := time.Since(start)
		latencyStr := fmt.Sprintf("%.2fms", float64(latency.Microseconds())/1000)
		code := status.Code(err)

		entry := log.WithFields(logrus.Fields{
			"\nMethod           ": info.FullMethod,
			"\nLatency          ": latencyStr,
			"\nClient IP        ": utils.GetClientIP(ctx),
			"\nRequest          ": fmt.Sprintf("%+v", req),
			"\nResponse         ": fmt.Sprintf("%+v", resp),
			"\nStatus Code      ": code,
		})

		if err != nil {
			entry.WithError(err).Error("❌ Request error")
		} else {
			entry.Info("✅ Request successful")
		}

		return resp, err
	}
}
