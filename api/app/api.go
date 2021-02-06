package api

import (
	"compress/flate"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/zea7ot/web_api_aeyesafe/logging"
)

// New configures application resources and routes
func New(enableCORS bool) (*chi.Mux, error) {
	logger := logging.NewLogger()

	// database access

	// app api

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	compressor := middleware.NewCompressor(flate.DefaultCompression)
	r.Use(compressor.Handler)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Use(logging.NewStructuredLogger(logger))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	if enableCORS {
		r.Use(corsConfig().Handler)
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return r, nil
}

func corsConfig() *cors.Cors {
	// Basic CORS
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // to allow specific origin host(s)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, // maximum value not ignored by any major browswer
	})
}
