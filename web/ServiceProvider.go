package web

import (
	"reflect"
)

type ServiceProvider struct {
	Services []ServiceDescriptor
}

func (sp *ServiceProvider) GetService(serviceType interface{}) interface{} {
	for _, service := range sp.Services {
		if service.ServiceType == reflect.TypeOf(serviceType) {
			return service.Implementation
		}
	}
	return nil
}

func (sp *ServiceProvider) GetServices(serviceType interface{}) []interface{} {
	var results []interface{}
	for _, service := range sp.Services {
		if service.ServiceType == reflect.TypeOf(serviceType) {
			results = append(results, service.Implementation)
		}
	}
	return results
}
