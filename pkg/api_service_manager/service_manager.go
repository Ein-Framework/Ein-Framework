package apiservicemanager

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func NewServiceManager(server *echo.Echo) *ServiceManager {
	return &ServiceManager{
		server:   server,
		services: make(map[string]*ApiService),
	}
}

func (m *ServiceManager) NewService(path string) (*ApiService, error) {
	_, ok := m.services[path]
	if ok {
		return nil, errors.New("service already exists")
	}

	service := &ApiService{
		server: m.server,
		path:   path,
	}
	m.services[path] = service
	return service, nil
}
