package api

import (
	"bsi/webapi/web"
	"net/http"
)

type HelloController struct {
	web.ControllerBase
}

func (c *HelloController) HandleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
