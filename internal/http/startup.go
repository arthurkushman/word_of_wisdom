package http

import (
	"bufio"
	"log/slog"
	"os"
)

const quoteFile = "quotes.txt"

func init() {
	file, err := os.Open(quoteFile)
	if err != nil {
		slog.Error("failed to open quotes file", "err", err)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			slog.Error("failed to close quotes file", "err", err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}
}
