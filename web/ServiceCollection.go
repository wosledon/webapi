package web

import (
	"reflect"
)

type ServiceDescriptor struct {
	ServiceType    reflect.Type
	Implementation interface{}
}

type ServiceCollection struct {
	services []ServiceDescriptor
}

func CreateServiceCollection() *ServiceCollection {
	return &ServiceCollection{}
}

func (sc *ServiceCollection) AddControllers(controller ...Controller) {
	for _, c := range controller {
		sc.AddController(c)
	}
}

func (sc *ServiceCollection) AddController(controller Controller) {
	sc.AddService((*Controller)(nil), controller)
}

func (sc *ServiceCollection) AddService(serviceType interface{}, implementation interface{}) {
	sc.services = append(sc.services, ServiceDescriptor{
		ServiceType:    reflect.TypeOf(serviceType),
		Implementation: implementation,
	})
}

func (sc *ServiceCollection) AddSingleton(implementation interface{}) {
	sc.services = append(sc.services, ServiceDescriptor{
		ServiceType:    reflect.TypeOf(implementation),
		Implementation: implementation,
	})
}

func (sc *ServiceCollection) BuildServiceProvider() *ServiceProvider {
	return &ServiceProvider{Services: sc.services}
}
