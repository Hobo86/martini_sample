package admin

import (
	"net/http"

	"github.com/martini-contrib/render"
)

func UserHandler(res http.ResponseWriter, req *http.Request, r render.Render) {

	r.HTML(200, "admin/user", "Steven")
}
