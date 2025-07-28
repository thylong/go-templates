package server

import (
	"net/http"
	"time"

	"github.com/go-fuego/fuego"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/thylong/go-templates/07-fuego-sqlc/api"
	"github.com/thylong/go-templates/07-fuego-sqlc/pkg/db"
)

func CreateApp(production bool, httpTimeout int64, queries *db.Queries) *fuego.Server {
	router := fuego.NewServer()

	/*
	   database setup:
	   - Initialize connection to the database
	   - Start session and store in gin context

	   Fuego setup:
	   - Chain global middlewares (logger, timeout, recovery, etc)
	   - Setup configuration (dev vs prod, log format, gracefull shutdown)
	   - Create healthcheck route /healthz
	   - Setup configuration (dev vs prod, log format)
	*/

	// Logger middleware will:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - Logs using JSON format.
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	fuego.Use(router, LoggingMiddleware)
	// timeout middleware will return 408 if timeout is reached.
	// router.Use(middlewares.TimeoutMiddleware(httpTimeout))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// router.Use(gin.Recovery())

	api.SetupRoutes(router, queries)

	return router
}

// LoggingMiddleware logs incoming HTTP requests and their responses.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Wrap the ResponseWriter to capture the status code.
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		// Record the start time.
		startTime := time.Now()

		// Process the request.
		next.ServeHTTP(lrw, r)

		// Log the details of the request and response.
		duration := time.Since(startTime)
		log.Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", lrw.statusCode).
			Str("remote_addr", r.RemoteAddr).
			Dur("duration", duration).
			Msg("Handled request")
	})
}

// loggingResponseWriter is a wrapper around http.ResponseWriter that captures the status code.
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
