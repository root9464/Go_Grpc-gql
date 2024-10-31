package utils

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net"

	"google.golang.org/grpc/peer"
)

func GenerateRandomID() string {
	id := make([]byte, 16) // 16 байт = 128 бит
	if _, err := rand.Read(id); err != nil {
		panic(err)
	}
	return hex.EncodeToString(id)
}

func GetClientIP(ctx context.Context) string {
	if p, ok := peer.FromContext(ctx); ok {
		if addr, ok := p.Addr.(*net.TCPAddr); ok {
			return addr.IP.String()
		}
	}
	return "unknown"
}
