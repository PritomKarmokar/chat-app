package main

import (
	"net/http"

	"github.com/PritomKarmokar/chat-app/cmd/config"
	"github.com/go-chi/chi"
)

func main() {
	config.LoadEnv()
	config.LoggerConfig()

	logger := config.GetLogger()

	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	logger.Info().Msg("Server starting on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Fatal().Err(err).Msg("Server stopped")
	}
}
