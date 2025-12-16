package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "Конвертер Морзе: ", log.Ldate|log.Ltime|log.Lshortfile)
	srv := server.New(logger)
	if err := srv.Start(); err != nil {
		logger.Fatalf("Не получилось запустить сервер: %v", err)
	}
}
