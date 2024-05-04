package http

import (
	"context"
	"log/slog"
	"math/big"
	"net"
	"time"
)

const powDifficulty = 10

var quotes []string

type Handlerer interface {
	HandleConnection(ctx context.Context, conn net.Conn)
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleConnection(ctx context.Context, conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			slog.ErrorContext(ctx, "failed to close connection", "err", err)
		}
	}(conn)

	challenge := make([]byte, 32)
	_, err := conn.Read(challenge)
	if err != nil {
		slog.ErrorContext(ctx, "failed to read challenge", "err", err)
		return
	}

	if !pow(challenge, powDifficulty) {
		slog.ErrorContext(ctx, "invalid proof of work")
		return
	}

	quote := quotes[time.Now().Unix()%int64(len(quotes))]
	slog.InfoContext(ctx, "connected and running", "conn", conn, "quote", quote)
}

func pow(data []byte, difficulty int) bool {
	target := big.NewInt(0)
	target.Exp(big.NewInt(2), big.NewInt(int64(8*len(data))), nil).Sub(target, big.NewInt(1))
	target.Lsh(target, uint(256-difficulty)).Div(target, big.NewInt(int64(time.Second/time.Millisecond)))

	guess := big.NewInt(0).SetBytes(data)
	for guess.Cmp(target) <= 0 {
		guess.Add(guess, big.NewInt(1))
	}

	return guess.Cmp(target) <= 0
}
