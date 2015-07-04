package demo

import (
	"net/http"

	"github.com/martini-contrib/render"
)

func DemoHandler(res http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "demo/index", "")
}
