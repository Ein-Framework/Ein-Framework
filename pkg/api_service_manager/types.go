package apiservicemanager

import "github.com/labstack/echo/v4"

type ApiService struct {
	server *echo.Echo
	path   string
}

type ServiceManager struct {
	server   *echo.Echo
	services map[string]*ApiService
}
