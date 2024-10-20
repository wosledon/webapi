package web

type Middleware interface{}

type WebApplicationBuilder struct {
	Services    *ServiceCollection
	Middleware  *ServiceCollection
	Controllers *ServiceCollection
}

func CreateBuilder() *WebApplicationBuilder {
	return &WebApplicationBuilder{
		Services: &ServiceCollection{},
	}
}

func (builder *WebApplicationBuilder) Build() *WebApplication {
	serviceProvider := builder.Services.BuildServiceProvider()
	return NewWebApplication(serviceProvider)
}
