package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/arthurkushman/word_of_wisdom/internal/http"
)

const (
	serverAddr = ":8080"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	h := http.NewHandler()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go h.HandleConnection(ctx, conn)
	}
}
