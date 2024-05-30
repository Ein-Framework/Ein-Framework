package apiservicemanager

import (
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func (s *ApiService) POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.server.POST(filepath.Join(s.path, path), h, m...)
}

func (s *ApiService) GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.server.GET(filepath.Join(s.path, path), h, m...)
}

func (s *ApiService) PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.server.PATCH(filepath.Join(s.path, path), h, m...)
}

func (s *ApiService) OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.server.OPTIONS(filepath.Join(s.path, path), h, m...)
}

func (s *ApiService) PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.server.PUT(filepath.Join(s.path, path), h, m...)
}

func (s *ApiService) DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.server.DELETE(filepath.Join(s.path, path), h, m...)
}

func (s *ApiService) HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.server.HEAD(filepath.Join(s.path, path), h, m...)
}

func (s *ApiService) TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return s.server.TRACE(filepath.Join(s.path, path), h, m...)
}
