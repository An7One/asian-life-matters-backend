package api

import (
	"compress/flate"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/zea7ot/web_api_aeyesafe/api/app"
	"github.com/zea7ot/web_api_aeyesafe/database"
	"github.com/zea7ot/web_api_aeyesafe/logging"
)

// New configures application resources and routes
func New(enableCORS bool) (*chi.Mux, error) {
	logger := logging.NewLogger()

	// database access
	db := database.DBConn()
	// if err != nil {
	// 	logger.WithField("module", "database").Error(err)
	// 	return nil, err
	// }

	// app api
	appAPI, err := app.NewAPI(db)
	if err != nil {
		logger.WithField("module", "app").Error(err)
		return nil, err
	}

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

	r.Group(func(r chi.Router) {
		r.Mount("/api", appAPI.Router())
	})

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
