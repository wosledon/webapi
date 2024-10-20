package main

import (
	"bsi/webapi/api"
	"bsi/webapi/services"
	WebApplication "bsi/webapi/web"
	"net/http"
)

func main() {

	builder := WebApplication.CreateBuilder()

	builder.Services.AddControllers(
		&api.HelloController{},
		&api.TestController{},
	)

	builder.Services.AddSingleton(&services.UserService{})

	app := builder.Build()

	app.UseStaticFiles("index.html")

	app.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	app.Run(":8080")
}
