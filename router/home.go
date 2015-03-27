package router

import (
	"net/http"

	"github.com/martini-contrib/render"
)

func HomeHandler(res http.ResponseWriter, req *http.Request, r render.Render) {
	r.HTML(200, "home", "Steven")
}
