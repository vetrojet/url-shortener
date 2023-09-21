package main

import (
	"fmt"
	"url-shortener/internal/config"
)

func main() {
	// TODO: init config: cleanevn
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// TODO: init logger: slog log/slog

	// TODO: init storage: sqlite

	// TODO: init router: chi, 'chi render'

	// TODO: run server
}
