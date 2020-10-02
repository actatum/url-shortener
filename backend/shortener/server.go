package shortener

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type server struct {
	service *service
}

func newServer(s *service) *server {
	return &server{
		service: s,
	}
}

func (s *server) Create(c echo.Context) error {
	var req Request
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	resp, err := s.service.Create(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) Read(c echo.Context) error {
	var req Request
	req.Slug = c.Param("id")

	resp, err := s.service.Read(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *server) Top5(c echo.Context) error {
	resp, err := s.service.Top5()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}
