package server

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/rs/zerolog"
	"github.com/thylong/go-templates/05-gin-templ-htmx/api"
	"github.com/thylong/go-templates/05-gin-templ-htmx/pkg/middleware"
)

func CreateApp(production bool, httpTimeout int64) *gin.Engine {
	r := gin.New()

	/*
	   database setup:
	   - Initialize connection to the database
	   - Start session and store in gin context

	   Gin setup:
	   - Chain global middlewares (logger, timeout, recovery, etc)
	   - Setup configuration (dev vs prod, log format, gracefull shutdown)
	   - Create healthcheck route /healthz
	   - Setup configuration (dev vs prod, log format)
	*/

	// Disable Console Color, switch to gin ReleaseMode, configure logger.
	if production {
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)

		// Logger middleware will:
		//   - Logs all requests, like a combined access and error log.
		//   - Logs to stdout.
		//   - Logs using JSON format.
		r.Use(logger.SetLogger(logger.WithLogger(func(_ *gin.Context, l zerolog.Logger) zerolog.Logger {
			return l.Output(gin.DefaultWriter).With().Logger()
		})))

		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		// Logger middleware will:
		//   - Logs all requests, like a combined access and error log.
		//   - Logs to stdout.
		r.Use(logger.SetLogger())
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// timeout middleware will return 408 if timeout is reached.
	r.Use(middleware.TimeoutMiddleware(httpTimeout))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.HTMLRender = &TemplRender{}

	api.SetupRoutes(r)

	return r
}

type TemplRender struct {
	Code int
	Data templ.Component
}

func (t TemplRender) Render(w http.ResponseWriter) error {
	t.WriteContentType(w)
	w.WriteHeader(t.Code)
	if t.Data != nil {
		return t.Data.Render(context.Background(), w)
	}
	return nil
}

func (t TemplRender) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (t *TemplRender) Instance(name string, data interface{}) render.Render {
	if templData, ok := data.(templ.Component); ok {
		return &TemplRender{
			Code: http.StatusOK,
			Data: templData,
		}
	}
	return nil
}
