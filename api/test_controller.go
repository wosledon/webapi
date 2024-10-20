package api

import (
	"bsi/webapi/web"
	"net/http"
)

type TestController struct {
	web.ControllerBase
}

func (c *HelloController) HandleTest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Test!"))
}
