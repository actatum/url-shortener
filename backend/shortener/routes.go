package shortener

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func routes(s *server) *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(cors()))

	e.POST("/url", s.Create)
	e.GET("/url/:id", s.Read)
	e.GET("/url", s.Top5)

	return e
}

func cors() middleware.CORSConfig {
	return middleware.CORSConfig{
		// AllowOrigins: []string{"http://localhost:3000", "*"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderAccept, echo.HeaderOrigin, echo.HeaderContentType},
	}
}

// Run starts the http server
func Run() error {
	repo, err := newRepository()
	if err != nil {
		return err
	}
	logger, err := newLogger()
	if err != nil {
		return err
	}
	service := newService(repo, logger)
	server := newServer(service)
	r := routes(server)
	port := os.Getenv("PORT")
	return r.Start(":" + port)
}
