package main

import (
	"ai-saas/handler"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := initEverything(); err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()

	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))

	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("Application is running on port: ", "port", port)
	log.Fatal(http.ListenAndServe(os.Getenv(port), router))

}

func initEverything() error {
	// if err := godotenv.Load; err != nil {
	// 	return err
	// }
	// return nil
	return godotenv.Load()
}
