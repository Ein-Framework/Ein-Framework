package apiservicemanager

import (
	"github.com/labstack/echo/v4"
)

func (s *ApiService) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.group.POST(path, h, m...)
}

func (s *ApiService) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.group.GET(path, h, m...)
}

func (s *ApiService) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.group.PATCH(path, h, m...)
}

func (s *ApiService) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.group.OPTIONS(path, h, m...)
}

func (s *ApiService) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.group.PUT(path, h, m...)
}

func (s *ApiService) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.group.DELETE(path, h, m...)
}

func (s *ApiService) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.group.HEAD(path, h, m...)
}

func (s *ApiService) TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.group.TRACE(path, h, m...)
}

func (s *ApiService) Use(middleware ...echo.MiddlewareFunc) {
	s.group.Use(middleware...)
}
