package web

import (
	"net/http"
)

type WebApplication struct {
	ServiceProvider *ServiceProvider
	Mux             *http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func (app *WebApplication) Use(middleware func(http.Handler) http.Handler) {
    app.middlewares = append(app.middlewares, middleware)
}

func NewWebApplication(serviceProvider *ServiceProvider) *WebApplication {
	return &WebApplication{
		ServiceProvider: serviceProvider,
		Mux:             http.NewServeMux(),
	}
}

func (app *WebApplication) registerControllers() {
	controllers := app.ServiceProvider.GetServices((*Controller)(nil))
	for _, ctrl := range controllers {
		controller := ctrl.(Controller)
		controller.RegisterRoutes(app.Mux, controller)
	}
}

func (app *WebApplication) UseHttpsRedirection(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil {
			http.Redirect(w, r, "https://"+r.Host+r.RequestURI, http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *WebApplication) UseStaticFiles(path string) {
	app.Mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path))))
}


func (app *WebApplication) Run(addr string) {
	app.registerControllers()

	finalHandler := http.Handler(app.Mux)
    for i := len(app.middlewares) - 1; i >= 0; i-- {
        finalHandler = app.middlewares[i](finalHandler)
    }

	http.ListenAndServe(addr, app.Mux)
}
